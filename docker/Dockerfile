from golang:1.12-alpine3.9 as builder


arg LIBVIRT_REPO=https://github.com/zhangjianweibj/prometheus-libvirt-exporter.git
arg LIBVIRT_BRANCH=master

env GOPATH=/go

run apk add --no-cache \
        git

run mkdir -p $GOPATH/src/github.com/zhangjianweibj/prometheus-libvirt-exporter && \
  cd $GOPATH/src/github.com/zhangjianweibj/prometheus-libvirt-exporter && \
  git init && \
  git remote add origin $LIBVIRT_REPO && \
  git fetch origin $LIBVIRT_BRANCH && \
  git reset --hard FETCH_HEAD && \
  go build -o prometheus-libvirt-exporter


from alpine:3.9

ENV \
    WEB_LISTEN_ADDRESS=web.listen-address \
    WEB_TELEMETRY_PATH=/metrics \
    LIBVIRT_URI=/var/run/libvirt/libvirt-sock

copy --from=builder /go/src/github.com/zhangjianweibj/prometheus-libvirt-exporter/prometheus-libvirt-exporter /

entrypoint ["/bin/sh", "-c", "/prometheus-libvirt-exporter"]
cmd ["-web.listen-address", "$WEB_LISTEN_ADDRESS", "-web.telemetry-path", "$WEB_TELEMETRY_PATH", "libvirt.uri", "$LIBVIRT_URI" ]
