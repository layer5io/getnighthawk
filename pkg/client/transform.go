package nighthawk

import (
	"errors"
	"fmt"
	"strings"
	"time"

	nighthawk_client "github.com/layer5io/nighthawk-go/pkg/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/durationpb"
)

// func Transform transforms nighthawk's json output to fortio compatible json output
// This implementation is adpated from nighthawk's own transformer
// https://github.com/envoyproxy/nighthawk/blob/main/source/client/output_formatter_impl.cc
func Transform(res *nighthawk_client.ExecutionResponse) ([]byte, error) {
	resFortio := &nighthawk_client.FortioResult{}

	var workers int
	if workers = len(res.Output.Results); workers != 1 {
		workers--
	}
	resFortio.Labels = strings.Join(res.Output.GetOptions().GetLabels(), " ")
	resFortio.Version = res.Output.GetVersion().GetVersion().String()
	resFortio.StartTime = res.Output.GetTimestamp()
	resFortio.RequestedQPS = uint32(workers) * res.Output.Options.RequestsPerSecond.GetValue()
	resFortio.URL = res.Output.Options.GetUri().GetValue()
	resFortio.RequestedDuration = durationpb.New(res.Output.Options.GetDuration().AsDuration())
	// actual duration
	avgExecutionDuration, err := getAverageExecutionDuration(res)
	if err != nil {
		return nil, err
	}
	resFortio.ActualDuration = avgExecutionDuration.Seconds()
	// set jitter
	hasJitter := false
	if jitter := res.Output.Options.GetJitterUniform(); jitter != nil {
		hasJitter = true
	}
	resFortio.Jitter = hasJitter
	resFortio.RunType = "HTTP"
	resFortio.NumThreads = res.Output.Options.Connections.GetValue() * uint32(workers)
	globalResult := getGlobalResult(res)
	if globalResult == nil {
		return nil, errors.New("error")
	}
	resFortio.ActualQPS = float64(getCounterValue(globalResult, "upstream_rq_total").GetValue()) / resFortio.ActualDuration
	resFortio.BytesReceived = getCounterValue(globalResult, "upstream_cx_rx_bytes_total").GetValue()
	resFortio.BytesSent = getCounterValue(globalResult, "upstream_cx_tx_bytes_total").GetValue()
	mRetCodes := make(map[string]uint64, 1)
	mRetCodes["200"] = getCounterValue(globalResult, "benchmark.http_2xx").GetValue()
	resFortio.RetCodes = mRetCodes
	statistic := findStatistic(globalResult, "benchmark_http_client.request_to_response")
	if statistic != nil {
		resFortio.DurationHistogram = renderFortioDurationHistogram(statistic)
	}
	statistic = findStatistic(globalResult, "benchmark_http_client.response_body_size")
	if statistic != nil {
		resFortio.Sizes = renderFortioDurationHistogram(statistic)
	}
	statistic = findStatistic(globalResult, "benchmark_http_client.response_header_size")
	if statistic != nil {
		resFortio.HeaderSizes = renderFortioDurationHistogram(statistic)
	}

	out, err := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(resFortio)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func getAverageExecutionDuration(res *nighthawk_client.ExecutionResponse) (time.Duration, error) {
	resultsLen := len(res.Output.Results)
	if resultsLen == 0 {
		return 0, errors.New("no results found")
	}

	avgExecutionDuration := res.Output.Results[resultsLen-1].ExecutionDuration.AsDuration()
	return avgExecutionDuration, nil
}

func getGlobalResult(res *nighthawk_client.ExecutionResponse) *nighthawk_client.Result {
	for _, result := range res.Output.Results {
		if result.GetName() == "global" {
			return result
		}
	}
	return nil
}

func getCounterValue(result *nighthawk_client.Result, counterName string) *nighthawk_client.Counter {
	for _, counter := range result.Counters {
		if counter.GetName() == counterName {
			return counter
		}
	}
	return &nighthawk_client.Counter{Value: 0}
}

func findStatistic(result *nighthawk_client.Result, statID string) *nighthawk_client.Statistic {
	for _, stat := range result.Statistics {
		if stat.Id == statID {
			return stat
		}
	}
	return nil
}

func renderFortioDurationHistogram(stat *nighthawk_client.Statistic) *nighthawk_client.DurationHistogram {
	fortioHistogram := &nighthawk_client.DurationHistogram{}

	var prevFortioCount uint64 = 0
	var prevFortioEnd float64 = 0

	for i, p := range stat.Percentiles {
		dataEntry := &nighthawk_client.DataEntry{}
		dataEntry.Percent = (p.GetPercentile() * 100)
		dataEntry.Count = p.GetCount() - prevFortioCount

		var value float64
		if d := p.GetDuration(); d != nil {
			value = p.GetDuration().AsDuration().Seconds()
		} else {
			value = p.GetRawValue()
		}

		dataEntry.End = value

		// fortioStart = prevFortioEnd
		// If this is the first entry, force the start and end time to be the same.
		// This prevents it from starting at 0, making it disproportionally big in the UI.
		i++
		if i == 0 {
			prevFortioEnd = value
		}
		dataEntry.Start = prevFortioEnd

		prevFortioCount = p.GetCount()
		prevFortioEnd = value

		fortioHistogram.Data = append(fortioHistogram.Data, dataEntry)
	}

	fortioHistogram.Count = stat.GetCount()

	if mean := stat.GetMean(); mean != nil {
		fortioHistogram.Avg = mean.AsDuration().Seconds()
	} else {
		fortioHistogram.Avg = stat.GetRawMean()
	}

	if min := stat.GetMin(); min != nil {
		fortioHistogram.Min = min.AsDuration().Seconds()
	} else {
		fortioHistogram.Min = float64(stat.GetRawMin())
	}

	fortioHistogram.Sum = float64(stat.GetCount() * uint64(fortioHistogram.GetAvg()))

	if max := stat.GetMax(); max != nil {
		fortioHistogram.Max = max.AsDuration().Seconds()
	} else {
		fortioHistogram.Max = float64(stat.GetRawMax())
	}

	if pstDev := stat.GetPstdev(); pstDev != nil {
		fortioHistogram.StdDev = pstDev.AsDuration().Seconds()
	} else {
		fortioHistogram.StdDev = stat.GetRawPstdev()
	}

	iteratePercentiles(fortioHistogram, stat, func(fortioHistogram *nighthawk_client.DurationHistogram,
		percentile *nighthawk_client.Percentile) {
		if percentile.GetPercentile() > 0 && percentile.GetPercentile() < 1 {
			p := &nighthawk_client.FortioPercentile{}
			p.Percentile = (percentile.Percentile * 1000) / 10
			if duration := percentile.GetDuration(); duration != nil {
				p.Value = duration.AsDuration().Seconds()
			} else {
				p.Value = percentile.GetRawValue()
			}
			fortioHistogram.Percentiles = append(fortioHistogram.Percentiles, p)
		}
	})

	return fortioHistogram
}

type callback func(*nighthawk_client.DurationHistogram, *nighthawk_client.Percentile)

func iteratePercentiles(fortioHistogram *nighthawk_client.DurationHistogram, stat *nighthawk_client.Statistic, fn callback) {
	var lastPercentile float64 = -1

	percentiles := []float64{.0, .5, .75, .8, .9, .95, .99, .999, 1}
	for _, p := range percentiles {
		for _, percentile := range stat.Percentiles {
			if percentile.GetPercentile() >= p && lastPercentile < percentile.GetPercentile() {
				lastPercentile = percentile.GetPercentile()
				fn(fortioHistogram, percentile)
				fmt.Println(fortioHistogram.GetPercentiles())
				break
			}
		}
	}
}
