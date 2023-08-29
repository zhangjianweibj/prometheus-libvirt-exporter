package exporter

import (
	"encoding/xml"
	"time"

	"github.com/Knalltuete5000/prometheus-libvirt-exporter/libvirt_schema"
	"github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-libvirt/socket/dialers"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	libvirtUpDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "", "up"),
		"Whether scraping libvirt's metrics was successful.",
		[]string{"host"},
		nil)

	libvirtDomainNumbers = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "", "domains_number"),
		"Number of the domain",
		[]string{"host"},
		nil)

	libvirtDomainState = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "", "domain_state_code"),
		"Code of the domain state",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "stateDesc"},
		nil)

	libvirtDomainInfoMaxMemDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_info", "maximum_memory_bytes"),
		"Maximum allowed memory of the domain, in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"},
		nil)
	libvirtDomainInfoMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_info", "memory_usage_bytes"),
		"Memory usage of the domain, in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"},
		nil)
	libvirtDomainStatMemorySwapInBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_stat", "memory_swap_in_bytes"),
		"Memory swap in of domain(the total amount of data read from swap space), in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"},
		nil)
	libvirtDomainStatMemorySwapOutBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_stat", "memory_swap_out_bytes"),
		"Memory swap out of the domain(the total amount of memory written out to swap space), in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"},
		nil)
	libvirtDomainStatMemoryUnusedBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_stat", "memory_unused_bytes"),
		"Memory unused of the domain, in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"},
		nil)
	libvirtDomainStatMemoryAvailableInBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_stat", "memory_available_bytes"),
		"Memory available of the domain, in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"},
		nil)
	libvirtDomainStatMemoryUsableBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_stat", "memory_usable_bytes"),
		"Memory usable of the domain(corresponds to 'Available' in /proc/meminfo), in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"},
		nil)
	libvirtDomainStatMemoryRssBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_stat", "memory_rss_bytes"),
		"Resident Set Size of the process running the domain, in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"},
		nil)
	libvirtDomainInfoNrVirtCpuDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_info", "virtual_cpus"),
		"Number of virtual CPUs for the domain.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"},
		nil)
	libvirtDomainInfoCpuTimeDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_info", "cpu_time_seconds_total"),
		"Amount of CPU time used by the domain, in seconds.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host"},
		nil)

	libvirtDomainBlockRdBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_block_stats", "read_bytes_total"),
		"Number of bytes read from a block device, in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_file", "target_device"},
		nil)
	libvirtDomainBlockRdReqDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_block_stats", "read_requests_total"),
		"Number of read requests from a block device.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_file", "target_device"},
		nil)
	libvirtDomainBlockWrBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_block_stats", "write_bytes_total"),
		"Number of bytes written from a block device, in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_file", "target_device"},
		nil)
	libvirtDomainBlockWrReqDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_block_stats", "write_requests_total"),
		"Number of write requests from a block device.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_file", "target_device"},
		nil)

	//DomainInterface
	libvirtDomainInterfaceRxBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "receive_bytes_total"),
		"Number of bytes received on a network interface, in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceRxPacketsDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "receive_packets_total"),
		"Number of packets received on a network interface.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceRxErrsDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "receive_errors_total"),
		"Number of packet receive errors on a network interface.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceRxDropDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "receive_drops_total"),
		"Number of packet receive drops on a network interface.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "transmit_bytes_total"),
		"Number of bytes transmitted on a network interface, in bytes.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxPacketsDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "transmit_packets_total"),
		"Number of packets transmitted on a network interface.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxErrsDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "transmit_errors_total"),
		"Number of packet transmit errors on a network interface.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxDropDesc = prometheus.NewDesc(
		prometheus.BuildFQName("libvirt", "domain_interface_stats", "transmit_drops_total"),
		"Number of packet transmit drops on a network interface.",
		[]string{"domain", "instanceName", "instanceId", "flavorName", "userName", "userId", "projectName", "projectId", "host", "source_bridge", "target_device"},
		nil)

	domainState = map[libvirt_schema.DomainState]string{
		libvirt_schema.DOMAIN_NOSTATE:     "no state",
		libvirt_schema.DOMAIN_RUNNING:     "the domain is running",
		libvirt_schema.DOMAIN_BLOCKED:     "the domain is blocked on resource",
		libvirt_schema.DOMAIN_PAUSED:      "the domain is paused by user",
		libvirt_schema.DOMAIN_SHUTDOWN:    "the domain is being shut down",
		libvirt_schema.DOMAIN_SHUTOFF:     "the domain is shut off",
		libvirt_schema.DOMAIN_CRASHED:     "the domain is crashed",
		libvirt_schema.DOMAIN_PMSUSPENDED: "the domain is suspended by guest power management",
		libvirt_schema.DOMAIN_LAST:        "this enum value will increase over time as new events are added to the libvirt API",
	}
)

