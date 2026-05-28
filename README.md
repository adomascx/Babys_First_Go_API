# Baby's first Go REST API

A Proof-of-Concept Go REST API

## Table of Contents

- [Baby's first Go REST API](#babys-first-go-REST API)
  - [Table of Contents](#table-of-contents)
  - [Project Description](#project-description)
    - [Tech Stack](#tech-stack)
    - [Features](#features)
    - [Project Structure](#project-structure)
  - [Getting started](#getting-started)
  - [Development Notes](#development-notes)

## Project Description

This project is meant as a training ground for learning Go's networking as well as API development fundamentals.
Adherence to both Golang and API standards was prioritized, whilst trying not to "overengineer" what is basically an example project.
So far, the goal is only to create a functional REST API, with specialization to come afterwards.

### Tech Stack

- **Go** - Main language for the backend
- **Google Cloud Run** - Hosting the server
- **Postman** - API testing
- **Air** - Live reload for development

### Features

- Integrated Postman `Local` and `Prod` environments
- Ready-to-use Air and Postman configs
- More to come ;)

### Project Structure

- `cmd/api` - Application entry point
- `internal` - Internal packages and future implementation details
- `postman` - Collection files and environments for API testing

## Getting started

The live server can be run with *air*:

```cmd
air
```

If you want to run it manually instead, use:

```cmd
go run ./cmd/api
```

To use a custom port, set `PORT` first:

```cmd
set PORT=8081
go run ./cmd/api
```

## Development Notes

This is still an early-stage project. The README will likely grow once the REST API layer, tests, and request flow settle down.

### Planned features

- Parallelization of page retrieval via Goroutines
- Front-end API demo
