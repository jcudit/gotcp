FROM golang:1.7
MAINTAINER Jeremy Udit <jcudit@gmail.com>

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get --quiet update && \
    apt-get install --yes --no-install-recommends net-tools netcat-openbsd
