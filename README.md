# GetNighthawk

- Project site: https://getnighthawk.dev
- Project doc: https://docs.google.com/document/d/1lHfMo4iIx2WXFZIspfHyxTsPR1T63_2IV5NUkgxoo0w/edit# 
- Project Slack: http://slack.layer5.io
## What is Nighthawk?
Nighthawk is a Layer 7 (HTTP/HTTPS/HTTP2) performance characterization tool. Nighthawk is Envoy’s load generator and is written in C++. 

## Nighthawk and Meshery
Meshery integrates Nighthawk as one of (currently) three choices of load generator for characterizing and managing the performance of service meshes and their workloads. 

## Why GetNighthawk?
Nighthawk is growing in popularity, but the core project only builds to one architecture / one Docker image. Recently, Nighthawk is being improved so that it can be horizontally scalable - such that multiple instances will be cognizant of one another and able to coordinate amongst each other. Nighthawk is a subproject of Envoy. Nighthawk is growing in popularity with Google, Red Hat, and AWS are investing into it. Istio is considering switching from Fortio to Nighthawk.

---
---

# Go-Nighthawk

Nighthawk adapter to run service mesh load tests with Meshery.

## Load Generators in Meshery

Users may prefer to use one load generator over the next given the difference of capabilities between load generators, so Meshery provides a load generator interface (a gRPC interface) behind which a load generator can be implemented. Meshery provides users with choice of which load generator they prefer to use for a given performance test. Users may set their configure their own preference of load generator different that the default load generator.

### What function do load generators in Meshery provide?

Load generators will provide the capability to run load tests from Meshery. As of today the load generators are embedded as libraries in Meshery and Meshery invokes the load generators APIs with the right load test options to run the load test. At the moment, Meshery has support for HTTP load generators. Support for GRPC and TCP load testing is on the roadmap. Meshery has functional integration with `fortio`, `wrk2`, and `nighthawk`.

### Why support multiple load generators?

Different use cases and different opinions call for different approaches to statistical analysis of the performance results. For example, wrk2 accounts for a concept called Coordinated Omission.

### Which are currently supported?

`fortio` - Fortio load testing library, command line tool, advanced echo server and web UI in go (golang). Allows to specify a set query-per-second load and record latency histograms and other useful stats.
`wrk2` - A constant throughput, correct latency recording variant of wrk.
`nighthawk` - Enables users to run distributed performance tests to better mimic real-world, distributed systems scenarios.




<div>&nbsp;</div>

## Join the service mesh community!

<a name="contributing"></a><a name="community"></a>
Our projects are community-built and welcome collaboration. 👍 Be sure to see the <a href="https://docs.google.com/document/d/17OPtDE_rdnPQxmk2Kauhm3GwXF1R5dZ3Cj8qZLKdo5E/edit">Layer5 Community Welcome Guide</a> for a tour of resources available to you and jump into our <a href="http://slack.layer5.io">Slack</a>!

<p style="clear:both;">
<a href ="https://layer5.io/community/meshmates"><img alt="MeshMates" src=".github/readme/images/Layer5-MeshMentors.png" style="margin-right:10px; margin-bottom:7px;" width="28%" align="left" /></a>
<h3>Find your MeshMate</h3>

<p>MeshMates are experienced Layer5 community members, who will help you learn your way around, discover live projects and expand your community network. 
Become a <b>Meshtee</b> today!</p>

Find out more on the <a href="https://layer5.io/community">Layer5 community</a>. <br />
<br /><br /><br /><br />
</p>

<div>&nbsp;</div>

<a href="https://meshery.io/community"><img alt="Layer5 Service Mesh Community" src=".github/readme/images//slack-128.png" style="margin-left:10px;padding-top:5px;" width="110px" align="right" /></a>

<a href="http://slack.layer5.io"><img alt="Layer5 Service Mesh Community" src=".github/readme/images//community.svg" style="margin-right:8px;padding-top:5px;" width="140px" align="left" /></a>

<p>
✔️ <em><strong>Join</strong></em> any or all of the weekly meetings on <a href="https://calendar.google.com/calendar/b/1?cid=bGF5ZXI1LmlvX2VoMmFhOWRwZjFnNDBlbHZvYzc2MmpucGhzQGdyb3VwLmNhbGVuZGFyLmdvb2dsZS5jb20">community calendar</a>.<br />
✔️ <em><strong>Watch</strong></em> community <a href="https://www.youtube.com/playlist?list=PL3A-A6hPO2IMPPqVjuzgqNU5xwnFFn3n0">meeting recordings</a>.<br />
✔️ <em><strong>Access</strong></em> the <a href="https://drive.google.com/drive/u/4/folders/0ABH8aabN4WAKUk9PVA">community drive</a>.<br />
</p>
<p align="center">
<i>Not sure where to start?</i> Grab an open issue with the <a href="https://github.com/issues?utf8=✓&q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Alayer5io+label%3A%22help+wanted%22+">help-wanted label</a>.
</p>
