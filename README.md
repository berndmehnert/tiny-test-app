# render-gin-tiny

Tiny Go + Gin app intended for deployment to Render as a Web Service.

## Quick start (local)

Build:
    go build -o main .

Run:
    ./main

Or with Docker:
    docker build -t render-gin-tiny .
    docker run -p 8080:8080 render-gin-tiny

## Render notes

- Service type: Web Service (Docker recommended)
- Port: app reads $PORT (default 8080)
- Build: Dockerfile uses golang:1.25.5 multi-stage build
- Start: default CMD in Dockerfile

## Endpoints

- GET /        -> "Hello from Render + Gin!"
- GET /health  -> "OK"

## Environment

- PORT (provided by Render)

## Go

- go 1.25 (go.mod: `go 1.25`)

## License

MIT
