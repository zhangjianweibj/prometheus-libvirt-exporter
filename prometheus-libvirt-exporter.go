package main

import (
	"encoding/xml"
	"flag"
	"github.com/digitalocean/go-libvirt"
	"github.com/prometheus-libvirt-exporter/libvirt_schema"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	libvirtUpDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "", "up"),
		"Whether scraping libvirt's metrics was successful.",
		nil,
		nil)

	libvirtDomainNumbers = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt","","domains_number"),
		"number of the domain",[]string{"number"},
		nil)

	libvirtDomainInfoMaxMemDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_info", "maximum_memory_bytes"),
		"Maximum allowed memory of the domain, in bytes.",
		[]string{"domainName", "instanceName", "instanceId"},
		nil)
	libvirtDomainInfoMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_info", "memory_usage_bytes"),
		"Memory usage of the domain, in bytes.",
		[]string{"domainName", "instanceName", "instanceId"},
		nil)
	libvirtDomainInfoNrVirtCpuDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_info", "virtual_cpus"),
		"Number of virtual CPUs for the domain.",
		[]string{"domainName", "instanceName", "instanceId"},
		nil)
	libvirtDomainInfoCpuTimeDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_info", "cpu_time_seconds_total"),
		"Amount of CPU time used by the domain, in seconds.",
		[]string{"domainName", "instanceName", "instanceId"},
		nil)

	libvirtDomainBlockRdBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_block_stats", "read_bytes_total"),
		"Number of bytes read from a block device, in bytes.",
		[]string{"domainName", "instanceName", "instanceId", "source_file", "target_device"},
		nil)
	libvirtDomainBlockRdReqDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_block_stats", "read_requests_total"),
		"Number of read requests from a block device.",
		[]string{"domainName", "instanceName", "instanceId", "source_file", "target_device"},
		nil)
	libvirtDomainBlockWrBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_block_stats", "write_bytes_total"),
		"Number of bytes written from a block device, in bytes.",
		[]string{"domainName", "instanceName", "instanceId", "source_file", "target_device"},
		nil)
	libvirtDomainBlockWrReqDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_block_stats", "write_requests_total"),
		"Number of write requests from a block device.",
		[]string{"domainName", "instanceName", "instanceId", "source_file", "target_device"},
		nil)

	//DomainInterface
	libvirtDomainInterfaceRxBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "receive_bytes_total"),
		"Number of bytes received on a network interface, in bytes.",
		[]string{"domainName", "instanceName", "instanceId", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceRxPacketsDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "receive_packets_total"),
		"Number of packets received on a network interface.",
		[]string{"domainName", "instanceName", "instanceId", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceRxErrsDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "receive_errors_total"),
		"Number of packet receive errors on a network interface.",
		[]string{"domainName", "instanceName", "instanceId", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceRxDropDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "receive_drops_total"),
		"Number of packet receive drops on a network interface.",
		[]string{"domainName", "instanceName", "instanceId", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "transmit_bytes_total"),
		"Number of bytes transmitted on a network interface, in bytes.",
		[]string{"domainName", "instanceName", "instanceId", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxPacketsDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "transmit_packets_total"),
		"Number of packets transmitted on a network interface.",
		[]string{"domainName", "instanceName", "instanceId", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxErrsDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "transmit_errors_total"),
		"Number of packet transmit errors on a network interface.",
		[]string{"domainName", "instanceName", "instanceId", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxDropDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "transmit_drops_total"),
		"Number of packet transmit drops on a network interface.",
		[]string{"domainName", "instanceName", "instanceId", "source_bridge", "target_device"},
		nil)
)

