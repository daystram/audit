# :satellite: Audit

[![Github Actions](https://github.com/daystram/audit/actions/workflows/push.yml/badge.svg)](https://github.com/daystram/audit/actions/workflows/push.yml)
[![Codecov](https://codecov.io/gh/daystram/audit/branch/master/graph/badge.svg?token=VI1CYFQ50N)](https://codecov.io/gh/daystram/audit)
[![Docker Pulls](https://img.shields.io/docker/pulls/daystram/audit)](https://hub.docker.com/r/daystram/audit)
[![MIT License](https://img.shields.io/github/license/daystram/audit)](https://github.com/daystram/audit/blob/master/LICENSE)

**Audit** is a service monitoring and incident management service.

## Features

- WIP: Service availability/uptime monitoring
- WIP: Response time logging
- WIP: HTTP service tracking
- WIP: TCP/UDP service tracking
- WIP: Incident log management
- WIP: Multi-region distributed tracker agents

## Services

The application comes in two parts:

| Name      |  Code Name  | Stack                                                                                                                                                                               |
| --------- | :---------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Back-end  | `audit-be` | [Go](https://golang.org/), [Gin](https://github.com/gin-gonic/gin) + [Gorm](https://github.com/go-gorm/gorm), [gRPC](https://grpc.io), [InfluxDB](https://www.influxdata.com/), [PostgreSQL](https://www.postgresql.org/) |
| Tracker  | `audit-tr` | [Go](https://golang.org/), [gRPC](https://grpc.io) |
| Front-end | `audit-fe` | [TypeScript](https://www.typescriptlang.org/), [Vue.js](https://vuejs.org/)                                                                                                         |

## Develop

### Generate Protobuf and Mocks
The project relies on some parts of the code being generated, notably the protocol buffers and the mocks used for testing. To generate these files, install [protoc](https://grpc.io/docs/protoc-installation``) and [mockgen](https://github.com/golang/mock). Then run the target on the Makefile.

Install the Go and Go gRPC plugin for protoc:
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

Generate the files:
```shell
$ make all
```

### audit-be

`audit-be` uses [Go Modules](https://blog.golang.org/using-go-modules) module/dependency manager, hence at least Go 1.11 is required. To ease development, [comstrek/air](https://github.com/cosmtrek/air) is used to live-reload the application. Install the tool as documented.

To begin developing, simply enter the sub-directory and run the development server:

```shell
$ cd audit-be
$ go mod tidy
$ air
```

### audit-tr

Populate `.env.development` with the required credentials.

To begin developing, simply enter the sub-director and run the development server:

```shell
$ cd audit-fe
$ yarn
$ yarn serve
```

### audit-fe

Populate `.env.development` with the required credentials.

To begin developing, simply enter the sub-directory and run the development server:

```shell
$ cd audit-fe
$ yarn
$ yarn serve
```

## Deploy

Both `audit-be` and `audit-fe` are containerized and pushed to [Docker Hub](https://hub.docker.com/r/daystram/audit). They are tagged based on their application name and version, e.g. `daystram/audit:be` or `daystram/audit:be-v1.1.0`.

To run `audit-be`, run the following:

```shell
$ docker run --name audit-be --env-file ./.env -p 8080:8080 -d daystram/audit:be
```

And `audit-fe` as follows:

```shell
$ docker run --name audit-fe -p 80:80 -d daystram/audit:fe
```

### Dependencies

The following are required for `audit-be` to function properly:

- InfluxDB
- PostgreSQL

Their credentials must be provided in the configuration file.

### Helm Chart

> WIP

### Docker Compose

> WIP

## License

This project is licensed under the [MIT License](./LICENSE).
