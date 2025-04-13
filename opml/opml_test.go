package opml

import (
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	f, err := os.Open("otl.opml")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	// The tricky thing is decoding child elements recursively
	// just using struct tags. So checking the number of elements
	// is expected gives us some confidence we're reading
	// the document correctly.
	var counts = map[string]int{
		"aggregators": 6,
		"apple":       6,
		"corp":        4,
		"email":       4,
		"Friends":     5,
		"lang":        4,
		"misc":        92,
		"repos":       6,
	}
	for _, node := range doc.Body {
		want, ok := counts[node.Title]
		if !ok {
			t.Errorf("unknown node title %s", node.Title)
			continue
		}
		if len(node.Children) != want {
			t.Errorf("%d child elements in %s, want %d", len(node.Children), node.Title, want)
		}
	}
}
