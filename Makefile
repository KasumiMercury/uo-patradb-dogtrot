build-dev:
	docker build -t patradb-dogtrot-dev:latest .
run-dev:
	docker run -it --rm -p 8080:8080 --name patradb-dogtrot-dev patradb-dogtrot-dev:latest
build-prod:
	docker build -t patradb-dogtrot-prod:latest -f Dockerfile.prod .
run-prod:
	docker run -it --rm -p 8080:8080 --name patradb-dogtrot-prod patradb-dogtrot-prod:latest
