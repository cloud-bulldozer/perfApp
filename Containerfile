FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

MAINTAINER Raúl Sevilla

COPY build/perfApp /usr/bin/perfApp
CMD perfApp
