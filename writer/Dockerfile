# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20 AS build-stage

WORKDIR go/src/app

COPY writer/go.mod writer/go.sum ./
RUN go mod download

COPY ./writer ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /shortify-writer

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /shortify-writer /shortify-writer

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/shortify-writer"]