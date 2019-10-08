APPLICATION_NAME    := prometheus-libvirt-exporter
APPLICATION_VERSION := $(shell cat VERSION)

LDFLAGS := -X $(APPLICATION_NAME)/meta.Version=$(APPLICATION_VERSION)

GO_BUILD := go build -v -ldflags "$(LDFLAGS)"

PATH := $(PWD)/build/dependencies:$(PATH)

TEMPDIR := $(shell mktemp -d)

all: build 

build_linux:
	env GOOS=linux GOARCH=amd64 GO111MODULE=on $(GO_BUILD) -o build/prometheus-libvirt-exporter ./prometheus-libvirt-exporter.go

build:
	$(GO_BUILD) -o build/prometheus-libvirt-exporter ./prometheus-libvirt-exporter.go

clean:
	go clean -v .
	rm -rf build

FPM-exists:
	@fpm -v || \
	(echo >&2 "FPM must be installed on the system. See https://github.com/jordansissel/fpm"; false)

deb: FPM-exists build_linux
	mkdir -p dist/$(APPLICATION_VERSION)/
	cd dist/$(APPLICATION_VERSION)/ && \
	fpm -s dir \
	    -t deb \
        -n prometheus-libvirt-exporter \
        -v $(APPLICATION_VERSION) \
        --url="https://github.com/zhangjianweibj/prometheus-libvirt-exporter" \
        --description "Prometheus-libvirt-exporter service (host and vm metrics exposed for prometheus, written in Go with pluggable metric collectors)" \
        --deb-priority optional \
        --workdir $(TEMPDIR) \
        --architecture amd64 \
        ../../build/prometheus-libvirt-exporter=/usr/bin/prometheus-libvirt-exporter \
        ../../debian/prometheus-libvirt-exporter.service=/etc/systemd/system/prometheus-libvirt-exporter.service \
        ../../debian/prometheus-libvirt-exporter.upstart=/etc/init/prometheus-libvirt-exporter.conf