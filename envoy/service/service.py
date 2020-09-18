# from flask import Flask
# from flask import request
from quart import make_response, Quart, render_template, url_for
import os
import requests
import socket
# from slackclient import SlackClient

app = Quart(__name__)

TRACE_HEADERS_TO_PROPAGATE = [
    'X-Ot-Span-Context',
    'X-Request-Id',

    # Zipkin headers
    'X-B3-TraceId',
    'X-B3-SpanId',
    'X-B3-ParentSpanId',
    'X-B3-Sampled',
    'X-B3-Flags',

    # Jaeger header (for native client)
    "uber-trace-id"
]

if 'SERVICE_NAME' not in os.environ:
  os.environ['SERVICE_NAME'] = 'default'

@app.route('/')
async def index():
  return ('Hello from behind Envoy (service {})! hostname: {} resolved'
          'hostname: {}\n'.format(os.environ['SERVICE_NAME'], socket.gethostname(),
                                  socket.gethostbyname(socket.gethostname())))

# @app.route('/service/<service_number>')
# def hello(service_number):
#   return ('Hello from behind Envoy (service {})! hostname: {} resolved'
#           'hostname: {}\n'.format(os.environ['SERVICE_NAME'], socket.gethostname(),
#                                   socket.gethostbyname(socket.gethostname())))


# @app.route('/trace/<service_number>')
# def trace(service_number):
#   headers = {}

#   # call service atreus from service ares
#   if os.environ['SERVICE_NAME'] == 'ares':
#     for header in TRACE_HEADERS_TO_PROPAGATE:
#       if header in request.headers:
#         headers[header] = request.headers[header]
#         requests.get("http://localhost:9000/trace/", headers=headers)
#   return ('Hello from behind Envoy (service {})! hostname: {} resolved'
#           'hostname: {}\n'.format(os.environ['SERVICE_NAME'], socket.gethostname(),
#                                   socket.gethostbyname(socket.gethostname())))


if __name__ == "__main__":
  app.run(host='0.0.0.0', port=8080, debug=True)