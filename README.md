# prometheus-libvirt-exporter
prometheus-libvirt-exporter for host and vm metrics exposed for prometheus, written in Go with pluggable metric collectors.
By default, this exporter listens on TCP port 9000,Path '/metrics',to expose metrics.

#Building and running

1.install go dep

2.cp $GOPATH/bin/dep /usr/bin/

3.dep ensure

4.go  build prometheus-libvirt-exporter.go

5. ./prometheus-libvirt-exporter

To see all available configuration flags:
./prometheus-libvirt-exporter -h
