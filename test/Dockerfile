FROM debian:jessie
MAINTAINER Jeremy Udit <jcudit@gmail.com>

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get --quiet update && \
    apt-get install --yes --no-install-recommends socat

CMD socat -d -d TUN:10.0.0.1/30,iff-up,iff-no-pi UNIX-LISTEN:/dev/net/gotcp-tun,fork
