package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/koron/go-spafs"

	"github.com/cheeze2000/pastelia/backend/internal/postgres"
	"github.com/cheeze2000/pastelia/backend/internal/snippet"
)

func createHandler(dbpool *pgxpool.Pool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("MODE") == "development" {
			w.Header().Add("Access-Control-Allow-Origin", "*")
		}

		switch r.Method {
		case http.MethodGet:
			qs := r.URL.Query()
			if !qs.Has("q") {
				http.NotFound(w, r)
				return
			}

			snippet, err := postgres.ReadSnippet(dbpool, qs.Get("q"))
			if err != nil {
				http.NotFound(w, r)
				return
			}

			w.Header().Add("Content-Type", "application/json")
			w.Write([]byte(snippet.String()))
		case http.MethodPost:
			var snippet snippet.Snippet
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&snippet)
			if err != nil || snippet.Lang == "" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}

			slug, err := postgres.CreateSnippet(dbpool, &snippet)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Write([]byte(slug))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Unable to connect to database")
	}

	frontend := spafs.FileServer(http.Dir("../frontend/dist"))
	go http.ListenAndServe(":3006", frontend)

	backend := http.NewServeMux()
	backend.HandleFunc("/", createHandler(dbpool))
	http.ListenAndServe(":3009", backend)
}
