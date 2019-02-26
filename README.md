# prometheus-libvirt-exporter
prometheus-libvirt-exporter for host and vm metrics exposed for prometheus, written in Go with pluggable metric collectors.
By default, this exporter listens on TCP port 9000,Path '/metrics',to expose metrics.vm's tags contain userId,userName,ProjectId,ProjectName.

[![Build Status](https://travis-ci.org/zhangjianweibj/prometheus-libvirt-exporter.svg?branch=master)](https://travis-ci.org/zhangjianweibj/prometheus-libvirt-exporter)
[![codecov](https://codecov.io/gh/zhangjianweibj/prometheus-libvirt-exporter/branch/master/graph/badge.svg)](https://codecov.io/gh/zhangjianweibj/prometheus-libvirt-exporter)
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


## Example

+ libvirt_domain_block_stats_read_bytes_total{domain="instance-0000503f",host="cmp001.inspurcloud.com",instanceId="47e22753-6874-4aec-801f-7126bffa323b",instanceName="slb-vlan202-200",projectId="560529102c8e4f39a6dc13552f4bd08b",projectName="slb-test-project",source_file="",target_device="vda",userId="1990c9213f9f4b52bbf8931a2fadcc4e",userName="slb-admin"} 1.18640128e+08
+ libvirt_domain_block_stats_read_bytes_total{domain="instance-0000503f",host="cmp001.inspurcloud.com",instanceId="47e22753-6874-4aec-801f-7126bffa323b",instanceName="slb-vlan202-200",projectId="560529102c8e4f39a6dc13552f4bd08b",projectName="slb-test-project",source_file="",target_device="vdb",userId="1990c9213f9f4b52bbf8931a2fadcc4e",userName="slb-admin"} 733696
(```)
libvirt_domain_block_stats_read_bytes_total{domain="instance-00005042",host="cmp001.inspurcloud.com",instanceId="6227d8d1-d925-49e4-8aa8-727f93738424",instanceName="slb-vlan202-201",projectId="560529102c8e4f39a6dc13552f4bd08b",projectName="slb-test-project",source_file="",target_device="vda",userId="1990c9213f9f4b52bbf8931a2fadcc4e",userName="slb-admin"} 1.49806592e+08
libvirt_domain_block_stats_read_bytes_total{domain="instance-00005042",host="cmp001.inspurcloud.com",instanceId="6227d8d1-d925-49e4-8aa8-727f93738424",instanceName="slb-vlan202-201",projectId="560529102c8e4f39a6dc13552f4bd08b",projectName="slb-test-project",source_file="",target_device="vdb",userId="1990c9213f9f4b52bbf8931a2fadcc4e",userName="slb-admin"} 733696
libvirt_domain_block_stats_read_bytes_total{domain="instance-00005b9a",host="cmp001.inspurcloud.com",instanceId="7634531f-3b18-4245-a312-5530f0006159",instanceName="BM_VM_Ubuntu1604",projectId="140785795de64945b02363661eb9e769",projectName="admin",source_file="",target_device="vda",userId="99e8e0d525b54c63bbe67799c9118d15",userName="cps"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-00005b9a",host="cmp001.inspurcloud.com",instanceId="7634531f-3b18-4245-a312-5530f0006159",instanceName="BM_VM_Ubuntu1604",projectId="140785795de64945b02363661eb9e769",projectName="admin",source_file="",target_device="vdb",userId="99e8e0d525b54c63bbe67799c9118d15",userName="cps"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006269",host="cmp001.inspurcloud.com",instanceId="fd67d76e-7912-4931-ac49-202865eb8dcf",instanceName="horizon-test",projectId="0c6a6f2d9af346c1aebc3db3938aa255",projectName="sjt",source_file="",target_device="vda",userId="5ea973cadbe14a0fb1eb593ee8b9ff21",userName="sjt"} 3.48270592e+08
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006269",host="cmp001.inspurcloud.com",instanceId="fd67d76e-7912-4931-ac49-202865eb8dcf",instanceName="horizon-test",projectId="0c6a6f2d9af346c1aebc3db3938aa255",projectName="sjt",source_file="",target_device="vdb",userId="5ea973cadbe14a0fb1eb593ee8b9ff21",userName="sjt"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-0000635f",host="cmp001.inspurcloud.com",instanceId="ac21a592-e802-4dd6-ad31-68376830154f",instanceName="guo.qos.test",projectId="140785795de64945b02363661eb9e769",projectName="admin",source_file="",target_device="vda",userId="e1a5c2d73f714b7f80866838fd20e102",userName="guochunting"} 1.09243392e+08
libvirt_domain_block_stats_read_bytes_total{domain="instance-0000635f",host="cmp001.inspurcloud.com",instanceId="ac21a592-e802-4dd6-ad31-68376830154f",instanceName="guo.qos.test",projectId="140785795de64945b02363661eb9e769",projectName="admin",source_file="",target_device="vdb",userId="e1a5c2d73f714b7f80866838fd20e102",userName="guochunting"} 1.901056e+06
libvirt_domain_block_stats_read_bytes_total{domain="instance-000063bf",host="cmp001.inspurcloud.com",instanceId="0212dd40-08f8-4ee6-815f-dbce6549b939",instanceName="jiasr-firewall-勿删",projectId="65a859f362f749ce95237cbd08c30edf",projectName="vpc",source_file="",target_device="vda",userId="11c95fd28341443592128de9b33f1c16",userName="vpc"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-000063bf",host="cmp001.inspurcloud.com",instanceId="0212dd40-08f8-4ee6-815f-dbce6549b939",instanceName="jiasr-firewall-勿删",projectId="65a859f362f749ce95237cbd08c30edf",projectName="vpc",source_file="",target_device="vdb",userId="11c95fd28341443592128de9b33f1c16",userName="vpc"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006578",host="cmp001.inspurcloud.com",instanceId="db55adf7-634d-4fe2-905e-881b0d9d070f",instanceName="slb-f5-test-tian",projectId="4db6d6b8742442bab42e06a0d6932fbb",projectName="slb-shared",source_file="",target_device="vda",userId="1990c9213f9f4b52bbf8931a2fadcc4e",userName="slb-admin"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006578",host="cmp001.inspurcloud.com",instanceId="db55adf7-634d-4fe2-905e-881b0d9d070f",instanceName="slb-f5-test-tian",projectId="4db6d6b8742442bab42e06a0d6932fbb",projectName="slb-shared",source_file="",target_device="vdb",userId="1990c9213f9f4b52bbf8931a2fadcc4e",userName="slb-admin"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-000066a1",host="cmp001.inspurcloud.com",instanceId="6770b291-0076-4847-a272-a0209163af95",instanceName="centos7_qga_guo",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",source_file="",target_device="vda",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 1.04454656e+08
libvirt_domain_block_stats_read_bytes_total{domain="instance-000066a1",host="cmp001.inspurcloud.com",instanceId="6770b291-0076-4847-a272-a0209163af95",instanceName="centos7_qga_guo",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",source_file="",target_device="vdb",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 4.237312e+06
libvirt_domain_block_stats_read_bytes_total{domain="instance-000066a1",host="cmp001.inspurcloud.com",instanceId="6770b291-0076-4847-a272-a0209163af95",instanceName="centos7_qga_guo",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",source_file="",target_device="vdc",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 3.744256e+06
libvirt_domain_block_stats_read_bytes_total{domain="instance-000066a1",host="cmp001.inspurcloud.com",instanceId="6770b291-0076-4847-a272-a0209163af95",instanceName="centos7_qga_guo",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",source_file="",target_device="vdd",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 3.744256e+06
libvirt_domain_block_stats_read_bytes_total{domain="instance-000066a1",host="cmp001.inspurcloud.com",instanceId="6770b291-0076-4847-a272-a0209163af95",instanceName="centos7_qga_guo",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",source_file="",target_device="vde",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 3.277312e+06
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006a4c",host="cmp001.inspurcloud.com",instanceId="71c36eb3-aa72-4235-9c4f-083cc9bce7a8",instanceName="CKSgaossmove3",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",source_file="",target_device="vda",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006a4c",host="cmp001.inspurcloud.com",instanceId="71c36eb3-aa72-4235-9c4f-083cc9bce7a8",instanceName="CKSgaossmove3",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",source_file="",target_device="vdb",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006adf",host="cmp001.inspurcloud.com",instanceId="125d677a-6649-4b35-aef8-73fa8236c5a9",instanceName="SLB-LVS-SHARED-pod-0000001093-1",projectId="4db6d6b8742442bab42e06a0d6932fbb",projectName="slb-shared",source_file="",target_device="vda",userId="1990c9213f9f4b52bbf8931a2fadcc4e",userName="slb-admin"} 1.52518144e+08
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006adf",host="cmp001.inspurcloud.com",instanceId="125d677a-6649-4b35-aef8-73fa8236c5a9",instanceName="SLB-LVS-SHARED-pod-0000001093-1",projectId="4db6d6b8742442bab42e06a0d6932fbb",projectName="slb-shared",source_file="",target_device="vdb",userId="1990c9213f9f4b52bbf8931a2fadcc4e",userName="slb-admin"} 733696
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006ae2",host="cmp001.inspurcloud.com",instanceId="23d43dad-a642-4481-85f3-d23f83de1306",instanceName="SLB-LVS-SHARED-pod-0000000112-1",projectId="71a70a350e734a33bde58b47c948c45f",projectName="slb-shared-test",source_file="",target_device="vda",userId="927aaabf2a4b4eea8fe4856ca15d48b5",userName="slb-admin-test"} 1.92544256e+08
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006ae2",host="cmp001.inspurcloud.com",instanceId="23d43dad-a642-4481-85f3-d23f83de1306",instanceName="SLB-LVS-SHARED-pod-0000000112-1",projectId="71a70a350e734a33bde58b47c948c45f",projectName="slb-shared-test",source_file="",target_device="vdb",userId="927aaabf2a4b4eea8fe4856ca15d48b5",userName="slb-admin-test"} 733696
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006ae5",host="cmp001.inspurcloud.com",instanceId="36549512-d362-4352-8b08-bc84405655f9",instanceName="SLB-LVS-SHARED-pod-0000001093-2",projectId="4db6d6b8742442bab42e06a0d6932fbb",projectName="slb-shared",source_file="",target_device="vda",userId="1990c9213f9f4b52bbf8931a2fadcc4e",userName="slb-admin"} 1.40893696e+08
(```)