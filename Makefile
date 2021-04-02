-include .env
export

start-db:
	docker-compose up -d

run: start-db
	go run *.go --token=$(TWITTER_BEARER_TOKEN) --query=$(QUERY)

log-postgres:
	$(DOCKER_COMPOSE) exec postgres tail -f /tmp/postgresql.log -n 100
