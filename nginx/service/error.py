from flask import Response, Flask, request, make_response

app = Flask(__name__)

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def catch_all(path):
  is_error = 'False'
  code_error = 401

  if 'X-Header-Error' in request.headers:
    is_error = request.headers['X-Header-Error']

  if 'X-Header-Status' in request.headers:
    code_error = request.headers['X-Header-Status']

  response = make_response('Hi guest, are you lost in page %s, with error status %s' % (path, code_error))
  response.status_code = 401
  response.headers['X-App-Name'] = 'tkp-error-app'

  # Default error handler simulation
  if is_error == 'True':
    response.status_code = 502

  return response

if __name__ == "__main__":
  app.run(host='0.0.0.0', port=8081, debug=True)