---
layout: docs
title: Using Nighthawk
section: "Getting Started"
---

# Using the Nighthawk CLI

```
bazel-bin/nighthawk_client --help
```

```
USAGE:

bazel-bin/nighthawk_client  [--latency-response-header-name <string>]
[--stats-flush-interval <uint32_t>]
[--stats-sinks <string>] ...
[--no-duration] [--simple-warmup]
[--request-source-plugin-config <string>]
[--request-source <uri format>] [--label
<string>] ... [--multi-target-use-https]
[--multi-target-path <string>]
[--multi-target-endpoint <string>] ...
[--experimental-h2-use-multiple-connections]
[--nighthawk-service <uri format>]
[--jitter-uniform <duration>] [--open-loop]
[--experimental-h1-connection-reuse-strategy
<mru|lru>] [--failure-predicate <string,
uint64_t>] ... [--termination-predicate
<string, uint64_t>] ... [--trace <uri
format>] [--sequencer-idle-strategy <spin
|poll|sleep>] [--max-concurrent-streams
<uint32_t>] [--max-requests-per-connection
<uint32_t>] [--max-active-requests
<uint32_t>] [--max-pending-requests
<uint32_t>] [--transport-socket <string>]
[--tls-context <string>]
[--request-body-size <uint32_t>]
[--request-header <string>] ...
[--request-method <GET|HEAD|POST|PUT|DELETE
|CONNECT|OPTIONS|TRACE>] [--address-family
<auto|v4|v6>] [--burst-size <uint32_t>]
[--prefetch-connections] [--output-format
<json|human|yaml|dotted|fortio
|experimental_fortio_pedantic>] [-v <trace
|debug|info|warn|error|critical>]
[--concurrency <string>] [-p <http1|http2
|http3>] [--h2] [--timeout <uint32_t>]
[--duration <uint32_t>] [--connections
<uint32_t>] [--rps <uint32_t>] [--]
[--version] [-h] <uri format>


Where:

--latency-response-header-name <string>
Set an optional header name that will be returned in responses, whose
values will be tracked in a latency histogram if set. Can be used in
tandem with the test server's response option
"emit_previous_request_delta_in_response_header" to record elapsed
time between request arrivals. Default: ""

--stats-flush-interval <uint32_t>
Time interval (in seconds) between flushes to configured stats sinks.
Default: 5.

--stats-sinks <string>  (accepted multiple times)
Stats sinks (in json or compact yaml) where Nighthawk metrics will be
flushed. This argument is intended to be specified multiple times.
Example (json): {name:"envoy.stat_sinks.statsd"
,typed_config:{"@type":"type.googleapis.com/envoy.config.metrics.v3.St
atsdSink",tcp_cluster_name:"statsd"}}

--no-duration
Request infinite execution. Note that the default failure predicates
will still be added. Mutually exclusive with --duration.

--simple-warmup
Perform a simple single warmup request (per worker) before starting
execution. Note that this will be reflected in the counters that
Nighthawk writes to the output. Default is false.

--request-source-plugin-config <string>
[Request
Source](https://github.com/envoyproxy/nighthawk/blob/main/docs/root/ov
erview.md#requestsource) plugin configuration in json or compact yaml.
Mutually exclusive with --request-source. Example (json):
{name:"nighthawk.stub-request-source-plugin"
,typed_config:{"@type":"type.googleapis.com/nighthawk.request_source.S
tubPluginConfig",test_value:"3"}}

--request-source <uri format>
Remote gRPC source that will deliver to-be-replayed traffic. Each
worker will separately connect to this source. For example
grpc://127.0.0.1:8443/. Mutually exclusive with
--request_source_plugin_config.

--label <string>  (accepted multiple times)
Label. Allows specifying multiple labels which will be persisted in
structured output formats.

--multi-target-use-https
Use HTTPS to connect to the target endpoints. Otherwise HTTP is used.
Mutually exclusive with providing a URI.

--multi-target-path <string>
The single absolute path Nighthawk should request from each target
endpoint. Required when using --multi-target-endpoint. Mutually
exclusive with providing a URI.

--multi-target-endpoint <string>  (accepted multiple times)
Target endpoint in the form IPv4:port, [IPv6]:port, or DNS:port. This
argument is intended to be specified multiple times. Nighthawk will
spread traffic across all endpoints with round robin distribution.
Mutually exclusive with providing a URI.

--experimental-h2-use-multiple-connections
DO NOT USE: This option is deprecated, if this behavior is desired,
set --max-concurrent-streams to one instead.

--nighthawk-service <uri format>
Nighthawk service uri. Example: grpc://localhost:8843/. Default is
empty.

--jitter-uniform <duration>
Add uniformly distributed absolute request-release timing jitter. For
example, to add 10 us of jitter, specify .00001s. Default is empty /
no uniform jitter.

--open-loop
Enable open loop mode. When enabled, the benchmark client will not
provide backpressure when resource limits are hit.

--experimental-h1-connection-reuse-strategy <mru|lru>
Choose picking the most recently used, or least-recently-used
connections for re-use.(default: mru). WARNING: this option is
experimental and may be removed or changed in the future!

--failure-predicate <string, uint64_t>  (accepted multiple times)
Failure predicate. Allows specifying a counter name plus threshold
value for failing execution. Defaults to not tolerating error status
codes and connection errors.

--termination-predicate <string, uint64_t>  (accepted multiple times)
Termination predicate. Allows specifying a counter name plus threshold
value for terminating execution.

--trace <uri format>
Trace uri. Example: zipkin://localhost:9411/api/v2/spans. Default is
empty.

--sequencer-idle-strategy <spin|poll|sleep>
Choose between using a busy spin/yield loop or have the thread poll or
sleep while waiting for the next scheduled request (default: spin).

--max-concurrent-streams <uint32_t>
Max concurrent streams allowed on one HTTP/2 or HTTP/3 connection.
Does not apply to HTTP/1. (default: 2147483647).

--max-requests-per-connection <uint32_t>
Max requests per connection (default: 4294937295).

--max-active-requests <uint32_t>
The maximum allowed number of concurrently active requests. HTTP/2
only. (default: 100).

--max-pending-requests <uint32_t>
Max pending requests (default: 0, no client side queuing. Specifying
any other value will allow client-side queuing of requests).

--transport-socket <string>
Transport socket configuration in json or compact yaml. Mutually
exclusive with --tls-context. Example (json):
{name:"envoy.transport_sockets.tls"
,typed_config:{"@type":"type.googleapis.com/envoy.extensions.transport
_sockets.tls.v3.UpstreamTlsContext"
,common_tls_context:{tls_params:{cipher_suites:["-ALL:ECDHE-RSA-AES128
-SHA"]}}}}

--tls-context <string>
DEPRECATED, use --transport-socket instead. Tls context configuration
in json or compact yaml. Mutually exclusive with --transport-socket.
Example (json):
{common_tls_context:{tls_params:{cipher_suites:["-ALL:ECDHE-RSA-AES128
-SHA"]}}}

--request-body-size <uint32_t>
Size of the request body to send. NH will send a number of consecutive
'a' characters equal to the number specified here. (default: 0, no
data).

--request-header <string>  (accepted multiple times)
Raw request headers in the format of 'name: value' pairs. This
argument may specified multiple times.

--request-method <GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE>
Request method used when sending requests. The default is 'GET'.

--address-family <auto|v4|v6>
Network address family preference. Possible values: [auto, v4, v6].
The default output format is 'AUTO'.

--burst-size <uint32_t>
Release requests in bursts of the specified size (default: 0).

--prefetch-connections
Use proactive connection prefetching (HTTP/1 only).

--output-format <json|human|yaml|dotted|fortio
|experimental_fortio_pedantic>
Output format. Possible values: {"json", "human", "yaml", "dotted",
"fortio", "experimental_fortio_pedantic"}. The default output format
is 'human'.

-v <trace|debug|info|warn|error|critical>,  --verbosity <trace|debug
|info|warn|error|critical>
Verbosity of the output. Possible values: [trace, debug, info, warn,
error, critical]. The default level is 'info'.

--concurrency <string>
The number of concurrent event loops that should be used. Specify
'auto' to let Nighthawk leverage all vCPUs that have affinity to the
Nighthawk process. Note that increasing this results in an effective
load multiplier combined with the configured --rps and --connections
values. Default: 1.

-p <http1|http2|http3>,  --protocol <http1|http2|http3>
The protocol to encapsulate requests in. Possible values: [http1,
http2, http3]. The default protocol is 'http1' when neither of --h2 or
--protocol is used. Mutually exclusive with --h2.

--h2
DEPRECATED, use --protocol instead. Encapsulate requests in HTTP/2.
Mutually exclusive with --protocol. Requests are encapsulated in
HTTP/1 by default when neither of --h2 or --protocol is used.

--timeout <uint32_t>
Connection connect timeout period in seconds. Default: 30.

--duration <uint32_t>
The number of seconds that the test should run. Default: 5. Mutually
exclusive with --no-duration.

--connections <uint32_t>
The maximum allowed number of concurrent connections per event loop.
HTTP/1 only. Default: 100.

--rps <uint32_t>
The target requests-per-second rate. Default: 5.

--,  --ignore_rest
Ignores the rest of the labeled arguments following this flag.

--version
Displays version information and exits.

-h,  --help
Displays usage information and exits.

<uri format>
URI to benchmark. http:// and https:// are supported, but in case of
https no certificates are validated. Provide a URI when you need to
benchmark a single endpoint. For multiple endpoints, set
--multi-target-* instead.


```
<br/>

