.PHONY: envoy-build-deploy
envoy-build-deploy: .validator
	@ docker build -f envoy/service.Dockerfile -t prakasa1904/error-service .
	@ docker push prakasa1904/error-service
	@ docker build -f envoy/envoy.Dockerfile -t prakasa1904/envoy-error-fallback .
	@ docker push prakasa1904/envoy-error-fallback