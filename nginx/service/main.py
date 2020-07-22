from flask import Response, Flask, request, make_response

app = Flask(__name__)

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def catch_all(path):
  return location_response(path)

def location_response(path):
  response = make_response('Hi guest, you\'re in page %s' % path)
  response.status_code = 200
  response.headers['X-App-Name'] = 'tkp-backend-app'

  if path == "error-400":
    response.status_code = 400
  elif path == "error-401":
    response.status_code = 401
  elif path == "error-403":
    response.status_code = 403
  elif path == "error-404":
    response.status_code = 404
  elif path == "error-500":
    response.status_code = 500
  elif path == "error-502":
    response.status_code = 502
  elif path == "error-503":
    response.status_code = 503
  elif path == "error-504":
    response.status_code = 504

  return response
    


if __name__ == "__main__":
  app.run(host='0.0.0.0', port=8080, debug=True)