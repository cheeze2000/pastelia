package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/cheeze2000/pastelia/api/internal/snippet"
)

func TestPostgres(t *testing.T) {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		t.Fatal("Unable to connect to database")
	}

	defer dbpool.Close()

	want := &snippet.Snippet{
		Code: "fmt.Println(\"Hello, World!\")",
		Lang: "go",
	}

	slug, err := CreateSnippet(dbpool, want)
	got, err := ReadSnippet(dbpool, slug)

	if *got != *want {
		t.Errorf("Got %s, want %s", got.String(), want.String())
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
