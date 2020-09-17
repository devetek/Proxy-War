ENGINE	:= envoy
TYPE	:= static

include nginx/Makefile
include envoy/Makefile

.PHONY: gen-cert
gen-cert: .certval
	@ ./scripts/cert.sh

.PHONY: run-dev
run-dev: .validator
	@ cp -rf $(ENGINE)-docker-compose-$(TYPE).yaml docker-compose.yaml
	@ docker-compose down --remove-orphans
	@ docker-compose up

.PHONY: show-services
show-services: .validator
	@ docker-compose ps --all

.PHONY: show-log
show-log: .validator
	@ docker-compose logs -f

.PHONY: down-dev
down-dev: .validator
	@ docker-compose down --remove-orphans


.PHONY: .certval
.certval:
	$(eval WHICH_MKCERT := $(shell which mkcert))

	@ test -n "$(WHICH_MKCERT)" || sh -c 'echo "No mkcert binary, follow this link to install in mac https://github.com/FiloSottile/mkcert#mkcert" && exit 1'

.PHONY: .validator
.validator:
	$(eval WHICH_DOCKER := $(shell which docker))
	$(eval WHICH_COMPOSE := $(shell which docker-compose))

	@ test -n "$(WHICH_DOCKER)" || sh -c 'echo "No docker binary, follow this link to install in mac https://docs.docker.com/docker-for-mac/install/" && exit 1'
	@ test -n "$(WHICH_COMPOSE)" || sh -c 'echo "No compose binary, follow this link to install in mac https://docs.docker.com/docker-for-mac/install/" && exit 1' 
