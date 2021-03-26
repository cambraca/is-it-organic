-include .env
export

run:
	go run *.go --token=$(TWITTER_BEARER_TOKEN) --query=$(QUERY)
