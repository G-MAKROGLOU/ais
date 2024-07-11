run-local:
ifeq ($(OS),Windows_NT)
	@set AIS_ENV=local&& air -c .air.win.toml
else
	@AIS_ENV=local air -c .air.nix.toml
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
	
build:
	docker build --tag "latest" -f ./.docker/Dockerfile.prod --rm --no-cache .
.PHONY: build-prod

