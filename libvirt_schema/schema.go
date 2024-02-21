package libvirt_schema

import "encoding/xml"

type Domain struct {
	Devices    Devices    `xml:"devices"`
	Name       string     `xml:"name"`
	UUID       string     `xml:"uuid"`
	Metadata   Metadata   `xml:"metadata"`
	OSMetadata OSMetadata `xml:"os"`
}

type Metadata struct {
	NovaInstance NovaInstance `xml:"instance"`
}

type OSMetadata struct {
	Type OSType `xml:"type"`
}

type OSType struct {
	Arch    string `xml:"arch,attr"`
	Machine string `xml:"machine,attr"`
	Value   string `xml:",chardata"`
}

type NovaInstance struct {
	XMLName xml.Name   `xml:"instance"`
	Name    string     `xml:"name"`
	Owner   NovaOwner  `xml:"owner"`
	Flavor  NovaFlavor `xml:"flavor"`
}

type NovaOwner struct {
	XMLName xml.Name    `xml:"owner"`
	User    NovaUser    `xml:"user"`
	Project NovaProject `xml:"project"`
}

type NovaUser struct {
	UserId   string `xml:"uuid,attr"`
	UserName string `xml:",chardata"`
}

type NovaProject struct {
	ProjectId   string `xml:"uuid,attr"`
	ProjectName string `xml:",chardata"`
}

type NovaFlavor struct {
	FlavorName string `xml:"name,attr"`
}

type Devices struct {
	Disks      []Disk      `xml:"disk"`
	Interfaces []Interface `xml:"interface"`
}

type Disk struct {
	Device string     `xml:"device,attr"`
	Type   string     `xml:"type,attr"`
	Serial string     `xml:"serial"`
	Driver DiskDriver `xml:"driver"`
	Source DiskSource `xml:"source"`
	Target DiskTarget `xml:"target"`
}

type DiskDriver struct {
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"`
	Cache   string `xml:"cache,attr"`
	Discard string `xml:"discard,attr"`
}

type DiskSource struct {
	File     string `xml:"file,attr"`
	Protocol string `xml:"protocol,attr"`
}

type DiskTarget struct {
	Device string `xml:"dev,attr"`
	Bus    string `xml:"bus,attr"`
}

type Interface struct {
	Type   string          `xml:"type,attr"`
	Source InterfaceSource `xml:"source"`
	Target InterfaceTarget `xml:"target"`
	MAC    InterfaceMAC    `xml:"mac"`
	Model  InterfaceModel  `xml:"model"`
	MTU    InterfaceMTU    `xml:"mtu"`
}

type InterfaceSource struct {
	Bridge string `xml:"bridge,attr"`
}

type InterfaceTarget struct {
	Device string `xml:"dev,attr"`
}

type InterfaceMAC struct {
	Address string `xml:"address,attr"`
}

type InterfaceModel struct {
	Type string `xml:"type,attr"`
}

type InterfaceMTU struct {
	Size string `xml:"size,attr"`
}
