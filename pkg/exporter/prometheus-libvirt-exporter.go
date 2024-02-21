package exporter

import (
	"encoding/xml"
	"time"

	"github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-libvirt/socket/dialers"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/inovex/prometheus-libvirt-exporter/libvirt_schema"
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "libvirt"

var (
	libvirtUpDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Whether scraping libvirt's metrics was successful.",
		nil,
		nil)

	libvirtDomainNumbers = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "domains"),
		"Number of domains",
		nil,
		nil)

	//domain info
	libvirtDomainState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_info", "state"),
		"Code of the domain state",
		[]string{"domain", "state_desc"},
		nil)
	libvirtDomainInfoMaxMemDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_info", "maximum_memory_bytes"),
		"Maximum allowed memory of the domain, in bytes.",
		[]string{"domain"},
		nil)
	libvirtDomainInfoMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_info", "memory_usage_bytes"),
		"Memory usage of the domain, in bytes.",
		[]string{"domain"},
		nil)
	libvirtDomainInfoNrVirtCpuDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_info", "virtual_cpus"),
		"Number of virtual CPUs for the domain.",
		[]string{"domain"},
		nil)
	libvirtDomainInfoCpuTimeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_info", "cpu_time_seconds_total"),
		"Amount of CPU time used by the domain, in seconds.",
		[]string{"domain"},
		nil)

	//domain memory stats
	libvirtDomainMemoryStatsSwapInBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_memory_stats", "swap_in_bytes"),
		"Memory swapped in for this domain(the total amount of data read from swap space)",
		[]string{"domain"},
		nil)
	libvirtDomainMemoryStatsSwapOutBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_memory_stats", "swap_out_bytes"),
		"Memory swapped out for this domain (the total amount of memory written out to swap space)",
		[]string{"domain"},
		nil)
	libvirtDomainMemoryStatsUnusedBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_memory_stats", "unused_bytes"),
		"Memory unused by the domain",
		[]string{"domain"},
		nil)
	libvirtDomainMemoryStatsAvailableInBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_memory_stats", "available_bytes"),
		"Memory available to the domain",
		[]string{"domain"},
		nil)
	libvirtDomainMemoryStatsUsableBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_memory_stats", "usable_bytes"),
		"Memory usable by the domain (corresponds to 'Available' in /proc/meminfo)",
		[]string{"domain"},
		nil)
	libvirtDomainMemoryStatsRssBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_memory_stats", "rss_bytes"),
		"Resident Set Size of the process running the domain",
		[]string{"domain"},
		nil)

	//domain block stats
	libvirtDomainBlockStatsRdBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_block_stats", "read_bytes_total"),
		"Number of bytes read from a block device, in bytes.",
		[]string{"domain", "source_file", "target_device"},
		nil)
	libvirtDomainBlockStatsRdReqDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_block_stats", "read_requests_total"),
		"Number of read requests from a block device.",
		[]string{"domain", "source_file", "target_device"},
		nil)
	libvirtDomainBlockStatsWrBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_block_stats", "write_bytes_total"),
		"Number of bytes written from a block device, in bytes.",
		[]string{"domain", "source_file", "target_device"},
		nil)
	libvirtDomainBlockStatsWrReqDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_block_stats", "write_requests_total"),
		"Number of write requests from a block device.",
		[]string{"domain", "source_file", "target_device"},
		nil)

	//domain interface stats
	libvirtDomainInterfaceRxBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_interface_stats", "receive_bytes_total"),
		"Number of bytes received on a network interface, in bytes.",
		[]string{"domain", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceRxPacketsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_interface_stats", "receive_packets_total"),
		"Number of packets received on a network interface.",
		[]string{"domain", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceRxErrsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_interface_stats", "receive_errors_total"),
		"Number of packet receive errors on a network interface.",
		[]string{"domain", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceRxDropDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_interface_stats", "receive_drops_total"),
		"Number of packet receive drops on a network interface.",
		[]string{"domain", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_interface_stats", "transmit_bytes_total"),
		"Number of bytes transmitted on a network interface, in bytes.",
		[]string{"domain", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxPacketsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_interface_stats", "transmit_packets_total"),
		"Number of packets transmitted on a network interface.",
		[]string{"domain", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxErrsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_interface_stats", "transmit_errors_total"),
		"Number of packet transmit errors on a network interface.",
		[]string{"domain", "source_bridge", "target_device"},
		nil)
	libvirtDomainInterfaceTxDropDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain_interface_stats", "transmit_drops_total"),
		"Number of packet transmit drops on a network interface.",
		[]string{"domain", "source_bridge", "target_device"},
		nil)

	// info metrics
	libvirtDomainInfoDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain", "info"),
		"Metadata labels for the domain.",
		[]string{"domain", "os_type", "os_type_arch", "os_type_machine"},
		nil)

	// info metrics from metadata extracted OpenStack Nova
	libvirtDomainOpenstackInfoDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "domain", "openstack_info"),
		"OpenStack Metadata labels for the domain.",
		[]string{"domain", "instance_name", "instance_id", "flavor_name", "user_name", "user_id", "project_name", "project_id"},
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
	domainName      string
	instanceName    string
	instanceId      string
	flavorName      string
	os_type_arch    string
	os_type_machine string
	os_type         string

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
		_ = level.Error(logger).Log("err", "failed to load domains", "msg", err)
		return nil, err
	}

	lvDomains := make([]domainMeta, len(domains))
	for idx, domain := range domains {
		xmlDesc, err := l.DomainGetXMLDesc(domain, 0)
		if err != nil {
			_ = level.Error(logger).Log("err", "failed to DomainGetXMLDesc", "domain", domain.Name, "msg", err)
			continue
		}
		var libvirtSchema libvirt_schema.Domain
		if err = xml.Unmarshal([]byte(xmlDesc), &libvirtSchema); err != nil {
			_ = level.Error(logger).Log("err", "failed to unmarshal domain", "domain", domain.Name, "msg", err)
			continue
		}

		lvDomains[idx].libvirtDomain = domain
		lvDomains[idx].libvirtSchema = libvirtSchema

		lvDomains[idx].domainName = domain.Name
		lvDomains[idx].instanceName = libvirtSchema.Metadata.NovaInstance.Name
		lvDomains[idx].instanceId = libvirtSchema.UUID
		lvDomains[idx].flavorName = libvirtSchema.Metadata.NovaInstance.Flavor.FlavorName
		lvDomains[idx].os_type_arch = libvirtSchema.OSMetadata.Type.Arch
		lvDomains[idx].os_type_machine = libvirtSchema.OSMetadata.Type.Machine
		lvDomains[idx].os_type = libvirtSchema.OSMetadata.Type.Value

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
		_ = level.Error(e.logger).Log("err", "failed to collect metrics", "msg", err)
	}
}

