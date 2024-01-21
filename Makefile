DOCKER_POSTGRES_CONTAINER_NAME=api_postges
POSTGRES_HOST=localhost
POSTGRES_PORT=5431
POSTGRES_USER=realtemirov
POSTGRES_PASSWORD=123456
POSTGRES_DB=task-for-dell

##########################################################################
# Golang
run:
	go run ./cmd/main.go
build:
	go build -o ./bin/main ./cmd/main.go
test:
	go test -cover ./...

##########################################################################
# Postgres Container
psql-run: 
	docker run --name ${DOCKER_POSTGRES_CONTAINER_NAME} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_DB=${POSTGRES_DB} -p ${POSTGRES_PORT}:5432 -d postgres
psql-start: 
	docker start ${DOCKER_POSTGRES_CONTAINER_NAME}
psql-login:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} psql ${POSTGRES_DB} ${POSTGRES_USER} 
psql-stop: 
	docker stop ${DOCKER_POSTGRES_CONTAINER_NAME}
psql-down: 
	docker rm -f ${DOCKER_POSTGRES_CONTAINER_NAME}


##########################################################################
# Migration
migration-up:
	migrate -path ./migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable' up
migration-down:
	migrate -path ./migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable' down
createdb:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}
dropdb:
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} dropdb --username=${POSTGRES_USER} ${POSTGRES_DB}

	
##########################################################################
# Swagger
swag:
	swag init -g **/**/*.go

##########################################################################
# Docker
compose:
	docker-compose up -d --build
start:
	make psql-run && make migration-up