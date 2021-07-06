# prometheus-libvirt-exporter
prometheus-libvirt-exporter for host and vm metrics exposed for prometheus, written in Go with pluggable metric collectors.
By default, this exporter listens on TCP port 9000,Path '/metrics',to expose metrics.vm's tags contain userId,userName,ProjectId,ProjectName.

[![Build Status](https://travis-ci.org/zhangjianweibj/prometheus-libvirt-exporter.svg?branch=master)](https://travis-ci.org/zhangjianweibj/prometheus-libvirt-exporter)
[![codecov](https://codecov.io/gh/zhangjianweibj/prometheus-libvirt-exporter/branch/master/graph/badge.svg)](https://codecov.io/gh/zhangjianweibj/prometheus-libvirt-exporter)
# Building and running

## Requirements
1. Gorelease: `go install github.com/goreleaser/goreleaser`

2. Taskfile: `go install github.com/go-task/task/v3/cmd/task@latest`

## use go dep(depressed)
1. install go dep

2. cp $GOPATH/bin/dep /usr/bin/

3. dep ensure

4. go build prometheus-libvirt-exporter.go

5. ./prometheus-libvirt-exporter

## Building
1. Run `task build`

2. Afterwards all packages, binaries and archives are available in the `dist/` folder

## To see all available configuration flags:

./prometheus-libvirt-exporter -h


## metrics
Name | Label |Description
---------|---------|-------------
up|"host"|scraping libvirt's metrics state
domains_number|"host"|get number of domains
domain_state_code|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "stateDesc", "host"|code of the domain state,include state description
maximum_memory_bytes|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"|Maximum allowed memory of the domain, in bytes
memory_usage_bytes|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"|Memory usage of the domain, in bytes
virtual_cpus|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"|Number of virtual CPUs for the domain
cpu_time_seconds_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"|Amount of CPU time used by the domain, in seconds
read_bytes_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_file", "target_device", "host"|Number of bytes read from a block device, in bytes
read_requests_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_file", "target_device", "host"|Number of read requests from a block device
write_bytes_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_file", "target_device", "host"|Number of bytes written from a block device, in bytes
write_requests_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_file", "target_device", "host"|Number of write requests from a block device
receive_bytes_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of bytes received on a network interface, in bytes
receive_packets_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packets received on a network interface
receive_errors_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packet receive errors on a network interface
receive_drops_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packet receive drops on a network interface
transmit_bytes_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of bytes transmitted on a network interface, in bytes
transmit_packets_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packets transmitted on a network interface
transmit_errors_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packet transmit errors on a network interface
transmit_drops_total|"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "source_bridge", "target_device", "host"|Number of packet transmit drops on a network interface


## Example

```
libvirt_domain_block_stats_read_requests_total{domain="instance-000070dc",host="cmp001.inspurcloud.com",instanceId="d**08b-46c7-4b51-90e8-b43934c61756",instanceName="SLB-POD-webappslb7c8356fe-1266-42a3-a221-40388e52ac0c-1",flavorName="c1.micro",projectId="3b5c57***bbb6361807ae1368c",projectName="hk",source_file="",target_device="vda",userId="ed5d5ba4cdf0445387e2914f46e96e0c",userName="11**69-e7cf-406d-8f49-da9883e4546e"} 718219
libvirt_domain_block_stats_write_requests_total{domain="instance-00006e36",host="cmp001.inspurcloud.com",instanceId="4a0**7-33cb-4067-8eb8-fe2f7c2f27b5",instanceName="centos7.3_0201",flavorName="c1.micro",projectId="3b5****2f4c2bbb6361807ae1368c",projectName="hk",source_file="",target_device="vdb",userId="adf***845ee8c470f7497fb01e7",userName="ebs"} 0
libvirt_domain_info_cpu_time_seconds_total{domain="instance-000073be",host="cmp001.inspurcloud.com",instanceId="07340**d-be0c-4394-a3f4-4a49d56a853b",instanceName="CPS-201912915505",flavorName="c1.micro",projectId="894c**8b5bc4247c3a1841e5ff43d0d88",projectName="z**gyan",userId="51efbb**96742a58a69236a5fc8d5e3",userName="zh**an"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-00005b9a",host="cmp001.inspurcloud.com",instanceId="7**531f-3b18-4245-a312-5530f0006159",instanceName="BM_VM_Ubuntu1604",flavorName="c1.micro",projectId="14078***e64945b02363661eb9e769",projectName="admin",source_file="",target_device="vdb",userId="99e8***5b54c63bbe67799c9118d15",userName="cps"} 0
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006269",host="cmp001.inspurcloud.com",instanceId="fd6**6e-7912-4931-ac49-202865eb8dcf",instanceName="horizon-test",flavorName="c1.micro",projectId="0c6a**d9af346c1aebc3db3938aa255",projectName="sjt",source_file="",target_device="vda",userId="5ea973cad**b1eb593ee8b9ff21",userName="sjt"} 3.48270592e+08
libvirt_domain_block_stats_read_bytes_total{domain="instance-00006269",host="cmp001.inspurcloud.com",instanceId="f**76e-7912-4931-ac49-202865eb8dcf",instanceName="horizon-test",flavorName="c1.micro",projectId="0c6a6**346c1aebc3db3938aa255",projectName="sjt",source_file="",target_device="vdb",userId="5ea973c**4a0fb1eb593ee8b9ff21",userName="sjt"} 0
libvirt_domain_info_maximum_memory_bytes{domain="instance-0000635f",host="cmp001.inspurcloud.com",instanceId="ac2**2-e802-4dd6-ad31-68376830154f",instanceName="guo.qos.test",flavorName="c1.micro",projectId="140**5de64945b02363661eb9e769",projectName="admin",userId="e1a5c2d73**f80866838fd20e102",userName="guo**ing"} 1.073741824e+09
libvirt_domain_info_maximum_memory_bytes{domain="instance-000063bf",host="cmp001.inspurcloud.com",instanceId="0**d40-08f8-4ee6-815f-dbce6549b939",instanceName="jiasr-firewall-勿删",flavorName="c1.micro",projectId="65**f362f749ce95237cbd08c30edf",projectName="vpc",userId="11c9**8341443592128de9b33f1c16",userName="vpc"} 4.294967296e+09
libvirt_domain_info_memory_usage_bytes{domain="instance-000066a1",host="cmp001.inspurcloud.com",instanceId="67**91-0076-4847-a272-a0209163af95",instanceName="centos7_qga_guo",flavorName="c1.micro",projectId="76**8a6c648068d082c467e88fba6",projectName="gaoss",userId="0d4**7463b45a0b9c882c3355ac341",userName="gaoss"} 1.073741824e+09
libvirt_domain_info_memory_usage_bytes{domain="instance-00006a4c",host="cmp001.inspurcloud.com",instanceId="7**eb3-aa72-4235-9c4f-083cc9bce7a8",instanceName="CKSgaossmove3",flavorName="c1.micro",projectId="766**6c648068d082c467e88fba6",projectName="gaoss",userId="0d**47463b45a0b9c882c3355ac341",userName="gaoss"} 2.147483648e+09
libvirt_domain_interface_stats_receive_bytes_total{domain="instance-00006ecf",host="cmp001.inspurcloud.com",instanceId="5c6**9-037a-4b97-95ea-5521414e0ff0",instanceName="testvpc2",flavorName="c1.micro",projectId="3b5c**2f4c2bbb6361807ae1368c",projectName="hk",source_bridge="qbr693de2e9-f7",target_device="tap693de2e9-f7",userId="adf255d1f79**e8c470f7497fb01e7",userName="ebs"} 0
libvirt_domain_interface_stats_receive_errors_total{domain="instance-00006ae5",host="cmp001.inspurcloud.com",instanceId="3**512-d362-4352-8b08-bc84405655f9",instanceName="SLB-LVS-SHARED-pod-0000001093-2",flavorName="c1.micro",projectId="4db6**2442bab42e06a0d6932fbb",projectName="slb-shared",source_bridge="qbr5c166814-65",target_device="tap5c166814-65",userId="1990c921****931a2fadcc4e",userName="slb-admin"} 0
libvirt_domain_interface_stats_receive_packets_total{domain="instance-000066a1",host="cmp001.inspurcloud.com",instanceId="6**291-0076-4847-a272-a0209163af95",instanceName="centos7_qga_guo",flavorName="c1.micro",projectId="7**0cb8a6c648068d082c467e88fba6",projectName="gaoss",source_bridge="qbr3a9f355a-92",target_device="tap3a9f355a-92",userId="0d**47463b45a0b9c882c3355ac341",userName="gaoss"} 31459
libvirt_domain_interface_stats_receive_packets_total{domain="instance-000073a0",host="cmp001.inspurcloud.com",instanceId="a**f9db-8767-4bd4-99a7-bdf96ee03384",instanceName="ECS-DLY-test",flavorName="c1.micro",projectId="766b0**6c648068d082c467e88fba6",projectName="gaoss",source_bridge="qbr9f38cb1f-9a",target_device="tap9f38cb1f-9a",userId="0d4**463b45a0b9c882c3355ac341",userName="gaoss"} 1
libvirt_domain_interface_stats_transmit_bytes_total{domain="instance-000070f4",host="cmp001.inspurcloud.com",instanceId="c**7e1-9cff-445e-b196-3c3b148088e8",instanceName="SLB-POD-SLB-201902140946333-1",flavorName="c1.micro",projectId="3b**d6102f4c2bbb6361807ae1368c",projectName="hk",source_bridge="qbr1f4ecfdb-9f",target_device="tap1f4ecfdb-9f",userId="e**a4cdf0445387e2914f46e96e0c",userName="11190**-e7cf-406d-8f49-da9883e4546e"} 3.399858e+06
libvirt_domain_interface_stats_transmit_drops_total{domain="instance-000074e1",host="cmp001.inspurcloud.com",instanceId="32**dc-77e5-430e-a9be-24cbc5d85fe3",instanceName="ECS-2019225165058",flavorName="c1.micro",projectId="894c**c4247c3a1841e5ff43d0d88",projectName="zh**an",source_bridge="qbr6bf53fe6-2f",target_device="tap6bf53fe6-2f",userId="068dd**03b49a485b39d33ac4ba80e",userName="223f6fac-***447-ba2e-ee42ab125148"} 0
libvirt_domain_interface_stats_transmit_packets_total{domain="instance-00006a4c",host="cmp001.inspurcloud.com",instanceId="**6eb3-aa72-4235-9c4f-083cc9bce7a8",instanceName="CKSgaossmove3",flavorName="c1.micro",projectId="7***a6c648068d082c467e88fba6",projectName="gaoss",source_bridge="qbraeea3b2f-d2",target_device="tapaeea3b2f-d2",userId="0d461f4746***9c882c3355ac341",userName="gaoss"} 0
libvirt_domain_state_code{domain="instance-0000722f",host="cmp001.inspurcloud.com",instanceId="7d1***d-f95d-45a7-af99-8b6d917c1bda",instanceName="retrytest",flavorName="c1.micro",projectId="766***a6c648068d082c467e88fba6",projectName="gaoss",stateDesc="the domain is running",userId="0d4***7463b45a0b9c882c3355ac341",userName="gaoss"} 1
libvirt_domain_state_code{domain="instance-000074ed",host="cmp001.inspurcloud.com",instanceId="53**4-f150-45f8-bba0-009cd47e1013",instanceName="SLB-POD-yaoyifei-test-1",flavorName="c1.micro",projectId="3b5c57d610**bbb6361807ae1368c",projectName="hk",stateDesc="the domain is running",userId="ed5**f0445387e2914f46e96e0c",userName="1***-e7cf-406d-8f49-da9883e4546e"} 1
libvirt_domain_state_code{domain="instance-000074f0",host="cmp001.inspurcloud.com",instanceId="e**db-f688-4460-81a5-9a6383bb1399",instanceName="ECS-201922695542",flavorName="c1.micro",projectId="8**c4247c3a1841e5ff43d0d88",projectName="z**an",stateDesc="the domain is running",userId="068***603b49a485b39d33ac4ba80e",userName="2***-cacf-4447-ba2e-ee42ab125148"} 1

```
