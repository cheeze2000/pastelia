package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/cheeze2000/pastelia/backend/internal/slug"
	"github.com/cheeze2000/pastelia/backend/internal/snippet"
)

func CreateSnippet(dbpool *pgxpool.Pool, snippet *snippet.Snippet) (string, error) {
	slug := slug.RandomSlug()
	_, err := dbpool.Exec(
		context.Background(),
		"insert into snippets (slug, code, lang) values ($1, $2, $3)",
		slug,
		snippet.Code,
		snippet.Lang,
	)

	if err != nil {
		return "", err
	}

	return slug, nil
}

func ReadSnippet(dbpool *pgxpool.Pool, slug string) (*snippet.Snippet, error) {
	var code string
	var lang string

	err := dbpool.
		QueryRow(
			context.Background(),
			"select code, lang from snippets where slug=$1",
			slug,
		).
		Scan(&code, &lang)

	snippet := &snippet.Snippet{Code: code, Lang: lang}
	if err != nil {
		return snippet, err
	}

	return snippet, nil
}
