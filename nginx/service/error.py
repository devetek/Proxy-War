from flask import Response, Flask, request
from helpers import graphs, count_process, prometheus_client, time

app = Flask(__name__)

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def catch_all(path):
  return 'Hi guest, are you lost in page %s' % path

if __name__ == "__main__":
  app.run(host='0.0.0.0', port=8081, debug=True)