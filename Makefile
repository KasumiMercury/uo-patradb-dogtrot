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

atlas-inspect:
	atlas schema inspect -u "ent://ent/schema" --dev-url "sqlite://file?mode=memory&_fk=1" -w
atlas-diff:
	atlas migrate diff $(ARG) \
	  --dir "file://ent/migrate/migrations/ent" \
	  --to "ent://ent/schema" \
	  --dev-url "docker://mysql/8/ent"
atlas-custom:
	atlas migrate new $(ARG) \
	  --dir "file://ent/migrate/migrations/custom"
integration:
	cd ent/migrate/migrations/ &&
	bash integrate_migration.sh
atlas-push:
	atlas migrate push patradb \
	  --dir "file://ent/migrate/migrations/integrate" \
	  --dev-url "docker://mysql/8/dev"
atlas-lint:
	atlas migrate lint \
	  --dir "file://ent/migrate/migrations/integrate" \
	  --dev-url "docker://mysql/8/dev" \
	  -w
atlas-hash-custom:
	atlas migrate hash \
	  --dir "file://ent/migrate/migrations/custom"

generate-sql:
	go run -mod=mod ./cmd/migration/main.go $(ARG)