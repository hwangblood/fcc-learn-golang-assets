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
	"github.com/hwangblood/fcc-learn-golang-assets/rssagg/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // database driver
)

type apiConfig struct {
	DB *database.Queries
}

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
	apiCfg := apiConfig{
		DB: db,
	}

	// run aggregation worker forever
	go startScraping(db, 10, time.Minute)

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
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	// users
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.meddlewareAuth(apiCfg.handlerGetUser))
	// feeds
	v1Router.Post("/feeds", apiCfg.meddlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	// posts
	v1Router.Get("/posts", apiCfg.meddlewareAuth(apiCfg.handlerGetPostsForUser))
	// feed_follows
	v1Router.Post("/feed_follows", apiCfg.meddlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.meddlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.meddlewareAuth(apiCfg.handlerDeleteFeedFollow))

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
