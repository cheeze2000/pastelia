package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/cheeze2000/pastelia/api/internal/postgres"
	"github.com/cheeze2000/pastelia/api/internal/snippet"
)

func createHandler(dbpool *pgxpool.Pool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("API_MODE") == "development" {
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
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@postgres:5432/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	fmt.Println(databaseUrl)

	dbpool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		panic("Unable to connect to database")
	}

	api := http.NewServeMux()
	api.HandleFunc("/", createHandler(dbpool))
	http.ListenAndServe(":3000", api)
}
