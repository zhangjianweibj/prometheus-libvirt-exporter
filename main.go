package main

import (
	"flag"
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	exp "github.com/zhangjianweibj/prometheus-libvirt-exporter/pkg/prometheus-libvirt-exporter"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var (
		listenAddress = flag.String("web.listen-address", ":9000", "Address to listen on for web interface and telemetry.")
		metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
		libvirtURI    = flag.String("libvirt.uri", "/var/run/libvirt/libvirt-sock-ro", "Libvirt URI from which to extract metrics.")
		driver        = flag.String("libvirt.driver", string(libvirt.QEMUSystem), fmt.Sprintf("Available drivers: %s (Default), %s, %s and %s ", libvirt.QEMUSystem, libvirt.QEMUSession, libvirt.XenSystem, libvirt.TestDefault))
	)
	flag.Parse()

	exporter, err := exp.NewLibvirtExporter(
		*libvirtURI,
		libvirt.ConnectURI(*driver),
		exp.WithLogger(logger),
	)
	if err != nil {
		panic(err)
	}
	prometheus.MustRegister(exporter)

	http.Handle(*metricsPath, promhttp.Handler())
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
	if err = http.ListenAndServe(*listenAddress, nil); err != nil {
		logger.Error("unexpected server shutdown", zap.Error(err))
	}
}
