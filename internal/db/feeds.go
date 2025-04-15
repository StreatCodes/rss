package db

import (
	"encoding/json"

	"github.com/streatCodes/rss/rss"
	bolt "go.etcd.io/bbolt"
)

var feedsBucket = []byte("feeds")

func (db *DB) SaveFeed(key []byte, feed *rss.Channel) error {
	value, err := json.Marshal(feed)
	if err != nil {
		return err
	}

	err = db.raw.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(feedsBucket)
		if err := bucket.Put(key, value); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (db *DB) GetFeed(key []byte) (*rss.Channel, error) {
	var channelBytes []byte
	err := db.raw.View(func(tx *bolt.Tx) error {
		if bucket := tx.Bucket(feedsBucket); bucket != nil {
			channelBytes = bucket.Get(key)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	var channel rss.Channel
	err = json.Unmarshal(channelBytes, &channel)
	return &channel, err
}
