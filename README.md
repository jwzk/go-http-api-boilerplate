# Go-HTTP-API-Boilerplate

Golang HTTP API Boilerplate using REST and Hexagonal architectural style and SOLID design principles.

## Requirements

- [Go 1.22](https://go.dev/dl/) or [Docker](https://docs.docker.com/get-docker/)

## Development requirements

Install the `mockery` tool:

```bash
$ go install github.com/vektra/mockery/v2@latest
```

## Project architecture

The project is built on top of this structure:

- **cmd:** Application main
- **internal:** Private application
  - **adapter:** Adapter layer
    - **dao:** Data access object layer
    - **http:** HTTP layer
  - **domain:** Domain business layer
    - **model:** Entity layer
    - **port:** Interface adapter layer
    - **usecase:** Application business rules layer
- **pkg:** External library code

## Local setup

Diplay make commands with:

```bash
$ make help
```

Run Go API with:

```bash
$ make run
```

Run Docker project with:

```bash
$ make start
```

Run tests with:

```bash
$ make test
```

Generate mocks with:

```bash
$ mockery
```

Access to your local service:

- HTTP API on port :4100
