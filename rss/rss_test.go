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
	if feed.Channel.PubDate.String() != "2025-04-11 14:42:00 +1000 AEST" {
		t.Errorf("Channel.PubDate: expected %q, got %q", "2025-04-11 14:42:00 +1000 AEST", feed.Channel.PubDate.String())
	}
	if feed.Channel.Image.Href != "https://risky.biz/static/img/rb-news.png" {
		t.Errorf("feed.Channel.Image: expected %q, got %q", "https://risky.biz/static/img/rb-news.png", feed.Channel.Image.Href)
	}
	if feed.Channel.Explicit != false {
		t.Errorf("Channel.Explicit: expected %t, got %t", false, feed.Channel.Explicit)
	}
	if feed.Channel.Categories[0].Text != "News" {
		t.Errorf("Channel.Explicit: expected %q, got %q", "News", feed.Channel.Categories[0].Text)
	}
	if feed.Channel.Categories[1].Text != "Technology" {
		t.Errorf("Channel.Explicit: expected %q, got %q", "Technology", feed.Channel.Categories[0].Text)
	}

	item := feed.Channel.Items[0]
	if item.Title != "Risky Bulletin: Trump orders investigation into former CISA director Chris Krebs" {
		t.Errorf("Entry.Title: expected %q, got %q", "Risky Bulletin: Trump orders investigation into former CISA director Chris Krebs", item.Title)
	}
	if item.Link != "https://risky.biz/RBNEWS410/" {
		t.Errorf("Entry.Link: expected %q, got %q", "https://risky.biz/RBNEWS410/", item.Link)
	}
}
