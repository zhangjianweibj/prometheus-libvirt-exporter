# prometheus-libvirt-exporter
prometheus-libvirt-exporter for host and vm metrics exposed for prometheus, written in Go with pluggable metric collectors.
By default, this exporter listens on TCP port 9000,Path '/metrics',to expose metrics.

[![Build Status](https://travis-ci.org/zhangjianweibj/prometheus-libvirt-exporter.svg?branch=master)(https://travis-ci.org/zhangjianweibj/prometheus-libvirt-exporter)

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


##Example
# TYPE libvirt_domain_state_code gauge
libvirt_domain_state_code{domainName="instance-00000126",instanceId="4a50e208-eb6d-4a6e-904b-f9d9ef4ec483",instanceName="test",stateDesc="the domain is shut off"} 5
libvirt_domain_state_code{domainName="instance-00000157",instanceId="cee27a5f-278f-4f1c-b6b8-3d0879834cd1",instanceName="AWvlZ9qA-DRDSSingle-FeqKU6N4",stateDesc="the domain is shut off"} 5
libvirt_domain_state_code{domainName="instance-0000016c",instanceId="821e161e-ab48-41ba-914c-7febca28cb76",instanceName="aXdazRlV-ESCluster-d8MwUMzB",stateDesc="the domain is shut off"} 5
libvirt_domain_state_code{domainName="instance-00000192",instanceId="3f403d75-85fe-4965-9d37-dcfbfd06482f",instanceName="wrDCxC5p",stateDesc="the domain is running"} 1
libvirt_domain_state_code{domainName="instance-0000021a",instanceId="9bcd4f1f-a02f-47bc-bebc-600f0a798be1",instanceName="like",stateDesc="the domain is shut off"} 5
libvirt_domain_state_code{domainName="instance-0000021c",instanceId="2687aba1-f86e-4370-8993-b652234fc102",instanceName="KKAQH52k",stateDesc="the domain is shut off"} 5
libvirt_domain_state_code{domainName="instance-0000021d",instanceId="c86e3187-8f35-4256-9995-9e4dc3cd696f",instanceName="RF8aueUv",stateDesc="the domain is shut off"} 5