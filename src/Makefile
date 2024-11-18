run:
	docker build -t go_env -f Dockerfile.env .
	docker compose up -d

ci-e2e:
	export ALLURE_OUTPUT_PATH="${GITHUB_WORKSPACE}" && \
	export ALLURE_OUTPUT_FOLDER="e2e-allure" && \
	go test -v -tags=e2e ./... --race

local-e2e:
	go clean -testcache
	docker build -t go_env -f Dockerfile.env .
	docker compose up -d
	export ALLURE_OUTPUT_PATH="/c/Users/mshsv/OneDrive/Рабочий стол/studying/7_sem/Test" && \
	go test -v -tags=e2e ./... --race
	docker compose down

rm-docker:
	docker compose down