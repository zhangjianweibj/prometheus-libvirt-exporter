# prometheus-libvirt-exporter
prometheus-libvirt-exporter for host and vm metrics exposed for prometheus, written in Go with pluggable metric collectors.
By default, this exporter listens on TCP port 9000,Path '/metrics',to expose metrics.vm's tags contain userId,userName,ProjectId,ProjectName.

[![Build Status](https://travis-ci.org/zhangjianweibj/prometheus-libvirt-exporter.svg?branch=master)](https://travis-ci.org/zhangjianweibj/prometheus-libvirt-exporter)
[![codecov](https://codecov.io/gh/zhangjianweibj/prometheus-libvirt-exporter/branch/master/graph/badge.svg)](https://codecov.io/gh/zhangjianweibj/prometheus-libvirt-exporter)
# Building and running

## use go dep(depressed)
1. install go dep

2. cp $GOPATH/bin/dep /usr/bin/

3. dep ensure

4. go build prometheus-libvirt-exporter.go

5. ./prometheus-libvirt-exporter

## use go mod tool
1. go build prometheus-libvirt-exporter.go 
2. ./prometheus-libvirt-exporter

## To see all available configuration flags:

./prometheus-libvirt-exporter -h


## metrics
Name | Label |Description
---------|---------|-------------
up|"host"|scraping libvirt's metrics state
domains_number|"host"|get number of domains
domain_state_code|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "stateDesc", "host"|code of the domain state,include state description
maximum_memory_bytes|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "host"|Maximum allowed memory of the domain, in bytes
memory_usage_bytes|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "host"|Memory usage of the domain, in bytes
virtual_cpus|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "host"|Number of virtual CPUs for the domain
cpu_time_seconds_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "host"|Amount of CPU time used by the domain, in seconds
read_bytes_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_file", "target_device", "host"|Number of bytes read from a block device, in bytes
read_requests_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_file", "target_device", "host"|Number of read requests from a block device
write_bytes_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_file", "target_device", "host"|Number of bytes written from a block device, in bytes
write_requests_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_file", "target_device", "host"|Number of write requests from a block device
receive_bytes_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of bytes received on a network interface, in bytes
receive_packets_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packets received on a network interface
receive_errors_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packet receive errors on a network interface
receive_drops_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packet receive drops on a network interface
transmit_bytes_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of bytes transmitted on a network interface, in bytes
transmit_packets_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packets transmitted on a network interface
transmit_errors_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packet transmit errors on a network interface
transmit_drops_total|"domain", "instanceName", "instanceId", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packet transmit drops on a network interface


## Example

```
libvirt_domain_block_stats_read_requests_total{domain="instance-000070dc",host="cmp001.inspurcloud.com",instanceId="d83e808b-46c7-4b51-90e8-b43934c61756",instanceName="SLB-POD-webappslb7c8356fe-1266-42a3-a221-40388e52ac0c-1",projectId="3b5c57d6102f4c2bbb6361807ae1368c",projectName="hk",source_file="",target_device="vda",userId="ed5d5ba4cdf0445387e2914f46e96e0c",userName="11190969-e7cf-406d-8f49-da9883e4546e"} 718219
libvirt_domain_block_stats_write_requests_total{domain="instance-00006e36",host="cmp001.inspurcloud.com",instanceId="4a0de6e7-33cb-4067-8eb8-fe2f7c2f27b5",instanceName="centos7.3_0201",projectId="3b5c57d6102f4c2bbb6361807ae1368c",projectName="hk",source_file="",target_device="vdb",userId="adf255d1f79845ee8c470f7497fb01e7",userName="ebs"} 0
libvirt_domain_info_cpu_time_seconds_total{domain="instance-000073be",host="cmp001.inspurcloud.com",instanceId="07340c2d-be0c-4394-a3f4-4a49d56a853b",instanceName="CPS-201912915505",projectId="894c48b5bc4247c3a1841e5ff43d0d88",projectName="zhangyan",userId="51efbb05996742a58a69236a5fc8d5e3",userName="zhangyan"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-00005b9a",host="cmp001.inspurcloud.com",instanceId="7634531f-3b18-4245-a312-5530f0006159",instanceName="BM_VM_Ubuntu1604",projectId="140785795de64945b02363661eb9e769",projectName="admin",source_file="",target_device="vdb",userId="99e8e0d525b54c63bbe67799c9118d15",userName="cps"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006269",host="cmp001.inspurcloud.com",instanceId="fd67d76e-7912-4931-ac49-202865eb8dcf",instanceName="horizon-test",projectId="0c6a6f2d9af346c1aebc3db3938aa255",projectName="sjt",source_file="",target_device="vda",userId="5ea973cadbe14a0fb1eb593ee8b9ff21",userName="sjt"} 3.48270592e+08
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006269",host="cmp001.inspurcloud.com",instanceId="fd67d76e-7912-4931-ac49-202865eb8dcf",instanceName="horizon-test",projectId="0c6a6f2d9af346c1aebc3db3938aa255",projectName="sjt",source_file="",target_device="vdb",userId="5ea973cadbe14a0fb1eb593ee8b9ff21",userName="sjt"} 0
libvirt_domain_info_maximum_memory_bytes{domain="instance-0000635f",host="cmp001.inspurcloud.com",instanceId="ac21a592-e802-4dd6-ad31-68376830154f",instanceName="guo.qos.test",projectId="140785795de64945b02363661eb9e769",projectName="admin",userId="e1a5c2d73f714b7f80866838fd20e102",userName="guochunting"} 1.073741824e+09
libvirt_domain_info_maximum_memory_bytes{domain="instance-000063bf",host="cmp001.inspurcloud.com",instanceId="0212dd40-08f8-4ee6-815f-dbce6549b939",instanceName="jiasr-firewall-勿删",projectId="65a859f362f749ce95237cbd08c30edf",projectName="vpc",userId="11c95fd28341443592128de9b33f1c16",userName="vpc"} 4.294967296e+09
libvirt_domain_info_memory_usage_bytes{domain="instance-000066a1",host="cmp001.inspurcloud.com",instanceId="6770b291-0076-4847-a272-a0209163af95",instanceName="centos7_qga_guo",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 1.073741824e+09
libvirt_domain_info_memory_usage_bytes{domain="instance-00006a4c",host="cmp001.inspurcloud.com",instanceId="71c36eb3-aa72-4235-9c4f-083cc9bce7a8",instanceName="CKSgaossmove3",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 2.147483648e+09
libvirt_domain_interface_stats_receive_bytes_total{domain="instance-00006ecf",host="cmp001.inspurcloud.com",instanceId="5c623799-037a-4b97-95ea-5521414e0ff0",instanceName="testvpc2",projectId="3b5c57d6102f4c2bbb6361807ae1368c",projectName="hk",source_bridge="qbr693de2e9-f7",target_device="tap693de2e9-f7",userId="adf255d1f79845ee8c470f7497fb01e7",userName="ebs"} 0
libvirt_domain_interface_stats_receive_errors_total{domain="instance-00006ae5",host="cmp001.inspurcloud.com",instanceId="36549512-d362-4352-8b08-bc84405655f9",instanceName="SLB-LVS-SHARED-pod-0000001093-2",projectId="4db6d6b8742442bab42e06a0d6932fbb",projectName="slb-shared",source_bridge="qbr5c166814-65",target_device="tap5c166814-65",userId="1990c9213f9f4b52bbf8931a2fadcc4e",userName="slb-admin"} 0
libvirt_domain_interface_stats_receive_packets_total{domain="instance-000066a1",host="cmp001.inspurcloud.com",instanceId="6770b291-0076-4847-a272-a0209163af95",instanceName="centos7_qga_guo",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",source_bridge="qbr3a9f355a-92",target_device="tap3a9f355a-92",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 31459
libvirt_domain_interface_stats_receive_packets_total{domain="instance-000073a0",host="cmp001.inspurcloud.com",instanceId="a74ff9db-8767-4bd4-99a7-bdf96ee03384",instanceName="ECS-DLY-test",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",source_bridge="qbr9f38cb1f-9a",target_device="tap9f38cb1f-9a",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 1
libvirt_domain_interface_stats_transmit_bytes_total{domain="instance-000070f4",host="cmp001.inspurcloud.com",instanceId="c05737e1-9cff-445e-b196-3c3b148088e8",instanceName="SLB-POD-SLB-201902140946333-1",projectId="3b5c57d6102f4c2bbb6361807ae1368c",projectName="hk",source_bridge="qbr1f4ecfdb-9f",target_device="tap1f4ecfdb-9f",userId="ed5d5ba4cdf0445387e2914f46e96e0c",userName="11190969-e7cf-406d-8f49-da9883e4546e"} 3.399858e+06
libvirt_domain_interface_stats_transmit_drops_total{domain="instance-000074e1",host="cmp001.inspurcloud.com",instanceId="32f9cfdc-77e5-430e-a9be-24cbc5d85fe3",instanceName="ECS-2019225165058",projectId="894c48b5bc4247c3a1841e5ff43d0d88",projectName="zhangyan",source_bridge="qbr6bf53fe6-2f",target_device="tap6bf53fe6-2f",userId="068ddef4603b49a485b39d33ac4ba80e",userName="223f6fac-cacf-4447-ba2e-ee42ab125148"} 0
libvirt_domain_interface_stats_transmit_packets_total{domain="instance-00006a4c",host="cmp001.inspurcloud.com",instanceId="71c36eb3-aa72-4235-9c4f-083cc9bce7a8",instanceName="CKSgaossmove3",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",source_bridge="qbraeea3b2f-d2",target_device="tapaeea3b2f-d2",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 0
libvirt_domain_state_code{domain="instance-0000722f",host="cmp001.inspurcloud.com",instanceId="7d10f5dd-f95d-45a7-af99-8b6d917c1bda",instanceName="retrytest",projectId="766b0cb8a6c648068d082c467e88fba6",projectName="gaoss",stateDesc="the domain is running",userId="0d461f47463b45a0b9c882c3355ac341",userName="gaoss"} 1
libvirt_domain_state_code{domain="instance-000074ed",host="cmp001.inspurcloud.com",instanceId="534cd6d4-f150-45f8-bba0-009cd47e1013",instanceName="SLB-POD-yaoyifei-test-1",projectId="3b5c57d6102f4c2bbb6361807ae1368c",projectName="hk",stateDesc="the domain is running",userId="ed5d5ba4cdf0445387e2914f46e96e0c",userName="11190969-e7cf-406d-8f49-da9883e4546e"} 1
libvirt_domain_state_code{domain="instance-000074f0",host="cmp001.inspurcloud.com",instanceId="e21115db-f688-4460-81a5-9a6383bb1399",instanceName="ECS-201922695542",projectId="894c48b5bc4247c3a1841e5ff43d0d88",projectName="zhangyan",stateDesc="the domain is running",userId="068ddef4603b49a485b39d33ac4ba80e",userName="223f6fac-cacf-4447-ba2e-ee42ab125148"} 1

```