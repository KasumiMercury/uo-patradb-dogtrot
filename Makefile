build-dev:
	docker compose build
build-dev-no-cache:
	docker compose build --no-cache
up:
	docker compose up -d
down:
	docker compose down --remove-orphans
ps:
	docker compose ps

build-prod:
	docker build --target runner -t patradb-dogtrot-prod:latest .
run-prod:
	docker run -it --rm -p 8080:8080 --name patradb-dogtrot-prod patradb-dogtrot-prod:latest
stop-prod:
	docker stop patradb-dogtrot-prod


bash-dev:
	docker compose exec dogtrot bash
air:
	docker compose exec dogtrot air -c .air.toml