type collectFunc func(ch chan<- prometheus.Metric, l *libvirt.Libvirt, domain domainMeta, promLabels []string, logger log.Logger) (err error)

type domainMeta struct {
	domainName   string
	instanceName string
	instanceId   string
	flavorName   string

	userName string
	userId   string

	projectName string
	projectId   string

	libvirtDomain libvirt.Domain
	libvirtSchema libvirt_schema.Domain
}

// LibvirtExporter implements a Prometheus exporter for libvirt state.
type LibvirtExporter struct {
	uri    string
	driver libvirt.ConnectURI

	logger log.Logger
}

// NewLibvirtExporter creates a new Prometheus exporter for libvirt.
func NewLibvirtExporter(uri string, driver libvirt.ConnectURI, logger log.Logger) (*LibvirtExporter, error) {
	return &LibvirtExporter{
		uri:    uri,
		driver: driver,
		logger: logger,
	}, nil
}

// DomainFromLibvirt retrives all domains from the libvirt socket and enriches them with some meta information.
func DomainsFromLibvirt(l *libvirt.Libvirt, logger log.Logger) ([]domainMeta, error) {
	domains, _, err := l.ConnectListAllDomains(1, 0)
	if err != nil {
		_ = level.Error(logger).Log("err", "failed to load domain")
		return nil, err
	}

	lvDomains := make([]domainMeta, len(domains))
	for idx, domain := range domains {
		xmlDesc, err := l.DomainGetXMLDesc(domain, 0)
		if err != nil {
			_ = level.Error(logger).Log("failed to DomainGetXMLDesc")
			return nil, err
		}
		var libvirtSchema libvirt_schema.Domain
		if err = xml.Unmarshal([]byte(xmlDesc), &libvirtSchema); err != nil {
			_ = level.Error(logger).Log("failed to Unmarshal domains")
			return nil, err
		}

		lvDomains[idx].libvirtDomain = domain
		lvDomains[idx].libvirtSchema = libvirtSchema

		lvDomains[idx].domainName = domain.Name
		lvDomains[idx].instanceName = libvirtSchema.Metadata.NovaInstance.Name
		lvDomains[idx].instanceId = libvirtSchema.UUID
		lvDomains[idx].flavorName = libvirtSchema.Metadata.NovaInstance.Flavor.FlavorName

		lvDomains[idx].userName = libvirtSchema.Metadata.NovaInstance.Owner.User.UserName
		lvDomains[idx].userId = libvirtSchema.Metadata.NovaInstance.Owner.User.UserId

		lvDomains[idx].projectName = libvirtSchema.Metadata.NovaInstance.Owner.Project.ProjectName
		lvDomains[idx].projectId = libvirtSchema.Metadata.NovaInstance.Owner.Project.ProjectId
	}

	return lvDomains, nil
}

// Collect scrapes Prometheus metrics from libvirt.
func (e *LibvirtExporter) Collect(ch chan<- prometheus.Metric) {
	if err := CollectFromLibvirt(ch, e.uri, e.driver, e.logger); err != nil {
		_ = level.Error(e.logger).Log("err", "failed to collect metrics")
	}
}

// CollectFromLibvirt obtains Prometheus metrics from all domains in a
// libvirt setup.
func CollectFromLibvirt(ch chan<- prometheus.Metric, uri string, driver libvirt.ConnectURI, logger log.Logger) (err error) {
	// var conn net.Conn
	// if conn, err = net.DialTimeout("unix", uri, 5*time.Second); err != nil {
	// 	level.Error(logger).Log("err", "failed to dial libvirt")
	// 	return err
	// }
	// defer conn.Close()

	dialer := dialers.NewLocal(dialers.WithSocket(uri), dialers.WithLocalTimeout((5 * time.Second)))
	l := libvirt.NewWithDialer(dialer)
	if err = l.ConnectToURI(driver); err != nil {
		_ = level.Error(logger).Log("err", "failed to connect")
		return err
	}

	defer func() {
		if err := l.Disconnect(); err != nil {
			_ = level.Error(logger).Log("err", "failed to disconnect")
		}
	}()

	var host string
	if host, err = l.ConnectGetHostname(); err != nil {
		_ = level.Error(logger).Log("err", "failed to get hostname")
		return err
	}

	ch <- prometheus.MustNewConstMetric(
		libvirtUpDesc,
		prometheus.GaugeValue,
		1.0,
		host)

	domains, err := DomainsFromLibvirt(l, logger)

	domainNumber := len(domains)
	ch <- prometheus.MustNewConstMetric(
		libvirtDomainNumbers,
		prometheus.GaugeValue,
		float64(domainNumber),
		host)

	for _, domain := range domains {
		if err = CollectDomain(ch, l, domain, logger); err != nil {
			_ = level.Error(logger).Log("err", "failed to Collect domains")
			return err
		}
	}
	return nil
}