// CollectFromLibvirt obtains Prometheus metrics from all domains in a libvirt setup.
func CollectFromLibvirt(ch chan<- prometheus.Metric, uri string, driver libvirt.ConnectURI, logger log.Logger) (err error) {
	dialer := dialers.NewLocal(dialers.WithSocket(uri), dialers.WithLocalTimeout((5 * time.Second)))
	l := libvirt.NewWithDialer(dialer)
	if err = l.ConnectToURI(driver); err != nil {
		_ = level.Error(logger).Log("err", "failed to connect", "msg", err)
		return err
	}

	defer func() {
		if err := l.Disconnect(); err != nil {
			_ = level.Error(logger).Log("err", "failed to disconnect", "msg", err)
		}
	}()

	ch <- prometheus.MustNewConstMetric(
		libvirtUpDesc,
		prometheus.GaugeValue,
		1.0)

	domains, err := DomainsFromLibvirt(l, logger)
	if err != nil {
		_ = level.Error(logger).Log("err", "failed to retrieve domains from Libvirt", "msg", err)
		return err
	}

	domainNumber := len(domains)
	ch <- prometheus.MustNewConstMetric(
		libvirtDomainNumbers,
		prometheus.GaugeValue,
		float64(domainNumber))

	// collect domain metrics from libvirt
	// see https://libvirt.org/html/libvirt-libvirt-domain.html
	for _, domain := range domains {
		if err = CollectDomain(ch, l, domain, logger); err != nil {
			_ = level.Error(logger).Log("err", "failed to collect domain", "domain", domain.domainName, "msg", err)
			return err
		}
	}
	return nil
}

