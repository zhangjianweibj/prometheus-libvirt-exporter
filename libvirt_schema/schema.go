package libvirt_schema

import "encoding/xml"

type Domain struct {
	Devices Devices `xml:"devices"`
	Name string `xml:"name"`
	UUID string `xml:"uuid"`
	Metadata Metadata `xml:"metadata"`
}

type Metadata struct {
	NovaInstance NovaInstance `xml:"instance"`
}

type NovaInstance struct {
	XMLName xml.Name `xml:"instance"`
	NovaName string `xml:"name"`
}

type Devices struct {
	Disks      []Disk      `xml:"disk"`
	Interfaces []Interface `xml:"interface"`
}

type Disk struct {
	Device string     `xml:"device,attr"`
	Source DiskSource `xml:"source"`
	Target DiskTarget `xml:"target"`
}

type DiskSource struct {
	File string `xml:"file,attr"`
}

type DiskTarget struct {
	Device string `xml:"dev,attr"`
}

type Interface struct {
	Source InterfaceSource `xml:"source"`
	Target InterfaceTarget `xml:"target"`
}

type InterfaceSource struct {
	Bridge string `xml:"bridge,attr"`
}

type InterfaceTarget struct {
	Device string `xml:"dev,attr"`
}