FROM envoyproxy/envoy-dev:latest

RUN apt-get update; \
  apt-get install curl iputils-ping -y;