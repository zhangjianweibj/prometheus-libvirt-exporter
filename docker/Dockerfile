from alpine:3.9

ARG BINARY_NAME

COPY "${BINARY_NAME}" "/usr/bin/prometheus-libvirt-exporter"

entrypoint ["prometheus-libvirt-exporter"] 
