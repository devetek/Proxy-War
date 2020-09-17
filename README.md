How to handle error pages for frontend microservices architecture nginx vs envoy.

# How it Works

The idea is we're implementing nginx error recursive to response user when error happen. Preparing location error handler and fill it with return static HTML error page when error service also down!.

# Development

```sh
make run-dev
```

# Test

Open https://devel.tokopedia.com/error-400, https://devel.tokopedia.com/error-401, https://devel.tokopedia.com/error-403, https://devel.tokopedia.com/error-404, https://devel.tokopedia.com/error-500, https://devel.tokopedia.com/error-502, https://devel.tokopedia.com/error-503, https://devel.tokopedia.com/error-504 for test service error handler. Open https://devel.tokopedia.com/totally-error for test default when all of the service down!.