// CollectDomain extracts Prometheus metrics from a libvirt domain.
func CollectDomain(ch chan<- prometheus.Metric, l *libvirt.Libvirt, domain domainMeta, logger log.Logger) (err error) {
	var host string
	if host, err = l.ConnectGetHostname(); err != nil {
		_ = level.Error(logger).Log("err", "failed to get hostname")
		return err
	}

	var rState uint8
	var rvirCpu uint16
	var rmaxmem, rmemory, rcputime uint64
	if rState, rmaxmem, rmemory, rvirCpu, rcputime, err = l.DomainGetInfo(domain.libvirtDomain); err != nil {
		_ = level.Error(logger).Log("err", "failed to get domainInfo")
		return err
	}

	promLabels := []string{
		domain.domainName,
		domain.instanceName,
		domain.instanceId,
		domain.flavorName,
		domain.userName,
		domain.userId,
		domain.projectName,
		domain.projectId,
		host}

	ch <- prometheus.MustNewConstMetric(libvirtDomainState, prometheus.GaugeValue, float64(rState), append(promLabels, domainState[libvirt_schema.DomainState(rState)])...)

	ch <- prometheus.MustNewConstMetric(libvirtDomainInfoMaxMemDesc, prometheus.GaugeValue, float64(rmaxmem)*1024, promLabels...)
	ch <- prometheus.MustNewConstMetric(libvirtDomainInfoMemoryDesc, prometheus.GaugeValue, float64(rmemory)*1024, promLabels...)
	ch <- prometheus.MustNewConstMetric(libvirtDomainInfoNrVirtCpuDesc, prometheus.GaugeValue, float64(rvirCpu), promLabels...)
	ch <- prometheus.MustNewConstMetric(libvirtDomainInfoCpuTimeDesc, prometheus.CounterValue, float64(rcputime)/1e9, promLabels...)

	var isActive int32
	if isActive, err = l.DomainIsActive(domain.libvirtDomain); err != nil {
		_ = level.Error(logger).Log("err", "failed to get active status of domain")
		return err
	}
	if isActive != 1 {
		_ = level.Info(logger).Log("err", "domain is not active")
		return nil
	}

	for _, collectFunc := range []collectFunc{CollectDomainBlockDeviceInfo, CollectDomainNetworkInfo, CollectDomainDomainStatInfo} {
		if err = collectFunc(ch, l, domain, promLabels, logger); err != nil {
			_ = level.Warn(logger).Log("warn", "failed to collect some domain info")
		}
	}

	return nil
}

func CollectDomainBlockDeviceInfo(ch chan<- prometheus.Metric, l *libvirt.Libvirt, domain domainMeta, promLabels []string, logger log.Logger) (err error) {
	// Report block device statistics.
	for _, disk := range domain.libvirtSchema.Devices.Disks {
		if disk.Device == "cdrom" || disk.Device == "fd" {
			continue
		}

		var rRdReq, rRdBytes, rWrReq, rWrBytes int64
		if rRdReq, rRdBytes, rWrReq, rWrBytes, _, err = l.DomainBlockStats(domain.libvirtDomain, disk.Target.Device); err != nil {
			_ = level.Warn(logger).Log("warn", "failed to get DomainBlockStats")
			return err
		}

		promDiskLabels := append(promLabels, disk.Source.File, disk.Target.Device)
		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockRdBytesDesc,
			prometheus.CounterValue,
			float64(rRdBytes),
			promDiskLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockRdReqDesc,
			prometheus.CounterValue,
			float64(rRdReq),
			promDiskLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockWrBytesDesc,
			prometheus.CounterValue,
			float64(rWrBytes),
			promDiskLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockWrReqDesc,
			prometheus.CounterValue,
			float64(rWrReq),
			promDiskLabels...)

	}
	return
}

