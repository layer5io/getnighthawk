---
layout: docs
title: Building Nighthawk
section: "Getting Started"
---

# Building Nighthawk

While GetNighthawk distribution are a convenient way of using Nighthawk, you can also build Nighthawk and run it from source.

## Clone the Nighthawk repo

Using GitHub CLI:

```
gh repo clone envoyproxy/nighthawk
```

Using HTTPS:

```
git clone https://github.com/envoyproxy/nighthawk.git
```

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

## Building Nighthawk

Run:

```
bazel build -c opt //:nighthawk
```
