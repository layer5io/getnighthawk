package nighthawk_client

import (
	"encoding/json"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (opt *CommandLineOptions) UnmarshalJSON(data []byte) error {

	type Duration struct {
		Seconds int64 `json:"seconds"`
		Nanos   int32 `json:"nanos"`
	}

	var set map[string]interface{}

	if err := json.Unmarshal(data, &set); err != nil {
		return err
	}

	o := &struct {
		RequestsPerSecond                     uint32            `json:"requests_per_second"`
		Connections                           uint32            `json:"connections"`
		Timeout                               Duration          `json:"timeout"`
		Concurrency                           string            `json:"concurrency"`
		Verbosity                             int32             `json:"verbosity"`
		OutputFormat                          int32             `json:"output_format"`
		PrefetchConnections                   bool              `json:"prefetch_connections"`
		BurstSize                             uint32            `json:"burst_size"`
		AddressFamily                         int32             `json:"address_family"`
		MaxPendingRequests                    uint32            `json:"max_pending_requests"`
		MaxActiveRequests                     uint32            `json:"max_active_requests"`
		MaxRequestsPerConnection              uint32            `json:"max_requests_per_connection"`
		SequencerIdleStrategy                 int32             `json:"sequencer_idle_strategy"`
		Trace                                 string            `json:"trace"`
		ExperimentalH1ConnectionReuseStrategy int32             `json:"experimental_h1_connection_reuse_strategy"`
		TerminationPredicates                 map[string]uint64 `json:"termination_predicates"`
		FailurePredicates                     map[string]uint64 `json:"failure_predicates"`
		OpenLoop                              bool              `json:"open_loop"`
		JitterUniform                         Duration          `json:"jitter_uniform"`
		NighthawkService                      string            `json:"nighthawk_service"`
		MaxConcurrentStreams                  uint32            `json:"max_concurrent_streams"`
		Labels                                []string          `json:"labels"`
		SimpleWarmup                          bool              `json:"simple_warmup"`
		StatsFlushInterval                    uint32            `json:"stats_flush_interval"`
		LatencyResponseHeaderName             string            `json:"latency_response_header_name"`
		ScheduledStart                        Duration          `json:"scheduled_start"`
		ExecutionId                           string            `json:"execution_id"`
	}{}

	if err := json.Unmarshal(data, &o); err != nil {
		return err
	}

	if set["requests_per_second"] != nil {
		opt.RequestsPerSecond = &wrappers.UInt32Value{Value: o.RequestsPerSecond}
	}

	if set["connections"] != nil {
		opt.Connections = &wrappers.UInt32Value{Value: o.Connections}
	}

	if set["timeout"] != nil {
		opt.Timeout = durationpb.New(time.Duration(o.Timeout.Seconds)*time.Second + time.Duration(o.Timeout.Nanos)*time.Nanosecond)
	}

	if set["concurrency"] != nil {
		opt.Concurrency = &wrappers.StringValue{Value: o.Concurrency}
	}

	if set["verbosity"] != nil {
		verbosity := Verbosity_VerbosityOptions(o.Verbosity)
		opt.Verbosity = &Verbosity{Value: verbosity}
	}

	if set["output_format"] != nil {
		outputFormat := OutputFormat_OutputFormatOptions(o.OutputFormat)
		opt.OutputFormat = &OutputFormat{Value: outputFormat}
	}

	if set["prefetch_connections"] != nil {
		opt.PrefetchConnections = &wrappers.BoolValue{Value: o.PrefetchConnections}
	}

	if set["burst_size"] != nil {
		opt.BurstSize = &wrappers.UInt32Value{Value: o.BurstSize}
	}

	if set["address_family"] != nil {
		addressFamily := AddressFamily_AddressFamilyOptions(o.AddressFamily)
		opt.AddressFamily = &AddressFamily{Value: addressFamily}
	}

	if set["max_pending_requests"] != nil {
		opt.MaxPendingRequests = &wrappers.UInt32Value{Value: o.MaxPendingRequests}
	}

	if set["max_active_requests"] != nil {
		opt.MaxActiveRequests = &wrappers.UInt32Value{Value: o.MaxActiveRequests}
	}

	if set["max_requests_per_connection"] != nil {
		opt.MaxRequestsPerConnection = &wrappers.UInt32Value{Value: o.MaxRequestsPerConnection}
	}

	if set["sequencer_idle_strategy"] != nil {
		sequencerIdleStrategy := SequencerIdleStrategy_SequencerIdleStrategyOptions(o.SequencerIdleStrategy)
		opt.SequencerIdleStrategy = &SequencerIdleStrategy{Value: sequencerIdleStrategy}
	}

	if set["trace"] != nil {
		opt.Trace = &wrappers.StringValue{Value: o.Trace}
	}

	if set["experimental_h1_connection_reuse_strategy"] != nil {
		experimentalH1ConnectionReuseStrategy := H1ConnectionReuseStrategy_H1ConnectionReuseStrategyOptions(o.ExperimentalH1ConnectionReuseStrategy)
		opt.ExperimentalH1ConnectionReuseStrategy = &H1ConnectionReuseStrategy{Value: experimentalH1ConnectionReuseStrategy}
	}

	if set["termination_predicates"] != nil {
		opt.TerminationPredicates = o.TerminationPredicates
	}

	if set["failure_predicates"] != nil {
		opt.FailurePredicates = o.FailurePredicates
	}

	if set["open_loop"] != nil {
		opt.OpenLoop = &wrappers.BoolValue{Value: o.OpenLoop}
	}

	if set["jitter_uniform"] != nil {
		opt.JitterUniform = durationpb.New(time.Duration(o.JitterUniform.Seconds)*time.Second + time.Duration(o.JitterUniform.Nanos)*time.Nanosecond)
	}

	if set["nighthawk_service"] != nil {
		opt.NighthawkService = &wrappers.StringValue{Value: o.NighthawkService}
	}

	if set["max_concurrent_streams"] != nil {
		opt.MaxConcurrentStreams = &wrappers.UInt32Value{Value: o.MaxConcurrentStreams}
	}

	if set["labels"] != nil {
		opt.Labels = o.Labels
	}

	if set["simple_warmup"] != nil {
		opt.SimpleWarmup = &wrappers.BoolValue{Value: o.SimpleWarmup}
	}

	if set["stats_flush_interval"] != nil {
		opt.StatsFlushInterval = &wrappers.UInt32Value{Value: o.StatsFlushInterval}
	}

	if set["latency_response_header_name"] != nil {
		opt.LatencyResponseHeaderName = &wrappers.StringValue{Value: o.LatencyResponseHeaderName}
	}

	if set["scheduled_start"] != nil {
		opt.ScheduledStart = timestamppb.New(time.Unix(o.ScheduledStart.Seconds, int64(o.ScheduledStart.Nanos)))
	}

	if set["execution_id"] != nil {
		opt.ExecutionId = &wrappers.StringValue{Value: o.ExecutionId}
	}

	return nil
}
