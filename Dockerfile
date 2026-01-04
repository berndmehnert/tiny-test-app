# build stage
FROM golang:1.25.5-alpine AS build
RUN apk add --no-cache git
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app/main .

# final stage
FROM alpine:3.19
RUN adduser -D app
USER app
WORKDIR /app

# copy binary and templates from build stage
COPY --from=build /app/main /app/main
COPY --from=build /src/templates /app/templates

EXPOSE 8080
CMD ["/app/main"]