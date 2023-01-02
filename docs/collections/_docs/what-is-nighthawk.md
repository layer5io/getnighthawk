---
layout: docs
title: What is Nighthawk?
section: "Overview"
---

# What is Nighthawk?

Nighthawk is a versatile HTTP load testing tool built out of a need to drill HTTP services with a constant request rate or with an adaptive request rate. Layer5 offers a custom distribution of Nighthawk with intelligent adaptive load controllers to automatically identify optimatal configurations for your service mesh deployment. As a Layer 7 performance characterization tool supporting HTTP/HTTPS/HTTP2, Nighthawk is Meshery's (and Envoy's) load generator and is written in C++.


<!-- GetNighthawk was built with the goal to make it easy to use Nighthawk. -->

You can download generally available distributions of Nighthawk under different architectures and platforms.
 <!-- also providing easy-to-use tooling for installation and operation. This involves creating distributions of Nighthawk and to build up existing tooling. -->

<img src="/assets/images/screenshots/Nighthawk+Meshery.svg" style="width:90%;padding:4%;align:center" />

Centric to the advancement of Nighthawk is the Meshery and Service Mesh Performance projects, which enable Nighthawk's standards-based distributed performance management. The intersection of these projects allow researchers and users to conveniently identify the optimal service mesh configuration while considering their specific environment, application and load.Meshery orchestrates multiple instances of Nighthawk (horizontal scaling) and provides an easy to use interface for Nighthawk's adaptive load controller capability.

Nighthawk also enables Nighthawk adoption by delivering trusted, certified builds, distributed via popular package managers like apt, yum, Homebrew and platforms including Docker and Kubernetes.

<!-- Nighthawk also bridges the gap between C++ code in Nighthawk and the language of the cloud, Golang. -->
First-class support for Nighthawk in [Service Mesh Patterns](https://github.com/service-mesh-patterns) is also available. Nighthawk to be the performance characterization tool that would be used in the 30 patterns in the [Service Mesh Patterns](https://layer5.io/learn/service-mesh-books/service-mesh-patterns) book.

{% include related-discussions.html tag="nighthawk" %}