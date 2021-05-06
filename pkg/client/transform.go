package nighthawk

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/layer5io/meshkit/utils"
	nighthawk_client "github.com/layer5io/nighthawk-go/pkg/proto"
)

type Verbosity struct {
	Value string `json:"value"`
}

type OutputFormat struct {
	Value string `json:"value"`
}

type AddressFamily struct {
	Value string `json:"value"`
}

type RequestOption struct {
	RequestMethod  string                  `json:"request_method"`
	RequestHeaders []*v3.HeaderValueOption `json:"request_headers"`
}

type SequencerIdleStrategy struct {
	Value string `json:"value"`
}

type ExperimentalH1ConnectionReuseStrategy struct {
	Value string `json:"value"`
}

type FailurePredicates struct {
	BenchmarkPoolConnectionFailure string `json:"benchmark.pool_connection_failure"`
	BenchmarkHTTP4Xx               string `json:"benchmark.http_4xx"`
	BenchmarkHTTP5Xx               string `json:"benchmark.http_5xx"`
	RequestsourceUpstreamRq5Xx     string `json:"requestsource.upstream_rq_5xx"`
}

type TransformOptions struct {
	RequestsPerSecond                     int                                    `json:"requests_per_second"`
	Connections                           int                                    `json:"connections"`
	Duration                              string                                 `json:"duration"`
	Timeout                               string                                 `json:"timeout"`
	H2                                    bool                                   `json:"h2"`
	Concurrency                           string                                 `json:"concurrency"`
	Verbosity                             *Verbosity                             `json:"verbosity"`
	OutputFormat                          *OutputFormat                          `json:"output_format"`
	AddressFamily                         *AddressFamily                         `json:"address_family"`
	RequestOptions                        *RequestOption                         `json:"request_options"`
	SequencerIdleStrategy                 *SequencerIdleStrategy                 `json:"sequencer_idle_strategy"`
	ExperimentalH1ConnectionReuseStrategy *ExperimentalH1ConnectionReuseStrategy `json:"experimental_h1_connection_reuse_strategy"`
	TerminationPredicates                 struct {
	} `json:"termination_predicates"`
	FailurePredicates                    *FailurePredicates    `json:"failure_predicates"`
	Labels                               []interface{}         `json:"labels"`
	StatsSinks                           []interface{}         `json:"stats_sinks"`
	PrefetchConnections                  bool                  `json:"prefetch_connections"`
	BurstSize                            *wrappers.UInt32Value `json:"burst_size"`
	MaxPendingRequests                   int                   `json:"max_pending_requests"`
	MaxActiveRequests                    int                   `json:"max_active_requests"`
	MaxRequestsPerConnection             int64                 `json:"max_requests_per_connection"`
	URI                                  string                `json:"uri"`
	Trace                                string                `json:"trace"`
	OpenLoop                             bool                  `json:"open_loop"`
	ExperimentalH2UseMultipleConnections bool                  `json:"experimental_h2_use_multiple_connections"`
	NighthawkService                     string                `json:"nighthawk_service"`
	SimpleWarmup                         bool                  `json:"simple_warmup"`
	StatsFlushInterval                   int                   `json:"stats_flush_interval"`
	LatencyResponseHeaderName            string                `json:"latency_response_header_name"`
}

type Percentile struct {
	Percentile int    `json:"percentile"`
	Count      string `json:"count"`
	Duration   string `json:"duration"`
}

type Statistic struct {
	Count       string       `json:"count"`
	ID          string       `json:"id"`
	Percentiles []Percentile `json:"percentiles"`
	Mean        string       `json:"mean,omitempty"`
	Pstdev      string       `json:"pstdev,omitempty"`
	Min         string       `json:"min,omitempty"`
	Max         string       `json:"max,omitempty"`
	RawMean     int          `json:"raw_mean,omitempty"`
	RawPstdev   int          `json:"raw_pstdev,omitempty"`
	RawMin      string       `json:"raw_min,omitempty"`
	RawMax      string       `json:"raw_max,omitempty"`
}

type Counter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Result struct {
	Name              string      `json:"name"`
	Statistics        []Statistic `json:"statistics"`
	Counters          []Counter   `json:"counters"`
	ExecutionDuration string      `json:"execution_duration"`
	ExecutionStart    time.Time   `json:"execution_start"`
}

type TransformResult struct {
	Options   *TransformOptions `json:"options"`
	Results   []Result          `json:"results"`
	Version   *v3.BuildVersion  `json:"version"`
	Timestamp string            `json:"timestamp"`
}

