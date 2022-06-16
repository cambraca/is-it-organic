-include .env
export

start-db:
	docker-compose up -d

run: start-db
	docker-compose run --rm go go run *.go --token=$(TWITTER_BEARER_TOKEN) --query=$(QUERY)

log-postgres:
	docker-compose exec postgres tail -f /tmp/postgresql.log -n 100
