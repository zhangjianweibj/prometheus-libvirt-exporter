FROM scratch
COPY prometheus-libvirt-exporter /prometheus-libvirt-exporter 
ENTRYPOINT ["/prometheus-libvirt-exporter"]