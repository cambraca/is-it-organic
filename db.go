package main

import (
	"fmt"
	"github.com/g8rswimmer/go-twitter"
	"time"
)

func saveSearchResults(recentSearch *twitter.TweetRecentSearch, query string) {
	for _, lookup := range recentSearch.LookUps {
		saveTweet(lookup.Tweet, query)
		if lookup.User != nil {
			saveUser(*lookup.User)
		}
	}
}

func saveTweet(tweet twitter.TweetObj, query string) {
	createdAt, err := time.Parse(time.RFC3339Nano, tweet.CreatedAt)
	if err != nil {
		fmt.Println("ERROR parsing tweet.CreatedAt: " + tweet.CreatedAt + " (id: " + tweet.ID + ")")
		return
	}

	_, err = db.Exec(`
		INSERT INTO tweets
		(id, text, author_id, conversation_id, created_at, query)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT DO NOTHING
	`, tweet.ID, tweet.Text, tweet.AuthorID, tweet.ConversationID, createdAt, query)

	if err != nil {
		fmt.Println("ERROR saving tweet: " + tweet.ID)
	}
}

func saveUser(user twitter.UserObj) {
	createdAt, err := time.Parse(time.RFC3339Nano, user.CreatedAt)
	if err != nil {
		fmt.Println("ERROR parsing user.CreatedAt: " + user.CreatedAt + " (id: " + user.ID + ")")
		return
	}

	_, err = db.Exec(`
		INSERT INTO users
		(id, username, name, created_at, followers_count, following_count, tweet_count)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT DO NOTHING
	`, user.ID, user.UserName, user.Name, createdAt, user.PublicMetrics.Followers, user.PublicMetrics.Following, user.PublicMetrics.Tweets)

	if err != nil {
		fmt.Println("ERROR saving user: " + user.ID)
	}
}