// CollectDomain extracts Prometheus metrics from a libvirt domain.
func CollectDomain(ch chan<- prometheus.Metric, l *libvirt.Libvirt, domain *libvirt.Domain) error {
	xmlDesc, err := l.DomainGetXMLDesc(*domain, 0)
	if err != nil {
		log.Fatalf("failed to DomainGetXMLDesc: %v", err)
		return err
	}
	var libvirtSchema libvirt_schema.Domain
	err = xml.Unmarshal([]byte(xmlDesc), &libvirtSchema)
	if err != nil {
		log.Fatalf("failed to Unmarshal domains: %v", err)
		return err
	}

	domainName := domain.Name
	instanceName := libvirtSchema.Metadata.NovaInstance.Name
	instanceId := libvirtSchema.UUID

	_, rmaxmem, rmemory, rvirCpu, rcputime, err := l.DomainGetInfo(*domain)

	if err != nil {
		log.Fatalf("failed to get domainInfo: %v", err)
		return err
	}
	ch <- prometheus.MustNewConstMetric(
		libvirtDomainInfoMaxMemDesc,
		prometheus.GaugeValue,
		float64(rmaxmem)*1024,
		domainName, instanceName, instanceId)
	ch <- prometheus.MustNewConstMetric(
		libvirtDomainInfoMemoryDesc,
		prometheus.GaugeValue,
		float64(rmemory)*1024,
		domainName, instanceName, string(instanceId[:]))
	ch <- prometheus.MustNewConstMetric(
		libvirtDomainInfoNrVirtCpuDesc,
		prometheus.GaugeValue,
		float64(rvirCpu),
		domainName, instanceName, string(instanceId[:]))
	ch <- prometheus.MustNewConstMetric(
		libvirtDomainInfoCpuTimeDesc,
		prometheus.CounterValue,
		float64(rcputime)/1e9,
		domainName, instanceName, string(instanceId[:]))

	// Report block device statistics.
	for _, disk := range libvirtSchema.Devices.Disks {
		if disk.Device == "cdrom" || disk.Device == "fd" {
			continue
		}

		isActive, err := l.DomainIsActive(*domain)
		var rRdReq, rRdBytes, rWrReq, rWrBytes int64
		if isActive == 1 {
			rRdReq, rRdBytes, rWrReq, rWrBytes, _, err = l.DomainBlockStats(*domain, disk.Target.Device)
		}
		if err != nil {
			log.Fatalf("failed to get DomainBlockStats: %v", err)
			return err
		}

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockRdBytesDesc,
			prometheus.CounterValue,
			float64(rRdBytes),
			domainName, instanceName, string(instanceId[:]),
			disk.Source.File,
			disk.Target.Device)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockRdReqDesc,
			prometheus.CounterValue,
			float64(rRdReq),
			domainName, instanceName, string(instanceId[:]),
			disk.Source.File,
			disk.Target.Device)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockWrBytesDesc,
			prometheus.CounterValue,
			float64(rWrBytes),
			domainName, instanceName, string(instanceId[:]),
			disk.Source.File,
			disk.Target.Device)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockWrReqDesc,
			prometheus.CounterValue,
			float64(rWrReq),
			domainName, instanceName, string(instanceId[:]),
			disk.Source.File,
			disk.Target.Device)

	}

	// Report network interface statistics.
	for _, iface := range libvirtSchema.Devices.Interfaces {
		if iface.Target.Device == "" {
			continue
		}
		isActive, err := l.DomainIsActive(*domain)
		var rRxBytes, rRxPackets, rRxErrs, rRxDrop, rTxBytes, rTxPackets, rTxErrs, rTxDrop int64
		if isActive == 1 {
			rRxBytes, rRxPackets, rRxErrs, rRxDrop, rTxBytes, rTxPackets, rTxErrs, rTxDrop, err = l.DomainInterfaceStats(*domain, iface.Target.Device)
		}

		if err != nil {
			log.Fatalf("failed to get DomainInterfaceStats: %v", err)
			return err
		}

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceRxBytesDesc,
			prometheus.CounterValue,
			float64(rRxBytes),
			domainName, instanceName, string(instanceId[:]),
			iface.Source.Bridge,
			iface.Target.Device)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceRxPacketsDesc,
			prometheus.CounterValue,
			float64(rRxPackets),
			domainName, instanceName, string(instanceId[:]),
			iface.Source.Bridge,
			iface.Target.Device)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceRxErrsDesc,
			prometheus.CounterValue,
			float64(rRxErrs),
			domainName, instanceName, string(instanceId[:]),
			iface.Source.Bridge,
			iface.Target.Device)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceRxDropDesc,
			prometheus.CounterValue,
			float64(rRxDrop),
			domainName, instanceName, string(instanceId[:]),
			iface.Source.Bridge,
			iface.Target.Device)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceTxBytesDesc,
			prometheus.CounterValue,
			float64(rTxBytes),
			domainName, instanceName, string(instanceId[:]),
			iface.Source.Bridge,
			iface.Target.Device)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceTxPacketsDesc,
			prometheus.CounterValue,
			float64(rTxPackets),
			domainName, instanceName, string(instanceId[:]),
			iface.Source.Bridge,
			iface.Target.Device)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceTxErrsDesc,
			prometheus.CounterValue,
			float64(rTxErrs),
			domainName, instanceName, string(instanceId[:]),
			iface.Source.Bridge,
			iface.Target.Device)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceTxDropDesc,
			prometheus.CounterValue,
			float64(rTxDrop),
			domainName, instanceName, string(instanceId[:]),
			iface.Source.Bridge,
			iface.Target.Device)

	}

	return nil
}

