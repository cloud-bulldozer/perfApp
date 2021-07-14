FROM docker.io/busybox:latest

MAINTAINER Raúl Sevilla

ARG ARCH=amd64
COPY build/perfApp-${ARCH} /usr/bin/perfApp
CMD perfApp
