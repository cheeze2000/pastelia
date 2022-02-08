package slug

import (
	"testing"
)

func TestSlugLength(t *testing.T) {
	slug := RandomSlug()
	if len(slug) != 10 {
		t.Errorf("Got %d, want 10", len(slug))
	}
}

func TestSlugRandomness(t *testing.T) {
	slug0 := RandomSlug()
	slug1 := RandomSlug()

	if slug0 == slug1 {
		t.Errorf("Slugs are not random")
	}
}
