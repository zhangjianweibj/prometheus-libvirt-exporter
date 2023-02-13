# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang alpine image
FROM golang:alpine AS builder

#Working Dir inside builder container
WORKDIR /app

#copy go files
COPY . .

#Requirements/Dependencies
RUN go mod vendor
RUN go mod download && go mod verify

#Build the app

RUN go build prometheus-libvirt-exporter.go

# Build a image from alpine
FROM alpine

COPY --from=builder /app/prometheus-libvirt-exporter /usr/bin/prometheus-libvirt-exporter

EXPOSE 9000

ENTRYPOINT ["/usr/bin/prometheus-libvirt-exporter"]