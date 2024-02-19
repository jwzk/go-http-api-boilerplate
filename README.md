# Go-HTTP-API-Boilerplate

![Coverage](https://raw.githubusercontent.com/jwzk/go-http-api-boilerplate/badges/.badges/main/coverage.svg)

Golang HTTP API Boilerplate using Hexagonal architectural style and SOLID design principles 🚀

## Requirements

- [Go 1.26.4](https://go.dev/dl/) or [Docker](https://docs.docker.com/get-docker/)

## Development requirements

Install the `mockery` tool:

```bash
$ go install github.com/vektra/mockery/v2@latest
```

Install `golangci-lint`:

```bash
$ go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1
```

Install `pre-commit` (e.g., `brew install pre-commit` or `pip install pre-commit`) and setup hooks:

```bash
$ make setup-hooks
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
$ make up
```

Run tests with:

```bash
$ make test
```

Generate mocks with:

```bash
$ mockery
```

Run Swagger UI locally (without API dependencies) with:

```bash
$ make swagger
```

Access to your local services:

- HTTP API on port `:4100` (http://localhost:4100)
- Swagger API Documentation on port `:4000` (http://localhost:4000)
