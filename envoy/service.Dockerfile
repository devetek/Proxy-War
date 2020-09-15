FROM python:alpine

WORKDIR /code

COPY envoy/service/requirements.txt .
COPY envoy/service/service.py .
COPY envoy/service/start_service.sh /usr/local/bin/start_service.sh

RUN apk update; \
  apk add curl; \
  pip install -r requirements.txt; \
  chmod u+x /usr/local/bin/start_service.sh;

ENTRYPOINT /usr/local/bin/start_service.sh
