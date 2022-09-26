FROM alpine:latest

LABEL service="gitd"

RUN mkdir -p /usr/local/dmcli

COPY sh /usr/local/dmcli/sh
COPY *.dev.yaml /usr/local/dmcli/
COPY --from=gitd-builder /usr/local/go/bin/linux_amd64/dmcli /usr/local/dmcli/dmcli

WORKDIR /usr/local/dmcli

ENTRYPOINT [ "./dmcli" ]