SERVICE :=zeus

.PHONY: build-deploy
build-deploy: .validator
	@ docker build -f nginx/Dockerfile -t prakasa1904/nginx-service-export .
	@ docker push prakasa1904/nginx-service-export

.PHONY: run-dev
run-dev: .validator
	@ docker-compose down --remove-orphans
	@ docker-compose up -d

.PHONY: show-services
show-services: .validator
	@ docker-compose ps --all

.PHONY: show-log
show-log: .validator
	@ docker-compose logs -f

.PHONY: down-dev
down-dev: .validator
	@ docker-compose down --remove-orphans

.PHONY: .validator
.validator:
	$(eval WHICH_DOCKER := $(shell which docker))
	$(eval WHICH_COMPOSE := $(shell which docker-compose))

	@ test -n "$(WHICH_DOCKER)" || sh -c 'echo "No docker binary, follow this link to install in mac https://docs.docker.com/docker-for-mac/install/" && exit 1'
	@ test -n "$(WHICH_COMPOSE)" || sh -c 'echo "No compose binary, follow this link to install in mac https://docs.docker.com/docker-for-mac/install/" && exit 1' 
