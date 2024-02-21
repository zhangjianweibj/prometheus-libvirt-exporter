
# Prometheus-libvirt-exporter
[![Build and Test](https://github.com/inovex/prometheus-libvirt-exporter/actions/workflows/build_and_test.yml/badge.svg)](https://github.com/inovex/prometheus-libvirt-exporter/actions/workflows/build_and_test.yml)
[![Lint Go Code](https://github.com/inovex/prometheus-libvirt-exporter/actions/workflows/lint.yml/badge.svg)](https://github.com/inovex/prometheus-libvirt-exporter/actions/workflows/lint.yml)

A prometheus-[libvirt](https://libvirt.org/)-exporter for host and vm metrics exposed for prometheus, written in Go with pluggable metric collectors.
By default, this exporter listens on TCP port 9177, path '/metrics', to expose metrics.

This exporter is built upon the [go-libvirt](https://github.com/digitalocean/go-libvirt) package developed by DigitalOcean. It offers a pure Go interface for interacting with Libvirt, leveraging the RPC interface provided by Libvirt. For detailed information about the Go bindings used, you can refer to the [Libvirt API reference](https://libvirt.org/html/index.html).



# Building and running

This release provides a set of assets for the prometheus-libvirt-exporter. It includes installation packages for various platforms (apk, deb, rpm) and the the binaries. Additionally, source code archives in both zip and tar.gz formats are available for download.

## Requirements
1. Gorelease: `go install github.com/goreleaser/goreleaser@latest`

2. Taskfile: `go install github.com/go-task/task/v3/cmd/task@latest`

## Local Building
1. Run `task build`

2. Afterwards all packages, binaries and archives are available in the `dist/` folder

## To see all available configuration flags:

`./prometheus-libvirt-exporter -h`


## metrics
Name | Label |Description
---------|---------|-------------
up||scraping libvirt's metrics state
libvirt_domains||number of domains
libvirt_domain_openstack_info | "domain", "instance_name", "instance_id", "flavor_name", "user_name", "user_id", "project_name", "project_id" | Aggregated OpenStack metadata as labels
libvirt_domain_info | "domain", "os_type", "os_type_machine", "os_type_arch" | e.g. os (operating system booting) settings as labels
libvirt_domain_info_state | "domain", "state_desc" | Code of the domain state,include state description
libvirt_domain_info_maximum_memory_bytes |"domain" | Maximum allowed memory of the domain
libvirt_domain_info_memory_usage_bytes |"domain" | Memory usage of the domain
libvirt_domain_info_virtual_cpus |"domain" | Number of virtual CPUs for the domain
libvirt_domain_info_cpu_time_seconds_total |"domain" | Amount of CPU time used by the domain
libvirt_domain_memory_stats_swap_in_bytes | "domain" | Memory swapped in for this domain(the total amount of data read from swap space)
libvirt_domain_memory_stats_swap_out_bytes | "domain" | Memory swapped out for this domain (the total amount of memory written out to swap space)
libvirt_domain_memory_stats_unused_bytes | "domain" | Memory unused by the domain
libvirt_domain_memory_stats_available_bytes | "domain" | Memory available to the domain
libvirt_domain_memory_stats_usable_bytes | "domain" | Memory usable by the domain (corresponds to 'Available' in /proc/meminfo)
libvirt_domain_memory_stats_rss_bytes |"domain" | Resident Set Size of the process running the domain
libvirt_domain_block_stats_info | "domain", "disk_type", "driver_cache", "driver_discard", "driver_name", "driver_type", "serial", "source_file", "target_bus", "target_device" | Metadata information on block devices
libvirt_domain_block_stats_read_bytes_total | "domain", "target_device", "host" | Number of bytes read from a block device, in bytes
libvirt_domain_block_stats_read_requests_total | "domain", "target_device", "host" | Number of read requests from a block device
libvirt_domain_block_stats_write_bytes_total | "domain", "target_device" | Number of bytes written from a block device, in bytes
libvirt_domain_block_stats_write_requests_total | "domain", "target_device" | Number of write requests from a block device
libvirt_domain_interface_stats_info | "domain", "interface_type", "mac_address", "model_type", "mtu_size", "source_bridge", "target_device" | Metadata on network interfaces
libvirt_domain_interface_stats_receive_bytes_total | "domain", "target_device", | Number of bytes received on a network interface, in bytes
libvirt_domain_interface_stats_receive_packets_total | "domain", "target_device" | Number of packets received on a network interface
libvirt_domain_interface_stats_receive_errors_total | "domain", "target_device" | Number of packet receive errors on a network interface
libvirt_domain_interface_stats_receive_drops_total| "domain", "target_device" | Number of packet receive drops on a network interface
libvirt_domain_interface_stats_transmit_bytes_total | "domain", "target_device" | Number of bytes transmitted on a network interface, in bytes
libvirt_domain_interface_stats_transmit_packets_total | "domain", "target_device" | Number of packets transmitted on a network interface
libvirt_domain_interface_stats_transmit_errors_total | "domain", "target_device" | Number of packet transmit errors on a network interface
libvirt_domain_interface_stats_transmit_drops_total | "domain", "target_device" | Number of packet transmit drops on a network interface
libvirt_domain_vcpu_current | "domain" | Number of current online vCPUs
libvirt_domain_vcpu_delay_seconds_total | "domain", "vcpu" | Time the vCPU spent waiting in the queue instead of running. Exposed to the VM as steal time
libvirt_domain_vcpu_maximum | "domain" | Number of maximum online vCPUs
libvirt_domain_vcpu_state | "domain", "vcpu" | State of the vCPU
libvirt_domain_vcpu_time_seconds_total | "domain", "vcpu" | Time spent by the virtual CPU
libvirt_domain_vcpu_wait_seconds_total | "domain", "vcpu" | Time the vCPU wants to run, but the host scheduler has something else running ahead of it


## Example

```
libvirt_domain_block_stats_info{disk_type="network",domain="instance-0001e06e",driver_cache="none",driver_discard="unmap",driver_name="qemu",driver_type="raw",serial="6ef20b92-1d9d-4de9-8a81-e324b98ae787",source_file="",target_bus="scsi",target_device="sda"} 1
libvirt_domain_block_stats_read_bytes_total{domain="instance-0001e06e",target_device="vda"} 1.497283072e+09
libvirt_domain_block_stats_read_requests_total{domain="instance-0001e06e",target_device="vda"} 23560
libvirt_domain_block_stats_write_bytes_total{domain="instance-0001e06e",target_device="vda"} 7.6914481664e+10
libvirt_domain_block_stats_write_requests_total{domain="instance-0001e06e",target_device="vda"} 2.409676e+06
libvirt_domain_info{domain="instance-0001e06e",os_type="hvm",os_type_arch="x86_64",os_type_machine="pc-q35-4.0"} 1
libvirt_domain_info_cpu_time_seconds_total{domain="instance-0001e06e"} 136215.65
libvirt_domain_info_maximum_memory_bytes{domain="instance-0001e06e"} 1.7179869184e+10
libvirt_domain_info_memory_usage_bytes{domain="instance-0001e06e"} 1.7179869184e+10
libvirt_domain_info_state{domain="instance-0001e06e",state_desc="the domain is running"} 1
libvirt_domain_info_virtual_cpus{domain="instance-0001e06e"} 4
libvirt_domain_interface_stats_info{domain="instance-0001e06e",interface_type="bridge",mac_address="fa:16:3e:fe:51:0a",model_type="virtio",mtu_size="7950",source_bridge="brq1cbc2c2b-af",target_device="tap3c5556a4-93"} 1
libvirt_domain_interface_stats_receive_bytes_total{domain="instance-0001e06e",target_device="tapab672ce4-11"} 1.589638794e+09
libvirt_domain_interface_stats_receive_drops_total{domain="instance-0001e06e",target_device="tapab672ce4-11"} 0
libvirt_domain_interface_stats_receive_errors_total{domain="instance-0001e06e",target_device="tapab672ce4-11"} 0
libvirt_domain_interface_stats_receive_packets_total{domain="instance-0001e06e",target_device="tapab672ce4-11"} 4.671267e+06
libvirt_domain_interface_stats_transmit_bytes_total{domain="instance-0001e06e",target_device="tapab672ce4-11"} 6.90886551e+08
libvirt_domain_interface_stats_transmit_drops_total{domain="instance-0001e06e",target_device="tapab672ce4-11"} 0
libvirt_domain_interface_stats_transmit_errors_total{domain="instance-0001e06e",target_device="tapab672ce4-11"} 0
libvirt_domain_interface_stats_transmit_packets_total{domain="instance-0001e06e",target_device="tapab672ce4-11"} 1.412009e+06
libvirt_domain_memory_stats_available_bytes{domain="instance-0001e06e"} 1.676619776e+10
libvirt_domain_memory_stats_rss_bytes{domain="instance-0001e06e"} 6.386884608e+09
libvirt_domain_memory_stats_swap_in_bytes{domain="instance-0001e06e"} 0
libvirt_domain_memory_stats_swap_out_bytes{domain="instance-0001e06e"} 0
libvirt_domain_memory_stats_unused_bytes{domain="instance-0001e06e"} 1.3844406272e+10
libvirt_domain_memory_stats_usable_bytes{domain="instance-0001e06e"} 1.4880370688e+10
libvirt_domain_openstack_info{domain="instance-0001e06e",flavor_name="z1.4xlarge",instance_id="a12423b02-4a36-4530-bf25-acb8ba80b1b1",instance_name="openstackInstanceName",project_id="hghngfhbf45435352353623gvfegt352",project_name="openstackProjectName",user_id="",user_name="openstackUserName"} 1
libvirt_domain_vcpu_current{domain="instance-00000131"} 2
libvirt_domain_vcpu_delay_seconds_total{domain="instance-00000131",vcpu="0"} 6309719178
libvirt_domain_vcpu_maximum{domain="instance-00000131"} 2
libvirt_domain_vcpu_state{domain="instance-00000131",vcpu="0"} 1
libvirt_domain_vcpu_time_seconds_total{domain="instance-00000131",vcpu="0"} 2111850000000
libvirt_domain_vcpu_wait_seconds_total{domain="instance-00000131",vcpu="0"} 4103560000000
```
