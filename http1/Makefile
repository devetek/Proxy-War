.PHONY: http2-init
http1-init: .validator
	@go mod init github.com/devetek/error-hanlder

.PHONY: http2-build
http1-build:
	@docker build -f http1/Dockerfile -t prakasa1904/go-http1:latest .
	@docker push prakasa1904/go-http1:latest