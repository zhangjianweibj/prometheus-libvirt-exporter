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
