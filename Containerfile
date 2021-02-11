FROM docker.io/busybox:latest

MAINTAINER Raúl Sevilla

COPY build/perfApp /usr/bin/perfApp
CMD perfApp
