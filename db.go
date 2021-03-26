package main

import (
	"fmt"
	"github.com/g8rswimmer/go-twitter"
	"time"
)

func saveSearchResults(lookups twitter.TweetLookups, query string) {
	countTweets := 0
	countUsers := 0

	for _, lookup := range lookups {
		if saveTweet(lookup.Tweet, query) {
			countTweets++
		}
		if lookup.User != nil {
			if saveUser(*lookup.User) {
				countUsers++
			}
		}
	}

	fmt.Printf("%s: %d tweet(s) received, saved %d (%d new user(s))\n", time.Now().Format(time.RFC3339), len(lookups), countTweets, countUsers)
}

func saveTweet(tweet twitter.TweetObj, query string) bool {
	createdAt, err := time.Parse(time.RFC3339Nano, tweet.CreatedAt)
	if err != nil {
		fmt.Println("ERROR parsing tweet.CreatedAt: " + tweet.CreatedAt + " (id: " + tweet.ID + ")")
		return false
	}

	result, err := db.Exec(`
		INSERT INTO tweets
		(id, text, author_id, conversation_id, created_at, query)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT DO NOTHING
	`, tweet.ID, tweet.Text, tweet.AuthorID, tweet.ConversationID, createdAt, query)

	if err != nil {
		fmt.Println("ERROR saving tweet: " + tweet.ID)
		return false
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("ERROR saving tweet: " + tweet.ID)
		return false
	}

	if rows == 0 {
		return false
	} else {
		return true
	}
}

func saveUser(user twitter.UserObj) bool {
	createdAt, err := time.Parse(time.RFC3339Nano, user.CreatedAt)
	if err != nil {
		fmt.Println("ERROR parsing user.CreatedAt: " + user.CreatedAt + " (id: " + user.ID + ")")
		return false
	}

	result, err := db.Exec(`
		INSERT INTO users
		(id, username, name, created_at, followers_count, following_count, tweet_count)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT DO NOTHING
	`, user.ID, user.UserName, user.Name, createdAt, user.PublicMetrics.Followers, user.PublicMetrics.Following, user.PublicMetrics.Tweets)

	if err != nil {
		fmt.Println("ERROR saving user: " + user.ID)
		return false
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("ERROR saving user: " + user.ID)
		return false
	}

	if rows == 0 {
		return false
	} else {
		return true
	}
}
