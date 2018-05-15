# prometheus-libvirt-exporter
prometheus-libvirt-exporter for host and vm metrics exposed for prometheus, written in Go with pluggable metric collectors.
By default, this exporter listens on TCP port 9000,Path '/metrics',to expose metrics.

# Building and running

1.install go dep

2.cp $GOPATH/bin/dep /usr/bin/

3.dep ensure

4.go build prometheus-libvirt-exporter.go

5../prometheus-libvirt-exporter

## To see all available configuration flags:

./prometheus-libvirt-exporter -h


## metrics
Name | Description
---------|-------------
up|scraping libvirt's metrics state
domains_number|get number of domains
domain_state_code|code of the domain state,include state description
maximum_memory_bytes|Maximum allowed memory of the domain, in bytes
memory_usage_bytes|Memory usage of the domain, in bytes
virtual_cpus|Number of virtual CPUs for the domain
cpu_time_seconds_total|Amount of CPU time used by the domain, in seconds
read_bytes_total|Number of bytes read from a block device, in bytes
read_requests_total|Number of read requests from a block device
write_bytes_total|Number of bytes written from a block device, in bytes
write_requests_total|Number of write requests from a block device
receive_bytes_total|Number of bytes received on a network interface, in bytes
receive_packets_total|Number of packets received on a network interface
receive_errors_total|Number of packet receive errors on a network interface
receive_drops_total|Number of packet receive drops on a network interface
transmit_bytes_total|Number of bytes transmitted on a network interface, in bytes
transmit_packets_total|Number of packets transmitted on a network interface
transmit_errors_total|Number of packet transmit errors on a network interface
transmit_drops_total|Number of packet transmit drops on a network interface
