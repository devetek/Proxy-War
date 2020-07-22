from flask import Response, Flask, request
from helpers import graphs, count_process, prometheus_client, time

app = Flask(__name__)


@app.route("/metrics")
def requests_count():
    res = []
    for _,value in graphs.items():
      res.append(prometheus_client.generate_latest(value))
    
    return Response(res, mimetype="text/plain")

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def catch_all(path):
    count_process()
    return 'Hi guest, welcome to %s' % path

if __name__ == "__main__":
  app.run(host='0.0.0.0', port=8080, debug=True)