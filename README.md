# Is It Organic

Small utility written in Go to search for tweets and store data about them and their authors in Postgres.

Hopefully useful to detect patterns in the early stages of a viral hashtag, for example.

## Limitations

The Twitter API endpoint used apparently won't give you results older than a week or so.

## Requirements

* Go 1.15
* Docker (for the Postgres database)

## Instructions

1. Set up a Twitter developer account and a new application. You should get a bearer token.
2. Create an `.env` file and add the token. The file should have one line and look like this:

   ```
   TWITTER_BEARER_TOKEN=...
   ```

3. Run `$ docker-compose up -d` to start the Postgres database server.
4. Using any Postgres client, execute the queries from `schema.sql`. This will create two tables. Here is the connection string:

   ```
   host=localhost port=10593 user=postgres password=postgres dbname=postgres sslmode=disable
   ```

5. Run `$ make run QUERY=#MyHashtag`. This will search until there are no more results and store everything in the database.
6. Run SQL queries to get useful data. The following will get data about the users, sorted by the date of their first tweet with the hashtag:

   ```SQL
   SELECT MIN(t.created_at) AS first_tweet_at,
       u.username,
       EXTRACT(DAYS FROM NOW() - u.created_at) AS days_since_account_was_created,
       u.followers_count,
       u.following_count,
       u.tweet_count
   FROM tweets t
       JOIN users u ON t.author_id = u.id
   WHERE t.query = '#MyHashtag'
   GROUP BY u.id
   ORDER BY first_tweet_at;
   ```