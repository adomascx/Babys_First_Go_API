# Skelbiu API

A REST API written in Go that returns structured skelbiu.lt listing data.

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

This project implements an API that returns a structured JSON representation of public Skelbiu.lt listings' data.
Aside from that, it serves as a personal training ground for learning Go's networking and API development fundamentals, both in serving and receiving HTTP requests, as well as best practices in production environments.

> [!IMPORTANT]  
> Cloud deployment is considered non-functioning due to Cloudflare blocking standard cloud IPs. As such, any API usage must be done on a local server instance until a self-hosted solution is implemented.

### Tech Stack

- **Go** - Main language for the backend
  - **Colly** - Web scraping framework
  - **uTLS** - TLS browser impersonator
- **Render** - Hosting the production server
- **Postman** - API endpoint testing
- **Air** - Live reload for development

### Features

- Filtering/conversion of listing data into predefinded structs
- Full OpenAPI specifications
- Integrated Postman `Local` and `Prod` environments
- Ready-to-use Air and Postman configs
- Full unit testing/benchmarks suite
- Parallel page retrieval and processing

### Project Structure

- `cmd/api` - Application entry point
- `api` - OpenAPI spec (`openapi.yaml`)
- `internal` - Internal packages and implementation logic
  - `internal/handler` - API HTTP handlers/parsers and routing
  - `internal/model` - Class definitions and methods
  - `internal/scraper` - Scraper implementation
- `postman` - Postman collections and environments for testing

## Getting started

The live development server can be run with *air*:

```cmd
air
```

If you want to run manually instead, use:

```cmd
go run ./cmd/api
```

Change:

```cmd
go run ./cmd/api \
    PORT=8081
```

## Design Notes

- LLMs were only used for answering technical questions and setup creation, in order to necessitate personal learning. For instance, all documentation was hand written by me :p
- Go's stdlib was used as much as possible, with the exception of a scraping framework ([Colly](https://github.com/gocolly/colly)) and a TLS client ([uTLS](https://github.com/refraction-networking/utls))
- The project's file structure was build from the ground up to follow standard practices
- A lack of caching/storage was a deliberate choice to fully comply with GDPR

### Process Flow Diagram

Our greatest scientists have created this wonderful process flow chart to illustrate the API's architecture and data flow:

![Cool and awesome flowchart](logic.png)

### Planned features

- Local server deployment
- Front-end API demo
- Opt-in listing caching w/ Redis
- Edge-case handling (Network interruptions, graceful shutdown)
- Rolling release system
