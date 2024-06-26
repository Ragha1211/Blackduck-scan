.PHONY: build run stop clean

DOCKER_COMPOSE := docker-compose.yaml

build:
	docker-compose -f $(DOCKER_COMPOSE) build
	go build .

run:
	docker-compose -f $(DOCKER_COMPOSE) up -d

run-go:
	bash ./go-run.sh

stop-go:
	bash ./go-stop.sh
    
stop:
	docker-compose -f $(DOCKER_COMPOSE) down

clean: stop
	docker-compose -f $(DOCKER_COMPOSE) rm -f