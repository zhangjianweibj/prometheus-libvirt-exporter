package main

import (
	"fmt"
	"log"
	"net"
	"time"
	"github.com/digitalocean/go-libvirt"
	_"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus-libvirt-exporter/libvirt_schema"
	"encoding/xml"
)

func main() {
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	v, err := l.Version()
	if err != nil {
		log.Fatalf("failed to retrieve libvirt version: %v", err)
	}
	fmt.Println("Version:", v)

	domains, err := l.Domains()
	if err != nil {
		log.Fatalf("failed to retrieve domains: %v", err)
	}


	fmt.Println("ID\tName\t\tUUID")
	fmt.Printf("--------------------------------------------------------\n")
	for _, d := range domains {
		//fmt.Printf("%d\t%s\t%x\n", d.ID, d.Name, d.UUID)
		xmlDesc,error := l.DomainGetXMLDesc(d,0)
		if error !=nil {
			log.Fatalf("failed to DomainGetXMLDesc: %v",error)
			continue
		}
		var libvirtSchema libvirt_schema.Domain
		error = xml.Unmarshal([]byte(xmlDesc),&libvirtSchema)
		if error !=nil {
			log.Fatalf("failed to Unmarshal domains: %v",error)
			continue
		}

		fmt.Printf("%s\n",libvirtSchema)

	}


	fmt.Println("ID\tName\trstate\trmaxmem\t\trmemory\trvirCpu\trcputime")
	fmt.Printf("--------------------------------------------------------\n")
	for _, d := range domains {
		//fmt.Printf("%d\t%s\t%x\n", d.ID, d.Name, d.UUID)
		rstate,rmaxmem,rmemory,rvirCpu,rcputime,err := l.DomainGetInfo(d)
		if err != nil{
			continue
		}
		fmt.Printf("%s\t%s\t%d\t%d\t%d\t%d\t%d\t\n",d.ID,d.Name,rstate,rmaxmem,rmemory,rvirCpu,rcputime)
	}


	if err := l.Disconnect(); err != nil {
		log.Fatal("failed to disconnect: %v", err)
	}
}
