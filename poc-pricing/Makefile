pkgs          = $(shell go list ./... | grep -vE '(vendor|mock)')
NOW=$(shell date)
compose_file=docker/docker-compose.yml
compose=docker-compose -f ${compose_file}
poc_pricing_service_binary=poc-pricing
order_service_binary=order
warehouse_service_binary=warehouse
current_dir=$(shell pwd)


run-dummy:
	go run dummy/frontend/main.go

docker-start:
	@echo "${NOW} STARTING CONTAINER..."
	@${compose} up -d --build


docker-rebuilddb-poc-pricing:
	@echo "${NOW} REBUILDDB..."
	@echo "${NOW} DROPING EXISTING DB..."
	@docker exec -it poc-pricing-db  mysql -u root -e 'drop database if exists `poc-pricing-db`'
	@echo "${NOW} CREATE DB..."
	@docker exec -it poc-pricing-db  mysql -u root -e 'create database `poc-pricing-db`'
	@echo "${NOW} RUN SQL SCRIPTS..."
	@cd database/migration && go run main.go

docker-stop:
	@echo "${NOW} STOPPING CONTAINER..."
	@${compose} down
	@echo "${NOW} CLEAN UP..."
	@rm -f ./bin/poc-pricing/${poc_pricing_service_binary} 

docker-run-rest:
	@echo "${NOW} BUILDING..."
	@go mod vendor && cd ./cmd/rest  && go build -o ./../../bin/${poc_pricing_service_binary}
	@echo "${NOW} RUNNING..."
	@docker exec -it poc-pricing /usr/local/bin/${poc_pricing_service_binary}

mock:
	mockgen --source=pkg/webservice/server.type.go --destination=pkg/webservice/mocks/server.type.go --package mocks
	mockgen --source=pkg/router/router.go --destination=pkg/router/mocks/router.go --package mocks
