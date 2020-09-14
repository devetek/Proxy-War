.PHONY: envoy-build-deploy
envoy-build-deploy: .validator
	@ docker build -f envoy/Dockerfile -t prakasa1904/envoy-error-fallback .
	@ docker push prakasa1904/envoy-error-fallback