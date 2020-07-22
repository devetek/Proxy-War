from flask import Response, Flask, request
from helpers import graphs, count_process, prometheus_client, time

app = Flask(__name__)

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def catch_all(path):
  headers = request.headers
  response = Response('Error', status=500)
  response.headers['X-App-Id'] = 'general-error'
  print(headers)
  return response

if __name__ == "__main__":
  app.run(host='0.0.0.0', port=8081, debug=True)