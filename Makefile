run:
	docker build -t go_env:latest . -f env.dockerfile
	docker compose up prometheus grafana echo-server gin-server -d

rerun-gatling:
	docker stop gatling-at-once gatling-per-second && docker rm gatling-at-once gatling-per-second
	docker compose up gatling-at-once gatling-per-second

rm:
	docker compose down
	docker image rm go_env:latest

.PHONY: run rerun-gatling rm