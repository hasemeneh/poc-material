pkgs          = $(shell go list ./... | grep -vE '(vendor|mock)')
NOW=$(shell date)
compose_file=docker/docker-compose.yml
compose=docker-compose -f ${compose_file}
poc_login_service_binary=poc-login
order_service_binary=order
warehouse_service_binary=warehouse
current_dir=$(shell pwd)



docker-start:
	@echo "${NOW} STARTING CONTAINER..."
	@${compose} up -d --build


docker-rebuilddb-poc-login:
	@echo "${NOW} REBUILDDB..."
	@echo "${NOW} DROPING EXISTING DB..."
	@docker exec -it poc-login-db  mysql -u root -e 'drop database if exists `poc-login-db`'
	@echo "${NOW} CREATE DB..."
	@docker exec -it poc-login-db  mysql -u root -e 'create database `poc-login-db`'
	@echo "${NOW} RUN SQL SCRIPTS..."
	@cd database/migration && go run main.go

docker-stop:
	@echo "${NOW} STOPPING CONTAINER..."
	@${compose} down
	@echo "${NOW} CLEAN UP..."
	@rm -f ./bin/poc-login/${poc_login_service_binary} 

docker-run-rest:
	@echo "${NOW} BUILDING..."
	@go mod vendor && cd ./cmd/rest  && go build -o ./../../bin/${poc_login_service_binary}
	@echo "${NOW} RUNNING..."
	@docker exec -it poc-login /usr/local/bin/${poc_login_service_binary}

mock:
	mockgen --source=pkg/webservice/server.type.go --destination=pkg/webservice/mocks/server.type.go --package mocks
	mockgen --source=pkg/router/router.go --destination=pkg/router/mocks/router.go --package mocks
