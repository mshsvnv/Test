create_db_command := psql -U postgres -d tests -f ./sql/init/init.sql
delete_db_command := psql -U postgres -d tests -f ./sql/drop.sql

init-db:
	$(create_db_command)

rm-db:
	$(delete_db_command)

run-local:
	go run cmd/main.go

gen-swag:
	swag init --parseDependency  --parseInternal --parseDepth 1 -g main.go -dir ./internal/controller/v2/http --instanceName v2 -o ./docs/v2

fmt-swag:
	swag f

convert-swag-to-3.0:
	chmod 755 ./docs/to3.sh && ./docs/to3.sh

swagger: fmt-swag gen-swag convert-swag-to-3.0

run-docker:
	docker build -t go_env -f Dockerfile.env .
	docker compose up -d

rm-docker:
	docker compose down

run-e2e: run-docker
	go test -v e2e/login_test.go
	# go test -v e2e/reset_test.go
