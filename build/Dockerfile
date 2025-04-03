# Start from the latest golang base image
FROM golang:1.24-alpine3.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /src/app

# Installing Git for alpine
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /nfs-metrics-exporter

# Second stage
FROM alpine:3.20

WORKDIR /usr/local/bin/nfs-metrics-exporter

COPY --from=builder /nfs-metrics-exporter /nfs-metrics-exporter

CMD ["--cmd", "exporter"]
ENTRYPOINT ["./nfs-metrics-exporter"]
