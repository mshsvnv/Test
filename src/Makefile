create_db_command := psql -U postgres -d tests -f ./sql/create.sql
delete_db_command := psql -U postgres -d tests -f ./sql/drop.sql

create_db:
	$(create_db_command)

delete_db:
	$(delete_db_command)

test:
	$(create_db_command)
	go clean -testcache
	# rm -rf ../allure-results
	# export ALLURE_OUTPUT_PATH="/c/Users/mshsv/OneDrive/Рабочий стол/studying/7_sem/Test" && go test -v ./... -shuffle=on --race --parallel 6
	# cp ../environment.properties ../allure-results
	go test -v ./... -shuffle=on --race --parallel 6
	$(delete_db_command)
	
allure:
	cp -R ../allure-reports/history ../allure-results
	rm -rf ../allure-reports
	allure generate ../allure-results -o ../allure-reports
	allure serve ../allure-results -p 4000

report: test allure

ci-unit:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="unit-allure" && \
	go test -v -tags=unit ./... --race

local-unit:
	go clean -testcache
	export ALLURE_OUTPUT_PATH="/c/Users/mshsv/OneDrive/Рабочий стол/studying/7_sem/Test" && \
	go test -v -tags=unit ./... --race

ci-integration:
	docker compose up -d
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="integration-allure" && \
 	export DB_INIT_PATH="${GITHUB_WORKSPACE}/src/sql/init/init.sql" && \
	go test -v -tags=integration ./... --race

local-integration:
	go clean -testcache
	docker compose up -d
	export ALLURE_OUTPUT_PATH="/c/Users/mshsv/OneDrive/Рабочий стол/studying/7_sem/Test" && \
 	export DB_INIT_PATH="/c/Users/mshsv/OneDrive/Рабочий стол/studying/7_sem/Test/src/sql/init/init.sql" && \
	go test -v -tags=integration ./... --race
	docker compose down

ci-e2e:
	docker compose up -d
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="e2e-allure" && \
	go test -v -tags=e2e ./... --race
	docker compose down
	docker image rm backend:latest bitnami/postgresql:16 alpine:latest

local-e2e:
	go clean -testcache
	docker compose up -d
	export ALLURE_OUTPUT_PATH="/c/Users/mshsv/OneDrive/Рабочий стол/studying/7_sem/Test" && \
	go test -v -tags=e2e ./... --race
	docker compose down

rm-docker:
	docker compose down

ci-concat-reports:
	mkdir ../allure-results
	cp unit-allure/* ../allure-results/
	# cp integration-allure/* ../allure-results/
	# cp e2e-allure/* allure-results/
	cp environment.properties allure-results

.PHONY: create_db delete_db test allure report