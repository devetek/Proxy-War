.PHONY: http2-init
http2-init: .validator
	@go mod init github.com/devetek/error-hanlder

.PHONY: http2-build
http2-build:
	@docker build -f http2/Dockerfile -t prakasa1904/go-http2:latest .
	@docker push prakasa1904/go-http2:latest