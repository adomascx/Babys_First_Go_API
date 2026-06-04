# Skelbiu API

A Go REST API that serves scraped skelbiu.lt listings.

## Table of Contents

- [Skelbiu API](#skelbiu-api)
  - [Table of Contents](#table-of-contents)
  - [Project Description](#project-description)
    - [Tech Stack](#tech-stack)
    - [Features](#features)
    - [Project Structure](#project-structure)
  - [Getting started](#getting-started)
  - [Design Notes](#design-notes)
    - [Process Flow Diagram](#process-flow-diagram)
    - [Planned features](#planned-features)

## Project Description

This project implements an API representation of public Skelbiu.lt listings, which is particularly useful for listing analyses via LLMs (e.g. "best deal" finder).
Aside from that, it serves as a personal training ground for learning Go's networking and API development fundamentals as well as best practices in production environments.

### Tech Stack

- **Go** - Main language for the backend
- **Google Cloud Run** - Hosting the server
- **Postman** - API testing
- **Air** - Live reload for development

### Features

- Full OpenAPI specifications
- Integrated Postman `Local` and `Prod` environments
- Ready-to-use Air and Postman configs
- Unit testing/benchmarks for most functionality

### Project Structure

- `cmd/api` - Application entry point
- `api` - OpenAPI spec (`openapi.yaml`)
- `internal` - Internal Go packages and implementation details
  - `internal/handler` - HTTP handlers and routing
  - `internal/model` - Domain models, listing definitions, and query helpers
  - `internal/scraper` - Scraper implementation and tests
- `postman` - Postman collections and environments for testing <!-- outdated, update once API is finalized -->

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
set PORT=8080
go run ./cmd/api
```

## Design Notes

- LLMs were only used for answering technical questions, in order to necessitate personal learning. For instance, all documentation was hand written by me :p
- Go's stdlib was used as much as possible, with the exception of a scraping framework ([Colly](https://github.com/gocolly/colly)) and a TLS impersonator ([uTLS](https://github.com/refraction-networking/utls))
- The project's file structure was build from the ground up to be standardized
- A lack of caching/storage was a deliberate choice, as this avoids possible violations of EU's GDPR laws

### Process Flow Diagram

Our greatest scientists have created this wonderful process flow diagram to illustrate the API's architecture and data flow:

![Flowchart](logic.png)

### Planned features

- Parallelization of page retrieval via Goroutines
- Opt-in listing caching w/ PostgreSQL
- Front-end API demo
- Front-end "best deal finder" application