// CollectFromLibvirt obtains Prometheus metrics from all domains in a
// libvirt setup.
func CollectFromLibvirt(ch chan<- prometheus.Metric, uri string) error {
	conn, err := net.DialTimeout("unix", uri, 5*time.Second)

	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
		return err
	}
	defer conn.Close()

	l := libvirt.New(conn)
	if err := l.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
		return err
	}

	domains, err := l.Domains()
	if err != nil {
		log.Fatalf("failed to load domain: %v", err)
		return err
	}

	//domains number
	domainNumber := len(domains)
	ch <- prometheus.MustNewConstMetric(
		libvirtDomainNumbers,
		prometheus.GaugeValue,
		float64(domainNumber))

	for _, domain := range domains {
		err = CollectDomain(ch, l, &domain)
		l.DomainShutdown(domain)
		//domain.Free()
		if err != nil {
			log.Fatalf("failed to Collect domain: %v", err)
			return err
		}
	}
	return nil
}

// LibvirtExporter implements a Prometheus exporter for libvirt state.
type LibvirtExporter struct {
	uri string
}

// NewLibvirtExporter creates a new Prometheus exporter for libvirt.
func NewLibvirtExporter(uri string) (*LibvirtExporter, error) {
	return &LibvirtExporter{
		uri: uri,
	}, nil
}

// Describe returns metadata for all Prometheus metrics that may be exported.
func (e *LibvirtExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- libvirtUpDesc
	ch <- libvirtDomainNumbers

	//domain info
	ch <- libvirtDomainInfoMaxMemDesc
	ch <- libvirtDomainInfoMemoryDesc
	ch <- libvirtDomainInfoNrVirtCpuDesc
	ch <- libvirtDomainInfoCpuTimeDesc

	//domain block
	ch <- libvirtDomainBlockRdBytesDesc
	ch <- libvirtDomainBlockRdReqDesc
	ch <- libvirtDomainBlockWrBytesDesc
	ch <- libvirtDomainBlockWrReqDesc

	//domain interface
	ch <- libvirtDomainInterfaceRxBytesDesc
	ch <- libvirtDomainInterfaceRxPacketsDesc
	ch <- libvirtDomainInterfaceRxErrsDesc
	ch <- libvirtDomainInterfaceRxDropDesc
	ch <- libvirtDomainInterfaceTxBytesDesc
	ch <- libvirtDomainInterfaceTxPacketsDesc
	ch <- libvirtDomainInterfaceTxErrsDesc
	ch <- libvirtDomainInterfaceTxDropDesc

}

// Collect scrapes Prometheus metrics from libvirt.
func (e *LibvirtExporter) Collect(ch chan<- prometheus.Metric) {
	err := CollectFromLibvirt(ch, e.uri)
	if err == nil {
		ch <- prometheus.MustNewConstMetric(
			libvirtUpDesc,
			prometheus.GaugeValue,
			1.0)
	} else {
		log.Printf("Failed to scrape metrics: %s", err)
		ch <- prometheus.MustNewConstMetric(
			libvirtUpDesc,
			prometheus.GaugeValue,
			0.0)
	}
}

func main() {
	var (
		listenAddress = flag.String("web.listen-address", ":9000", "Address to listen on for web interface and telemetry.")
		metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
		libvirtURI    = flag.String("libvirt.uri", "/var/run/libvirt/libvirt-sock", "Libvirt URI from which to extract metrics.")
	)
	flag.Parse()

	exporter, err := NewLibvirtExporter(*libvirtURI)
	if err != nil {
		panic(err)
	}
	prometheus.MustRegister(exporter)

	http.Handle(*metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
			<head><title>Libvirt Exporter</title></head>
			<body>
			<h1>Libvirt Exporter</h1>
			<p><a href='` + *metricsPath + `'>Metrics</a></p>
			</body>
			</html>`))
	})
	log.Fatal(http.ListenAndServe(*listenAddress, nil))

}
