.PHONY: nginx-build-deploy
nginx-build-deploy: .validator
	@ docker build -f nginx/Dockerfile -t prakasa1904/nginx-error-fallback .
	@ docker push prakasa1904/nginx-error-fallback