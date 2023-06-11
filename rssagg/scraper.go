package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/internal/database"
)

// fetch different RSS feeds and download all of their blog posts at the same time, it should be long-running job, never return
func startScraping(
	db *database.Queries,
	concurrency int, // the number of goroutines to do the scraping RSS feeds with the same number
	timeBetweenRequest time.Duration, // time between each request to go scrape a new rss feed
) {
	log.Printf("Scraping on %v goroutines every %s duration.\n", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)
	// when start scraping, for-loop will fire immediately, and then it will wait for the interval on ticker
	for ; ; <-ticker.C {
		feedsToFetch, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("error fetching feeds: %v.\n", err)
			continue // just continue when error happens, because we want always run the scraping job
		}
		log.Printf("Found %d feeds to fetch.\n", len(feedsToFetch))

		wg := &sync.WaitGroup{}
		for _, feedToFetch := range feedsToFetch {
			wg.Add(1) // fetch the feed in a individually goroutine
			go scrapeFeed(db, wg, feedToFetch)
		}
		wg.Wait() // only go to next for-loop, when the scraping is completed

		log.Printf("Fetched %d feeds completed.\n", len(feedsToFetch))
	}

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feedToFetch database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feedToFetch.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v\n.", err)
	}

	// fetching the feed
	rssFeed, err := urlToFeed(feedToFetch.Url)
	if err != nil {
		log.Printf("Error fetching feed: %v.\n", err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		log.Printf("Found post %s on feed %s.\n", item.Title, feedToFetch.Name)
	}

	log.Printf("Feed %s collected, %d posts found.\n", feedToFetch.Name, len(rssFeed.Channel.Items))

}
