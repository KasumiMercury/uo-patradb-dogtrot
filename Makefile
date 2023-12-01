build-dev:
	docker build --target dev -t patradb-dogtrot-dev:latest .
run-dev:
	docker run -it --rm -p 8080:8080 -p 2345:2345 --name patradb-dogtrot-dev -d patradb-dogtrot-dev:latest
stop-dev:
	docker stop patradb-dogtrot-dev
build-prod:
	docker build --target runner -t patradb-dogtrot-prod:latest .
run-prod:
	docker run -it --rm -p 8080:8080 --name patradb-dogtrot-prod patradb-dogtrot-prod:latest
stop-prod:
	docker stop patradb-dogtrot-prod

bash-dev:
	docker exec -it patradb-dogtrot-dev bash
air:
	docker exec -it patradb-dogtrot-dev air -c .air.toml