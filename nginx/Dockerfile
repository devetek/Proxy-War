FROM nginx:alpine

WORKDIR /code

# Service requirement
COPY nginx/service/requirements.txt .
COPY nginx/service/main.py .
COPY nginx/version.md .
COPY nginx/process/supervisor.conf /etc/supervisor.conf

# Nginx requirement
COPY nginx/default.conf /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/nginx.conf

# Exporter requirement
COPY nginx/access-log-exporter .

RUN apk update; \
  apk add python3 bash curl supervisor; \
  pip3 install --upgrade pip; \
  pip3 install -r requirements.txt; \
  mkdir -p /var/log/supervisor; \
  mkdir -p /etc/supervisor/conf.d; \
  mkdir -p /var/log/access-log-exporter/;

CMD ["supervisord", "-c", "/etc/supervisor.conf"]
