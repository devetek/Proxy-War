How to handle error pages for frontend microservices architecture nginx vs envoy.

# How it Works

The idea is we're implementing nginx error recursive to response user when error happen. Preparing location error handler and fill it with return static HTML error page when error service also down!.

# Development

```sh
make run-dev
```

# Test

Open http://localhost/error-400, http://localhost/error-401, http://localhost/error-403, http://localhost/error-404, http://localhost/error-500, http://localhost/error-502, http://localhost/error-503, http://localhost/error-504 for test service error handler. Open http://localhost/totally-error for test default when all of the service down!.