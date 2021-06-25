# | <img src="docs/images/logo.png" alt="Fury Logo" width="18" height="25"> |  http-status-check |

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/sighupio/http-status-check)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sighupio/http-status-check)
![GitHub](https://img.shields.io/github/license/sighupio/http-status-check)
[![Go Report Card](https://goreportcard.com/badge/github.com/sighupio/http-status-check)](https://goreportcard.com/report/github.com/sighupio/http-status-check)

## Table of Contents

- [| <img src="docs/images/logo.png" alt="Fury Logo" width="18" height="25"> |  http-status-check |](#----http-status-check-)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Usage](#usage)
      - [Usage from inside a docker image](#usage-from-inside-a-docker-image)
    - [Deploy in a cluster as a Job](#deploy-in-a-cluster-as-a-job)
    - [Examples](#examples)
  - [Developer Guide](#developer-guide)
  - [License](#license)

## Overview

http-status-check is a tool that TBD

The tool can be used to TBD:

- As a standalone tool to connect to the local Kubernetes cluster.
- As a standalone tool to connect to a remote Kubernetes cluster.
- As a Kubernetes Job accessing services in the same cluster authorized using RBAC.

## Getting Started

### Installation

The simplest way to use the tool is to install the binary from the source repo
as follows:

* Using Go get

```sh
$ go get -u github.com/sighupio/http-status-check/cmd/http-status-check
#
```

You should find the CLI installed in the `$GOPATH`. From this point mentioned as
`http-status-check`

### Usage

The basic usage info of the tool can be seen by using the following command:

```sh
$ http-status-check -h
#
```

> [Refer this extended documentation on CLI usage for more](./cmd/http-status-check/README.md)

#### Usage from inside a docker image

There is a [Dockerfile bundled](./build/builder/Dockerfile) with this repo which
can be used to build a Docker image and that can be used to run the binary. To
build docker image one can use the make rule `build`. You can read more about
Makefile in [CONTRIBUTING.md](./CONTRIBUTING.md). To build the image:

``` sh
$ make build
# The docker image will be created by the name http-status-check:local-build
```

The above image can be run exactly the way the CLI is used like shown by the
code block below:

``` sh
$ docker run -it http-status-check:local-build ./http-status-check -h
http-status-check TBD

Usage:
  http-status-check [flags]

Flags:
  -h, --help                help for http-status-check

$ docker run -v .kube/:/root/.kube/ -it http-status-check:local-build ./http-status-check \
                                            --flag-1=1 --flag-2=2
```

### Deploy in a cluster as a Job

As a part of the health check toolkit of Fury Intelligence Platform, this tool
is primary built to work as a Kubernetes Job or CronJob monitoring if TBD.
In order to do so, our preferred way of deployment is by using a
kustomization file that deploys the RBAC policy letting the job look into the
services and enpoints in a namespace and the cron job itself.

We ship the deployment files under the `deployment` directory. To
understand the usage of these files in detail head over to the [README of the
directory](./deployments/).

### Examples

TBD of example deployments that are being monitored by our
http-status-check tool can be found in the `examples` directory. Follow
the usage information in the [corresponding README](./examples/) for more info.

## Developer Guide

To set the code up locally, build, run tests, etc. Please refer the
[contributor's guide](./CONTRIBUTING.md).

## License

Check the [License here](LICENSE)
