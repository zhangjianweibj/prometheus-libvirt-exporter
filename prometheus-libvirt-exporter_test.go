package main

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	libvirt_schema "github.com/zhangjianweibj/prometheus-libvirt-exporter/libvirt_schema"
	"testing"
)

func init() {

}

//go lang bug ,1.9 bug still alive
/*
https://www.oschina.net/question/203177_2213070
https://www.oschina.net/question/203177_2213070

**/

func TestUnmarshal(t *testing.T) {
	var (
		str = `
<domain type='kvm'>
  <name>instance-00000212</name>
  <uuid>a6b57d2e-dad0-4860-9104-6eb072935126</uuid>
  <metadata>
    <nova:instance xmlns:nova="http://openstack.org/xmlns/libvirt/nova/1.0">
      <nova:package version="15.0.4-1.el7"/>
      <nova:name>LqnyzNfe</nova:name>
      <nova:creationTime>2018-04-23 07:08:32</nova:creationTime>
      <nova:flavor name="c1.micro">
        <nova:memory>1024</nova:memory>
        <nova:disk>20</nova:disk>
        <nova:swap>1024</nova:swap>
        <nova:ephemeral>0</nova:ephemeral>
        <nova:vcpus>1</nova:vcpus>
      </nova:flavor>
      <nova:owner>
        <nova:user uuid="f06c4cf1c4e14fda958c73f750049e7b">tenant</nova:user>
        <nova:project uuid="93ac887ce5794c778320a88c3024b1ad">ssdaW2wp</nova:project>
      </nova:owner>
      <nova:root type="image" uuid="674e506f-1791-435a-a85f-f983d7c9fef6"/>
    </nova:instance>
  </metadata>
  <memory unit='KiB'>1048576</memory>
  <currentMemory unit='KiB'>1048576</currentMemory>
  <vcpu placement='static'>1</vcpu>
  <cputune>
    <shares>1024</shares>
  </cputune>
  <sysinfo type='smbios'>
    <system>
      <entry name='manufacturer'>RDO</entry>
      <entry name='product'>OpenStack Compute</entry>
      <entry name='version'>15.0.4-1.el7</entry>
      <entry name='serial'>abfc2417-f63c-48df-a8ca-df3b37bae262</entry>
      <entry name='uuid'>a6b57d2e-dad0-4860-9104-6eb072935126</entry>
      <entry name='family'>Virtual Machine</entry>
    </system>
  </sysinfo>
  <os>
    <type arch='x86_64' machine='pc-i440fx-rhel7.3.0'>hvm</type>
    <boot dev='hd'/>
    <smbios mode='sysinfo'/>
  </os>
  <features>
    <acpi/>
    <apic/>
  </features>
  <cpu mode='host-model' check='partial'>
    <model fallback='allow'/>
    <topology sockets='1' cores='1' threads='1'/>
  </cpu>
  <clock offset='utc'>
    <timer name='pit' tickpolicy='delay'/>
    <timer name='rtc' tickpolicy='catchup'/>
    <timer name='hpet' present='no'/>
  </clock>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>destroy</on_crash>
  <devices>
    <emulator>/usr/libexec/qemu-kvm</emulator>
    <disk type='block' device='disk'>
      <driver name='qemu' type='raw' cache='none' io='native'/>
      <source dev='/dev/disk/by-path/ip-10.110.20.107:3260-iscsi-iqn.2010-10.org.openstack:volume-9bdce751-6bd2-495a-abc0-cfedbdfdc8be-lun-0'/>
      <target dev='vda' bus='virtio'/>
      <serial>9bdce751-6bd2-495a-abc0-cfedbdfdc8be</serial>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0'/>
    </disk>
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2' cache='none'/>
      <source file='/var/lib/nova/instances/a6b57d2e-dad0-4860-9104-6eb072935126/disk.swap'/>
      <target dev='vdb' bus='virtio'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0'/>
    </disk>
    <disk type='file' device='cdrom'>
      <driver name='qemu' type='raw' cache='none'/>
      <source file='/var/lib/nova/instances/a6b57d2e-dad0-4860-9104-6eb072935126/disk.config'/>
      <target dev='hda' bus='ide'/>
      <readonly/>
      <address type='drive' controller='0' bus='0' target='0' unit='0'/>
    </disk>
    <controller type='usb' index='0' model='piix3-uhci'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x2'/>
    </controller>
    <controller type='pci' index='0' model='pci-root'/>
    <controller type='ide' index='0'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x1'/>
    </controller>
    <interface type='bridge'>
      <mac address='fa:16:3e:76:c3:48'/>
      <source bridge='brqc2028cd6-12'/>
      <target dev='tap032a4f2e-8b'/>
      <model type='virtio'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
    </interface>
    <serial type='file'>
      <source path='/var/lib/nova/instances/a6b57d2e-dad0-4860-9104-6eb072935126/console.log'/>
      <target port='0'/>
    </serial>
    <serial type='pty'>
      <target port='1'/>
    </serial>
    <console type='file'>
      <source path='/var/lib/nova/instances/a6b57d2e-dad0-4860-9104-6eb072935126/console.log'/>
      <target type='serial' port='0'/>
    </console>
    <input type='tablet' bus='usb'>
      <address type='usb' bus='0' port='1'/>
    </input>
    <input type='mouse' bus='ps2'/>
    <input type='keyboard' bus='ps2'/>
    <graphics type='vnc' port='-1' autoport='yes' listen='10.110.20.107' keymap='en-us'>
      <listen type='address' address='10.110.20.107'/>
    </graphics>
    <video>
      <model type='cirrus' vram='16384' heads='1' primary='yes'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
    </video>
    <memballoon model='virtio'>
      <stats period='10'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
    </memballoon>
  </devices>
</domain>

    `
	)

	r := libvirt_schema.Domain{}
	err := xml.Unmarshal([]byte(str), &r)
	if err != nil {
		fmt.Println(err, r)
	}
	assert.Equal(nil, r.Metadata.NovaInstance.Name, "LqnyzNfe")
	assert.Equal(nil, r.Metadata.NovaInstance.Flavor.FlavorName, "c1.micro")
	fmt.Printf("xml name=%#v\n", r.Metadata.NovaInstance.XMLName)
	fmt.Printf("nova name=%#v\n", r.Metadata.NovaInstance.Name)
	fmt.Printf("nova =%#v\n", r.Metadata)
}
