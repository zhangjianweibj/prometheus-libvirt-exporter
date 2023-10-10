FROM golang:1.21 as builder
WORKDIR /usr/src/app

# Requirements/Dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o prometheus-libvirt-exporter

FROM scratch
COPY --from=builder /usr/src/app/prometheus-libvirt-exporter /prometheus-libvirt-exporter
# Default listen on port 9177
EXPOSE 9177
# Start
CMD ["/prometheus-libvirt-exporter"]