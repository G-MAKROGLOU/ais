run-local:
ifeq ($(OS),Windows_NT)
	@set AIS_ENV=local&& go run ./cmd/ais/main.go
else
	@AIS_ENV=local go run ./cmd/ais/main.go
endif

docker-start:
	docker compose -f ./.docker/docker-compose.yml -p ais-api up

docker-stop:
	docker compose -f ./.docker/docker-compose.yml -p ais-api down --rmi="all"

