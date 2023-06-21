package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/api"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/internal/database"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/routes"
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/scraper"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // database driver postgresql
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("POSTGRES_URL")
	if dbURL == "" {
		log.Fatal("POSTGRES_URL is not found in the environment")
	}

	conn, dbErr := sql.Open("postgres", dbURL)
	if dbErr != nil {
		log.Fatal("Can't connect to database:", dbErr)
	}

	db := database.New(conn)
	apiCfg := api.ApiConfig{
		DB: db,
	}

	// run aggregation worker forever
	go scraper.StartScraping(db, 10, time.Minute)

	// setup router
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// setup routes
	routesCaller := routes.New(&apiCfg)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", routesCaller.Healthz())
	v1Router.Get("/err", routesCaller.Err())
	// users
	v1Router.Post("/users", routesCaller.CreateUser())
	v1Router.Get("/users", routesCaller.GetUser())
	// feeds
	v1Router.Post("/feeds", routesCaller.CreateFeed())
	v1Router.Get("/feeds", routesCaller.GetFeeds())
	// posts
	v1Router.Get("/posts", routesCaller.GetPostsForUser())
	// feed_follows
	v1Router.Post("/feed_follows", routesCaller.CreateFeedFollow())
	v1Router.Get("/feed_follows", routesCaller.GetFeedFollows())
	v1Router.Delete("/feed_follows/{feedFollowID}", routesCaller.DeleteFeedFollow())

	// setup server
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}

	fmt.Println("Hello, Welcome to RSS Aggregator!")
	fmt.Printf("Server starting at port: %v\n", portStr)
	srvErr := srv.ListenAndServe()
	if srvErr != nil {
		log.Fatal(srvErr)
	}

}