func CollectDomainNetworkInfo(ch chan<- prometheus.Metric, l *libvirt.Libvirt, domain domainMeta, promLabels []string, logger log.Logger) (err error) {
	// Report network interface statistics.
	for _, iface := range domain.libvirtSchema.Devices.Interfaces {
		if iface.Target.Device == "" {
			continue
		}
		var rRxBytes, rRxPackets, rRxErrs, rRxDrop, rTxBytes, rTxPackets, rTxErrs, rTxDrop int64
		if rRxBytes, rRxPackets, rRxErrs, rRxDrop, rTxBytes, rTxPackets, rTxErrs, rTxDrop, err = l.DomainInterfaceStats(domain.libvirtDomain, iface.Target.Device); err != nil {
			_ = level.Warn(logger).Log("warn", "failed to get DomainInterfaceStats")
			return err
		}

		promInterfaceLabels := append(promLabels, iface.Source.Bridge, iface.Target.Device)
		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceRxBytesDesc,
			prometheus.CounterValue,
			float64(rRxBytes),
			promInterfaceLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceRxPacketsDesc,
			prometheus.CounterValue,
			float64(rRxPackets),
			promInterfaceLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceRxErrsDesc,
			prometheus.CounterValue,
			float64(rRxErrs),
			promInterfaceLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceRxDropDesc,
			prometheus.CounterValue,
			float64(rRxDrop),
			promInterfaceLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceTxBytesDesc,
			prometheus.CounterValue,
			float64(rTxBytes),
			promInterfaceLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceTxPacketsDesc,
			prometheus.CounterValue,
			float64(rTxPackets),
			promInterfaceLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceTxErrsDesc,
			prometheus.CounterValue,
			float64(rTxErrs),
			promInterfaceLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainInterfaceTxDropDesc,
			prometheus.CounterValue,
			float64(rTxDrop),
			promInterfaceLabels...)
	}
	return
}

func CollectDomainDomainStatInfo(ch chan<- prometheus.Metric, l *libvirt.Libvirt, domain domainMeta, promLabels []string, logger log.Logger) (err error) {
	//collect stat info
	var rStats []libvirt.DomainMemoryStat
	if rStats, err = l.DomainMemoryStats(domain.libvirtDomain, uint32(libvirt.DomainMemoryStatNr), 0); err != nil {
		_ = level.Warn(logger).Log("warn", "failed to get domainstat")
		return err
	}
	for _, stat := range rStats {
		switch stat.Tag {
		case int32(libvirt.DomainMemoryStatSwapIn):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainStatMemorySwapInBytesDesc,
				prometheus.GaugeValue,
				float64(stat.Val)*1024,
				promLabels...)
		case int32(libvirt.DomainMemoryStatSwapOut):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainStatMemorySwapOutBytesDesc,
				prometheus.GaugeValue,
				float64(stat.Val)*1024,
				promLabels...)
		case int32(libvirt.DomainMemoryStatUnused):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainStatMemoryUnusedBytesDesc,
				prometheus.GaugeValue,
				float64(stat.Val*1024),
				promLabels...)
		case int32(libvirt.DomainMemoryStatAvailable):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainStatMemoryAvailableInBytesDesc,
				prometheus.GaugeValue,
				float64(stat.Val*1024),
				promLabels...)
		case int32(libvirt.DomainMemoryStatUsable):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainStatMemoryUsableBytesDesc,
				prometheus.GaugeValue,
				float64(stat.Val*1024),
				promLabels...)
		case int32(libvirt.DomainMemoryStatRss):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainStatMemoryRssBytesDesc,
				prometheus.GaugeValue,
				float64(stat.Val*1024),
				promLabels...)
		}
	}
	return
}

// Describe returns metadata for all Prometheus metrics that may be exported.
func (e *LibvirtExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- libvirtUpDesc
	ch <- libvirtDomainNumbers

	//domain info
	ch <- libvirtDomainState
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

	//domain mem stat
	ch <- libvirtDomainStatMemorySwapInBytesDesc
	ch <- libvirtDomainStatMemorySwapOutBytesDesc
	ch <- libvirtDomainStatMemoryUnusedBytesDesc
	ch <- libvirtDomainStatMemoryAvailableInBytesDesc
	ch <- libvirtDomainStatMemoryUsableBytesDesc
	ch <- libvirtDomainStatMemoryRssBytesDesc
}