# Using Nighthawk gRPC Service

The gRPC service can be used to start a server which is able to perform back-to-back benchmark runs upon request. The service interface definition [can be found here](https://github.com/envoyproxy/nighthawk/blob/59a37568783272a6438b5697277d4e56aa16ebbe/api/client/service.proto).

```
bazel-bin/nighthawk_service --help
```

```
USAGE:

bazel-bin/nighthawk_service  [--service <traffic-generator-service
|dummy-request-source>]
[--listener-address-file <>] [--listen
<address:port>] [--] [--version] [-h]


Where:

--service <traffic-generator-service|dummy-request-source>
Specifies which service to run. Default 'traffic-generator-service'.

--listener-address-file <>
Location where the service will write the final address:port on which
the Nighthawk grpc service listens. Default empty.

--listen <address:port>
The address:port on which the Nighthawk gRPC service should listen.
Default: 0.0.0.0:8443.

--,  --ignore_rest
Ignores the rest of the labeled arguments following this flag.

--version
Displays version information and exits.

-h,  --help
Displays usage information and exits.


L7 (HTTP/HTTPS/HTTP2) performance characterization tool.
```
<br/>

# Nighthawk Output Transform Utility

Nighthawk comes with a tool to transform its json output to its other supported output formats.

```
bazel-bin/nighthawk_output_transform --help
```

```
USAGE:

bazel-bin/nighthawk_output_transform  --output-format <json|human|yaml
|dotted|fortio
|experimental_fortio_pedantic> [--]
[--version] [-h]


Where:

--output-format <json|human|yaml|dotted|fortio
|experimental_fortio_pedantic>
(required)  Output format. Possible values: {"json", "human", "yaml",
"dotted", "fortio", "experimental_fortio_pedantic"}.

--,  --ignore_rest
Ignores the rest of the labeled arguments following this flag.

--version
Displays version information and exits.

-h,  --help
Displays usage information and exits.


L7 (HTTP/HTTPS/HTTP2) performance characterization transformation tool.
```

See [Nighthawk README](https://github.com/envoyproxy/nighthawk) for more info.

{% include related-discussions.html tag="nighthawk" %}
