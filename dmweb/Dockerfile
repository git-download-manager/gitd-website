FROM alpine:latest

LABEL service="gitd"

RUN mkdir -p /usr/local/dmweb

ARG ENV_FILE

COPY ${ENV_FILE} /usr/local/dmweb/.env
COPY assets /usr/local/dmweb/assets
COPY --from=gitd-builder /usr/local/go/bin/linux_amd64/dmweb /usr/local/dmweb/dmweb

WORKDIR /usr/local/dmweb

ENTRYPOINT [ "./dmweb",  "--env-file", ".env" ]