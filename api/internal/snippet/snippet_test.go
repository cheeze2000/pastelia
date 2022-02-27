package snippet

import (
	"testing"
)

func TestStringify(t *testing.T) {
	snippet := &Snippet{Code: "x := 3", Lang: "go"}
	want := "{\"code\":\"x := 3\",\"lang\":\"go\"}"
	got := snippet.String()

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}
