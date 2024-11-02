# Build Stage
FROM golang:1.23.2 as build

WORKDIR /app
COPY . .

# Download and cache Go dependencies
RUN go mod download

# Build the application from the main.go file in the /cmd/api directory
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /go/bin/app ./cmd/app

# Production Stage
FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /

ENTRYPOINT ["/app"]