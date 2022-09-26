FROM golang:1.18-alpine AS builder

ARG SERVICE_BUILD
ARG SERVICE_COMMIT_ID

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /usr/local/go/src/{gitd,gitdm}/

COPY internal /usr/local/go/src/gitdm/internal

WORKDIR /usr/local/go/src/gitd

COPY . .

RUN ls .

RUN go install -v -ldflags="-w -s -X main.ServiceBuild=${SERVICE_BUILD} -X main.ServiceCommitId=${SERVICE_COMMIT_ID}" ./...