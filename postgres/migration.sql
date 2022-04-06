CREATE TABLE IF NOT EXISTS snippets (
	id   serial PRIMARY KEY,
	slug char(10),
	code text,
	lang text
);