func Transform(res *nighthawk_client.ExecutionResponse, typ string) ([]byte, error) {

	//dur, err := time.ParseDuration(fmt.Sprintf("%ds%dµs", res.Output.Timestamp.Seconds, res.Output.Timestamp.Nanos))
	//if err != nil {
	//	return nil, err
	//}

	results := make([]Result, 0)

	for _, r := range res.Output.Results {
		statistics := make([]Statistic, 0)
		counters := make([]Counter, 0)
		for _, c := range r.Counters {
			counters = append(counters, Counter{
				Name:  c.Name,
				Value: strconv.Itoa(int(c.Value)),
			})
		}

		for _, s := range r.Statistics {
			percentiles := make([]Percentile, 0)
			for _, p := range s.Percentiles {
				percentiles = append(percentiles, Percentile{
					Percentile: int(p.Percentile),
					Count:      strconv.Itoa(int(p.Count)),
					Duration:   formatToNs(p.GetDuration()),
				})
			}

			sts := Statistic{
				Count:       strconv.Itoa(int(s.Count)),
				ID:          s.Id,
				Percentiles: percentiles,
				Mean:        formatToNs(s.GetMean()),
				Pstdev:      formatToNs(s.GetPstdev()),
				Min:         formatToNs(s.GetMin()),
				Max:         formatToNs(s.GetMax()),
			}
			if sts.Mean == "" {
				sts.RawMean = int(s.GetRawMean())
			}
			if sts.Pstdev == "" {
				sts.RawPstdev = int(s.GetRawPstdev())
			}
			if sts.Min == "" {
				sts.RawMin = strconv.Itoa(int(s.GetRawMin()))
			}
			if sts.Max == "" {
				sts.RawMax = strconv.Itoa(int(s.GetRawMax()))
			}
			statistics = append(statistics, sts)
		}
		results = append(results, Result{
			Name:              r.Name,
			ExecutionStart:    r.ExecutionStart.AsTime(),
			ExecutionDuration: formatToNs(r.ExecutionDuration),
			Statistics:        statistics,
			Counters:          counters,
		})
	}

	expStrategy := res.Output.Options.ExperimentalH1ConnectionReuseStrategy.String()
	if expStrategy == "" {
		expStrategy = "DEFAULT"
	}

	t := TransformResult{
		Version:   res.Output.Version,
		Timestamp: time.Now().Format(time.RFC3339),
		Options: &TransformOptions{
			RequestsPerSecond: int(res.Output.Options.RequestsPerSecond.GetValue()),
			Connections:       int(res.Output.Options.Connections.GetValue()),
			Duration:          formatToNs(res.Output.Options.GetDuration()),
			Timeout:           formatToNs(res.Output.Options.Timeout),
			H2:                res.Output.Options.H2.Value,
			Concurrency:       res.Output.Options.Concurrency.Value,
			Verbosity: &Verbosity{
				Value: res.Output.Options.Verbosity.Value.String(),
			},
			OutputFormat: &OutputFormat{
				Value: res.Output.Options.OutputFormat.Value.String(),
			},
			PrefetchConnections: res.Output.Options.PrefetchConnections.Value,
			BurstSize:           res.Output.Options.BurstSize,
			AddressFamily: &AddressFamily{
				Value: res.Output.Options.AddressFamily.Value.String(),
			},
			RequestOptions: &RequestOption{
				RequestMethod:  res.Output.Options.GetRequestOptions().RequestMethod.String(),
				RequestHeaders: res.Output.Options.GetRequestOptions().RequestHeaders,
			},
			MaxPendingRequests:       res.Output.Options.MaxPendingRequests.ProtoReflect().Descriptor().Index(),
			MaxActiveRequests:        res.Output.Options.MaxActiveRequests.ProtoReflect().Descriptor().Index(),
			MaxRequestsPerConnection: int64(res.Output.Options.MaxRequestsPerConnection.Value),
			SequencerIdleStrategy: &SequencerIdleStrategy{
				Value: res.Output.Options.SequencerIdleStrategy.Value.String(),
			},
			URI:   res.Output.Options.GetUri().Value,
			Trace: res.Output.Options.Trace.Value,
			ExperimentalH1ConnectionReuseStrategy: &ExperimentalH1ConnectionReuseStrategy{
				Value: expStrategy,
			},
			ExperimentalH2UseMultipleConnections: res.Output.Options.ExperimentalH2UseMultipleConnections.Value,
			FailurePredicates: &FailurePredicates{
				BenchmarkHTTP4Xx:               fmt.Sprint(res.Output.Options.FailurePredicates["benchmark.http_4xx"]),
				BenchmarkHTTP5Xx:               fmt.Sprint(res.Output.Options.FailurePredicates["benchmark.http_5xx"]),
				BenchmarkPoolConnectionFailure: fmt.Sprint(res.Output.Options.FailurePredicates["benchmark.pool_connection_failure"]),
				RequestsourceUpstreamRq5Xx:     fmt.Sprint(res.Output.Options.FailurePredicates["requestsource.upstream_rq_5xx"]),
			},
			OpenLoop:                  res.Output.Options.OpenLoop.Value,
			NighthawkService:          res.Output.Options.NighthawkService.Value,
			SimpleWarmup:              res.Output.Options.SimpleWarmup.Value,
			StatsFlushInterval:        res.Output.Options.GetStatsFlushInterval().ProtoReflect().Descriptor().Index(),
			LatencyResponseHeaderName: res.Output.Options.LatencyResponseHeaderName.Value,
		},
		Results: results,
	}
	input, err := utils.Marshal(t)
	if err != nil {
		return nil, err
	}
	fmt.Println("input: ", input)

	command := "/private/var/tmp/_bazel_abishekk/eaf1167d72f4772496616c435b301da8/execroot/nighthawk/bazel-out/darwin-opt/bin/nighthawk_output_transform"
	cmd := exec.Command(command, "--output-format", "fortio")
	cmd.Stdin = strings.NewReader(input)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	fmt.Println("output: ", string(out))

	// Hack due to bug in nighthawk
	m := map[string]interface{}{}
	err = json.Unmarshal(out, &m)
	m["RequestedQPS"] = fmt.Sprint(m["RequestedQPS"].(float64))

	if m["DurationHistogram"] != nil {
		dh, err := strconv.ParseInt(m["DurationHistogram"].(map[string]interface{})["Count"].(string), 10, 64)
		if err != nil {
			return nil, err
		}
		m["DurationHistogram"].(map[string]interface{})["Count"] = dh

		for _, c := range m["DurationHistogram"].(map[string]interface{})["Data"].([]interface{}) {
			h, err := strconv.ParseInt(c.(map[string]interface{})["Count"].(string), 10, 64)
			if err != nil {
				return nil, err
			}
			c.(map[string]interface{})["Count"] = h
		}
	}

	if m["HeaderSizes"] != nil {
		h, err := strconv.ParseInt(m["HeaderSizes"].(map[string]interface{})["Count"].(string), 10, 64)
		if err != nil {
			return nil, err
		}
		m["HeaderSizes"].(map[string]interface{})["Count"] = h
	}

	if m["RetCodes"] != nil {
		temp := make(map[int]int64)
		for key, val := range m["RetCodes"].(map[string]interface{}) {
			k, _ := strconv.Atoi(key)
			v, _ := strconv.ParseInt(val.(string), 10, 64)
			temp[k] = v
		}
		m["RetCodes"] = temp
	}

	if m["Sizes"] != nil {
		h, err := strconv.ParseInt(m["Sizes"].(map[string]interface{})["Count"].(string), 10, 64)
		if err != nil {
			return nil, err
		}
		m["Sizes"].(map[string]interface{})["Count"] = h
	}

	outTemp, _ := json.Marshal(m)

	return outTemp, nil
}

func formatToNs(s *duration.Duration) string {
	ss := s.AsDuration().String()
	if strings.Contains(ss, "ms") {
		sep := strings.Split(ss, "m")
		f, _ := strconv.ParseFloat(sep[0], 64)
		f = f / 1000
		st := fmt.Sprintf("%f", f)
		sep[0] = st
		ss = strings.Join(sep, "")
	} else if strings.Contains(ss, "µs") {
		sep := strings.Split(ss, "µ")
		f, _ := strconv.ParseFloat(sep[0], 64)
		f = f / 1000000
		st := fmt.Sprintf("%f", f)
		sep[0] = st
		ss = strings.Join(sep, "")
	} else if strings.Contains(ss, "ns") {
		sep := strings.Split(ss, "n")
		f, _ := strconv.ParseFloat(sep[0], 64)
		f = f / 10000000
		st := fmt.Sprintf("%f", f)
		sep[0] = st
		ss = strings.Join(sep, "")
	}
	return ss
}
