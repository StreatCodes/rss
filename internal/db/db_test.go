package db

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/streatCodes/rss/rss"
	bolt "go.etcd.io/bbolt"
)

func TestSaveFeed(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test.db")

	f, err := os.Open("example.rss")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	feed, err := rss.Decode(f)
	if err != nil {
		t.Error(err)
	}

	db, err := New(filePath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		t.Error(err)
	}

	feedUrl := "http://example.com"
	err = db.SaveChannel(feedUrl, &feed.Channel)
	if err != nil {
		t.Error(err)
	}

	retrievedFeed, err := db.GetChannel(feedUrl)
	if err != nil {
		t.Error(err)
	}

	if feed.Channel.Title != retrievedFeed.Title {
		t.Errorf("Retrieved feed title: expected %q, got %q", feed.Channel.Title, retrievedFeed.Title)
	}
}
