package snippet

import (
	"encoding/json"
)

type Snippet struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}

func (snippet *Snippet) String() string {
	s, _ := json.Marshal(snippet)
	return string(s)
}
