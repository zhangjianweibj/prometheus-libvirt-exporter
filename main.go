package main

import (
	"net/http"
	"os"

	kingpin "github.com/alecthomas/kingpin/v2"
	"github.com/digitalocean/go-libvirt"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
	webflag "github.com/prometheus/exporter-toolkit/web/kingpinflag"
)

func main() {
	metricsPath := kingpin.Flag(
		"web.telemetry-path", "Path under which to expose metrics",
	).Default("/metrics").String()
	toolkitFlags := webflag.AddFlags(kingpin.CommandLine, ":9000")

	promlogConfig := &promlog.Config{}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	kingpin.Version(version.Print("libvirt_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	logger := promlog.New(promlogConfig)

	level.Info(logger).Log("msg", "Starting libvirt_exporter", "version", version.Info())
	level.Info(logger).Log("msg", "Build context", "build_context", version.BuildContext())

	exporter, err := NewLibvirtExporter(*libvirtURI, libvirt.ConnectURI(*driver), logger)
	if err != nil {
		panic(err)
	}
	prometheus.MustRegister(exporter)

	http.Handle(*metricsPath, promhttp.Handler())
	if *metricsPath != "/" {
		landingCnf := web.LandingConfig{
			Name:        "Libvirt Exporter",
			Description: "Prometheus Libvirt Exporter",
			Version:     version.Info(),
			Links: []web.LandingLinks{
				{
					Address: *metricsPath,
					Text:    "Metrics",
				},
			},
		}
		landingPage, err := web.NewLandingPage(landingCnf)
		if err != nil {
			level.Error(logger).Log("err", err)
			os.Exit(1)
		}
		http.Handle("/", landingPage)
	}

	srv := &http.Server{}
	if err = web.ListenAndServe(srv, toolkitFlags, logger); err != nil {
		level.Error(logger).Log("err", err)
		os.Exit(1)
	}
}
