package rss

import (
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	f, err := os.Open("testdata/risky.biz.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	feed, err := Decode(f)
	if err != nil {
		t.Errorf("decode %v", err)
	}

	if feed.Version != "2.0" {
		t.Errorf("Version: expected %q, got %q", "2.0", feed.Version)
	}
	if feed.Channel.Title != "Risky Bulletin" {
		t.Errorf("Channel.Title: expected %q, got %q", "Risky Bulletin", feed.Channel.Title)
	}
	if feed.Channel.Link[0] != "https://risky.biz/" {
		t.Errorf("Channel.Link: expected %q, got %q", "https://risky.biz/", feed.Channel.Link)
	}
	// Fri, 11 Apr 2025 14:42:00 +1000
	if feed.Channel.PubDate.String() != "2025-04-11 14:42:00 +1000 AEST" {
		t.Errorf("Channel.PubDate: expected %q, got %q", "2025-04-11 14:42:00 +1000 AEST", feed.Channel.PubDate.String())
	}
	if feed.Channel.ITunesExplicit != false {
		t.Errorf("Channel.ITunesExplicit: expected %t, got %t", false, feed.Channel.ITunesExplicit)
	}

	item := feed.Channel.Items[0]
	if item.Title != "Risky Bulletin: Trump orders investigation into former CISA director Chris Krebs" {
		t.Errorf("Entry.Title: expected %q, got %q", "Risky Bulletin: Trump orders investigation into former CISA director Chris Krebs", item.Title)
	}
	if item.Link != "https://risky.biz/RBNEWS410/" {
		t.Errorf("Entry.Link: expected %q, got %q", "https://risky.biz/RBNEWS410/", item.Link)
	}
}
