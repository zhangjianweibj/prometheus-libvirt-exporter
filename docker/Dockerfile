FROM alpine:3.9

ARG BINARY_NAME

COPY "${BINARY_NAME}" "/usr/bin/prometheus-libvirt-exporter"

ENTRYPOINT ["prometheus-libvirt-exporter"]
