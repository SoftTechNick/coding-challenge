# Build Stage
FROM golang:1.23.2-alpine as build

# Installiere build tools für CGO und SQLite-Unterstützung
RUN apk add --no-cache gcc musl-dev libc-dev sqlite sqlite-dev

WORKDIR /app
COPY . .

# Download und Cache Go Abhängigkeiten
RUN go mod download

# Baue die Anwendung mit CGO-Unterstützung
RUN CGO_ENABLED=1 GOOS=linux go build -v -o /go/bin/app ./cmd/app

# Production Stage
FROM alpine:latest

# Installiere benötigte Laufzeit-Bibliotheken für SQLite
RUN apk add --no-cache sqlite-libs

# Kopiere die Anwendung aus der Build-Stage
COPY --from=build /go/bin/app /app

ENTRYPOINT ["/app"]
