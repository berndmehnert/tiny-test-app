# tiny multi-stage build
FROM golang:1.25.5-alpine AS build
RUN apk add --no-cache git
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/main .

FROM alpine:3.19
RUN adduser -D app
USER app
COPY --from=build /app/main /app/main
EXPOSE 8080
CMD ["/app/main"]
