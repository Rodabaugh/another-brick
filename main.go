package main

import (
	"another-brick/internal/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	platform  string
	db        *database.Queries
	siteTitle string
	subTitle  string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Using enviroment variables.")
	} else {
		fmt.Println("Loaded .env file.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	siteTitle := os.Getenv("SITE_TITLE")
	if siteTitle == "" {
		siteTitle = "Another Brick"
	}

	subTitle := os.Getenv("SUB_TITLE")
	if subTitle == "" {
		subTitle = "in the"
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	platform := os.Getenv("PLATFORM")
	if platform != "dev" && platform != "prod" {
		log.Fatal("PLATFORM must be set to either dev or prod")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		platform:  platform,
		db:        dbQueries,
		siteTitle: siteTitle,
		subTitle:  subTitle,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		MainPage(&apiCfg).Render(r.Context(), w)
	})

	mux.HandleFunc("POST /api/posts", apiCfg.handlerPostsCreate)
	mux.HandleFunc("GET /api/posts", apiCfg.handlerPostsGet)
	mux.HandleFunc("DELETE /api/posts/{post_id}", apiCfg.handlerPostsDelete)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting another-brick on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
