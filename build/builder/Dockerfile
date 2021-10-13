# Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# hadolint ignore=DL3007
FROM quay.io/sighup/policeman:latest as linter

RUN mkdir /app
WORKDIR /app

COPY . .

RUN /entrypoint.sh

FROM golang:1.16 as license

RUN go get -u github.com/google/addlicense

RUN mkdir /app
WORKDIR /app

COPY . .

RUN addlicense -c "SIGHUP s.r.l" -v -l bsd --check .

FROM golang:1.16 as tester

RUN mkdir /app
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd cmd
COPY internal internal

RUN go test -v ./... -cover

FROM golang:1.16 as builder

RUN mkdir /app
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd cmd
COPY internal internal

WORKDIR /app/cmd/http-status-check
RUN go build

FROM golang:1.16 as release

COPY --from=builder /app/cmd/http-status-check/http-status-check /http-status-check
ENTRYPOINT ["/http-status-check"]
