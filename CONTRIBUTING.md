# Contribution

Read the following guide about how to contribute to this project and get familiar with.

Project structure based on:
[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Make

This project contains an easy-to-use interface with a `Makefile`. It allows you to pass all the checks the pipeline
will pass on certain events. So it worth get familiar with all the targets. To know what targets are available in
the project just run:

```bash
$ make

 Choose a command to run in http-status-check:

  init            Init the project. GITHUB_PROJECT=demo make init
  drone-init      Init the drone-project. GITHUB_PROJECT=demo GITHUB_TOKEN=123token321 DRONE_TOKEN=tokenhere REGISTRY=registry.sighup.io REGISTRY_USER=robotuser REGISTRY_PASSWORD=thepassword make drone-init
  build           Build the container image
  lint            Run the policeman over the repository
  build-release   Build the release container image
  test            Run unit testing
  license         Check license headers are in-place in all files in the project
  e2e-test        Execute e2e-tests. CLUSTER_VERSION=v1.21.1 make e2e-test
  publish         Publish the container image

```

### init

The `init` target is required to be executed just one time after creating the project from the template.
Once the project is initialized you'll receive a message like:

```bash
$ GITHUB_PROJECT=a make init
Project already initialized with name demo-check
```

You can read more about this target in
[the template GitHub repository](https://github.com/sighupio/http-status-check)

### drone-init

The `drone-init` target is required to be executed just one time after creating the project from the template.
Once the project is initialized you'll receive a message like:

```bash
$ GITHUB_PROJECT=demo GITHUB_TOKEN=123token321 DRONE_TOKEN=tokenhere REGISTRY=registry.sighup.io REGISTRY_USER=robotuser REGISTRY_PASSWORD=thepassword make init
Project already initialized with name demo-check
```

It automates the configuration of the drone project with a single command, it requires a privilege token to run it.
If you don't have it, read more about how to set up it manually
[in the docs](https://github.com/sighupio/http-status-check)

### build

This project has been developed using some awesome features from `docker`. With the same
[`Dockerfile`](build/builder/Dockerfile) you'll be able to `build` the project and [`test`](#test) it locally among
other automation.

If you are adding more directories than the default ones to the project, you'll need to specify them in the
[`Dockerfile`](build/builder/Dockerfile):

```Dockerfile
.
..
...
RUN go test ./...

FROM golang:1.16 as builder

RUN mkdir /app
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd cmd
COPY pkg pkg
COPY internal internal

WORKDIR /app/cmd/http-status-check
RUN go build

FROM debian:buster as release
...
..
.
```

Same with the build command, if you need to specify some golang flags to the `build` command, this is the right place
where to customize it.

### lint

This `make` target makes easy check if the (company-wide) linter pass or not. This is a basic requirement in order to
build and release a project at SIGHUP.io

The `lint` target uses the
[`policeman` *(with actually it is a wrapper around GitHub super-linter)*](https://github.com/sighupio/fury-kubernetes-policeman)
to check if everything is fine.

This project contains [a couple of additional linter rules](.rules) that has been tested in a real environment.

### build-release

`build-release` target creates a ready to run container image. It gets the `build` output binary, adding to a
`debian:buster` container image base.

### test

The `test` target executes the command defined in the `tester` [`Dockerfile`](build/builder/Dockerfile)
target in order to unit test the application.

Again, if you need to add a different command to test the application, you can do it by modifying the right
chunk of the [`Dockerfile`](build/builder/Dockerfile):

```Dockerfile
.
..
...
RUN addlicense -c "SIGHUP s.r.l" -v -l bsd --check .

FROM golang:1.16 as tester

RUN mkdir /app
WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd cmd
COPY pkg pkg
COPY internal internal

RUN go test ./...

FROM golang:1.16 as builder

RUN mkdir /app
...
..
.
```

Currently it executes `go test ./...` but feel free to modify with the right command.

### license

This make target checks for the license header in all the files. If you add new files that doesn't contain the
required header, you'll need to add them manually by:

```bash
go get -u github.com/google/addlicense
addlicense -c "SIGHUP s.r.l" -v -l bsd .
```

Then, you'll be able to re-check the `license` again:

```bash
make license
```

### e2e-test

`e2e-test` target has been designed to create a local *(and easy)* cluster using
[`kind`](https://github.com/kubernetes-sigs/kind). Then runs e2e tests on top of the cluster.
Test runs using [`bats`](https://github.com/bats-core/bats-core). Some basic tests have been designed in order to
make simple extend with the right e2e-tests.

Add your own tests [here](scripts/e2e/tests.sh).

### publish

The `publish` make target is used only to build the container image (`make build-release`) then push to the registry.
In order to use this target you may need to specify a set of variables:

- **`REGISTRY`**: Registry endpoint where the resulting container image will be pushed.
- **`REGISTRY_USER`**: User of the `REGISTRY`
- **`REGISTRY_PASSWORD`**: Password for the `REGISTRY_USER` @ `REGISTRY`

```bash
REGISTRY_USER=angelbarrera92 IMAGE_TAG=latest IMAGE_NAME=sighupio/http-status-check REGISTRY_PASSWORD=supersuperrarepasswordeh REGISTRY=registry.sighup.io make publish
```

## Pipeline

This project pipeline executes the `make` targets mentioned above.

It executes on every action to the Git Repository:

- `license`. It checks for the license headers.
- `lint`. It checks for common linting rules.
- `test`. It runs the tests of the project.
- `build`. This one executes the build of the project.
- `e2e-test`. Executes the e2e tests of the project.

Then, on every event **except for a PR**:

- `publish`: It creates a container image with the `unstable` tag.

And finally, if the repo got a new `tag`:

- `release`: Creates a new GitHub release when a new `tag` is pushed to the git repo.
It uses the `docs/releases/${DRONE_TAG}.md`. It has to be present before creating the tag.


## Releases

This repository has been designed to follow trunk-based development. Then to create a new
release, you need to create a file describing the release with the name as the name of the
release.

You can see an example here: `docs/releases/v0.1.0.md`.

Then, in order to release it:

```bash
git tag v0.1.0
git push --tags
```
