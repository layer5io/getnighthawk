---
layout: docs
title: Building Nighthawk
section: "Getting Started"
---

# Building Nighthawk

While GetNighthawk distribution are a convenient way of using Nighthawk, you can also build Nighthawk and run it from source.
&nbsp;

## Prerequisite 

#### Installing Bazelisk as Bazel

It is recommended to use [Bazelisk](https://github.com/bazelbuild/bazelisk) installed as `bazel`, to avoid Bazel compatibility issues.

On Linux, run the following commands:

```console
sudo wget -O /usr/local/bin/bazel https://github.com/bazelbuild/bazelisk/releases/latest/download/bazelisk-linux-$([ $(uname -m) = "aarch64" ] && echo "arm64" || echo "amd64")
sudo chmod +x /usr/local/bin/bazel
```
&nbsp;

## Clone the Nighthawk repo

Using GitHub CLI:

```
gh repo clone envoyproxy/nighthawk
```

Using HTTPS:

```
git clone https://github.com/envoyproxy/nighthawk.git
```
&nbsp;

## Install Dependencies

On Ubuntu, run the following:

```console
sudo apt-get install \
    autoconf \
    automake \
    cmake \
    curl \
    libtool \
    make \
    ninja-build \
    patch \
    python3-pip \
    unzip \
    virtualenv
```

Install [Golang](https://golang.org/) on your machine. This is required as part of building [BoringSSL](https://boringssl.googlesource.com/boringssl/+/HEAD/BUILDING.md) and also for [Buildifer](https://github.com/bazelbuild/buildtools) which is used for formatting bazel BUILD files.
&nbsp;

## Building Nighthawk

Run:

```
cd nighthawk
bazel build -c opt //:nighthawk
```

<br>
{% include related-discussions.html tag="nighthawk" %}