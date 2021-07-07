# :satellite: Audit

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
| Back-end  | `audit-be` | [Go](https://golang.org/), [Gin](https://github.com/gin-gonic/gin) + [Gorm](https://github.com/go-gorm/gorm), [InfluxDB](https://www.influxdata.com/), [PostgreSQL](https://www.postgresql.org/) |
| Front-end | `audit-fe` | [TypeScript](https://www.typescriptlang.org/), [Vue.js](https://vuejs.org/)                                                                                                         |

## Develop

### audit-be

`audit-be` uses [Go Modules](https://blog.golang.org/using-go-modules) module/dependency manager, hence at least Go 1.11 is required. To ease development, [comstrek/air](https://github.com/cosmtrek/air) is used to live-reload the application. Install the tool as documented.

To begin developing, simply enter the sub-directory and run the development server:

```shell
$ cd audit-be
$ go mod tidy
$ air
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