// CollectDomain extracts Prometheus metrics from a libvirt domain.
func CollectDomain(ch chan<- prometheus.Metric, l *libvirt.Libvirt, domain domainMeta, logger log.Logger) (err error) {

	var rState uint8
	var rvirCpu uint16
	var rmaxmem, rmemory, rcputime uint64
	if rState, rmaxmem, rmemory, rvirCpu, rcputime, err = l.DomainGetInfo(domain.libvirtDomain); err != nil {
		_ = level.Error(logger).Log("err", "failed to get domainInfo", "domain", domain.libvirtDomain.Name, "msg", err)
		return err
	}

	promLabels := []string{
		domain.domainName,
	}

	openstackInfoLabels := []string{
		domain.domainName,
		domain.instanceName,
		domain.instanceId,
		domain.flavorName,
		domain.userName,
		domain.userId,
		domain.projectName,
		domain.projectId,
	}

	infoLabels := []string{
		domain.domainName,
		domain.os_type,
		domain.os_type_arch,
		domain.os_type_machine,
	}

	ch <- prometheus.MustNewConstMetric(libvirtDomainInfoDesc, prometheus.GaugeValue, 1.0, infoLabels...)
	ch <- prometheus.MustNewConstMetric(libvirtDomainOpenstackInfoDesc, prometheus.GaugeValue, 1.0, openstackInfoLabels...)

	ch <- prometheus.MustNewConstMetric(libvirtDomainState, prometheus.GaugeValue, float64(rState), append(promLabels, domainState[libvirt_schema.DomainState(rState)])...)

	ch <- prometheus.MustNewConstMetric(libvirtDomainInfoMaxMemDesc, prometheus.GaugeValue, float64(rmaxmem)*1024, promLabels...)
	ch <- prometheus.MustNewConstMetric(libvirtDomainInfoMemoryDesc, prometheus.GaugeValue, float64(rmemory)*1024, promLabels...)
	ch <- prometheus.MustNewConstMetric(libvirtDomainInfoNrVirtCpuDesc, prometheus.GaugeValue, float64(rvirCpu), promLabels...)
	ch <- prometheus.MustNewConstMetric(libvirtDomainInfoCpuTimeDesc, prometheus.CounterValue, float64(rcputime)/1e9, promLabels...)

	var isActive int32
	if isActive, err = l.DomainIsActive(domain.libvirtDomain); err != nil {
		_ = level.Error(logger).Log("err", "failed to get active status of domain", "domain", domain.libvirtDomain.Name, "msg", err)
		return err
	}
	if isActive != 1 {
		_ = level.Debug(logger).Log("debug", "domain is not active, skipping", "domain", domain.libvirtDomain.Name)
		return nil
	}

	for _, collectFunc := range []collectFunc{CollectDomainBlockDeviceInfo, CollectDomainNetworkInfo, CollectDomainDomainStatInfo} {
		if err = collectFunc(ch, l, domain, promLabels, logger); err != nil {
			_ = level.Warn(logger).Log("warn", "failed to collect some domain info", "domain", domain.libvirtDomain.Name, "msg", err)
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
			_ = level.Warn(logger).Log("warn", "failed to get DomainBlockStats", "domain", domain.libvirtDomain.Name, "msg", err)
			return err
		}

		promDiskLabels := append(promLabels, disk.Source.File, disk.Target.Device)
		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockStatsRdBytesDesc,
			prometheus.CounterValue,
			float64(rRdBytes),
			promDiskLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockStatsRdReqDesc,
			prometheus.CounterValue,
			float64(rRdReq),
			promDiskLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockStatsWrBytesDesc,
			prometheus.CounterValue,
			float64(rWrBytes),
			promDiskLabels...)

		ch <- prometheus.MustNewConstMetric(
			libvirtDomainBlockStatsWrReqDesc,
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
			_ = level.Warn(logger).Log("warn", "failed to get DomainInterfaceStats", "domain", domain.libvirtDomain.Name, "msg", err)
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
		_ = level.Warn(logger).Log("warn", "failed to get DomainMemoryStats", "domain", domain.libvirtDomain.Name, "msg", err)
		return err
	}
	for _, stat := range rStats {
		switch stat.Tag {
		case int32(libvirt.DomainMemoryStatSwapIn):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainMemoryStatsSwapInBytesDesc,
				prometheus.GaugeValue,
				float64(stat.Val)*1024,
				promLabels...)
		case int32(libvirt.DomainMemoryStatSwapOut):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainMemoryStatsSwapOutBytesDesc,
				prometheus.GaugeValue,
				float64(stat.Val)*1024,
				promLabels...)
		case int32(libvirt.DomainMemoryStatUnused):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainMemoryStatsUnusedBytesDesc,
				prometheus.GaugeValue,
				float64(stat.Val*1024),
				promLabels...)
		case int32(libvirt.DomainMemoryStatAvailable):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainMemoryStatsAvailableInBytesDesc,
				prometheus.GaugeValue,
				float64(stat.Val*1024),
				promLabels...)
		case int32(libvirt.DomainMemoryStatUsable):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainMemoryStatsUsableBytesDesc,
				prometheus.GaugeValue,
				float64(stat.Val*1024),
				promLabels...)
		case int32(libvirt.DomainMemoryStatRss):
			ch <- prometheus.MustNewConstMetric(
				libvirtDomainMemoryStatsRssBytesDesc,
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

	ch <- libvirtDomainInfoDesc
	ch <- libvirtDomainOpenstackInfoDesc

	//domain info
	ch <- libvirtDomainState
	ch <- libvirtDomainInfoMaxMemDesc
	ch <- libvirtDomainInfoMemoryDesc
	ch <- libvirtDomainInfoNrVirtCpuDesc
	ch <- libvirtDomainInfoCpuTimeDesc

	//domain block
	ch <- libvirtDomainBlockStatsRdBytesDesc
	ch <- libvirtDomainBlockStatsRdReqDesc
	ch <- libvirtDomainBlockStatsWrBytesDesc
	ch <- libvirtDomainBlockStatsWrReqDesc

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
	ch <- libvirtDomainMemoryStatsSwapInBytesDesc
	ch <- libvirtDomainMemoryStatsSwapOutBytesDesc
	ch <- libvirtDomainMemoryStatsUnusedBytesDesc
	ch <- libvirtDomainMemoryStatsAvailableInBytesDesc
	ch <- libvirtDomainMemoryStatsUsableBytesDesc
	ch <- libvirtDomainMemoryStatsRssBytesDesc
}
