run-local:
ifeq ($(OS),Windows_NT)
	@set AIS_ENV=local&& go run ./cmd/ais/main.go
else
	@AIS_ENV=local go run ./cmd/ais/main.go
endif

docker-start:
ifeq ($(OS),Windows_NT)
	@set PORT=3005&& docker compose -f ./.docker/docker-compose.yml -p ais-api up
else
	@PORT=3005 docker compose -f ./.docker/docker-compose.yml -p ais-api up
endif
	
docker-stop:
ifeq ($(OS),Windows_NT)
	@set PORT=3005&& docker compose -f ./.docker/docker-compose.yml -p ais-api down --rmi="all"
else
	@PORT=3005 docker compose -f ./.docker/docker-compose.yml -p ais-api down --rmi="all"
endif
	

