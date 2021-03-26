package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/g8rswimmer/go-twitter"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func search(token, query string) {
	tweet := &twitter.Tweet{
		Authorizer: authorize{
			Token: token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	fieldOpts := twitter.TweetFieldOptions{
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldConversationID, twitter.TweetFieldAuthorID},
		Expansions:  []twitter.Expansion{twitter.ExpansionAuthorID},
		UserFields:  []twitter.UserField{twitter.UserFieldCreatedAt, twitter.UserFieldPublicMetrics},
	}

	nextToken := ""
MainLoop:
	for true {
		searchOpts := twitter.TweetRecentSearchOptions{
			MaxResult: 100,
			NextToken: nextToken,
		}
		recentSearch, err := tweet.RecentSearch(context.Background(), query, searchOpts, fieldOpts)
		var tweetErr *twitter.TweetErrorResponse
		switch {
		case errors.As(err, &tweetErr):
			printTweetError(tweetErr)
			break MainLoop
		case err != nil:
			fmt.Println(err)
			break MainLoop
		default:
			saveSearchResults(recentSearch.LookUps, query)
			if recentSearch.Meta.ResultCount == 0 {
				fmt.Println("Finished!")
				break MainLoop
			}
			if recentSearch.Meta.NextToken == "" {
				fmt.Println("Finished! (next token not found)")
				break MainLoop
			}
			nextToken = recentSearch.Meta.NextToken
		}
	}
}

func printTweetError(tweetErr *twitter.TweetErrorResponse) {
	enc, err := json.MarshalIndent(tweetErr, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(enc))
}
