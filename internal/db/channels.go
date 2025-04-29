package db

import (
	"encoding/json"
	"errors"

	"github.com/streatCodes/rss/rss"
	bolt "go.etcd.io/bbolt"
)

var channelsBucket = []byte("channels")

func (db *DB) SaveChannel(url string, channel *rss.Channel) error {
	value, err := json.Marshal(channel)
	if err != nil {
		return err
	}

	err = db.raw.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(channelsBucket)
		if err := bucket.Put([]byte(url), value); err != nil {
			return err
		}

		return nil
	})

	return err
}

var ErrNotExist = errors.New("channel does not exist")

func (db *DB) GetChannel(url string) (*rss.Channel, error) {
	var channelBytes []byte
	err := db.raw.View(func(tx *bolt.Tx) error {
		if bucket := tx.Bucket(channelsBucket); bucket != nil {
			channelBytes = bucket.Get([]byte(url))
		}
		if len(channelBytes) == 0 {
			return ErrNotExist
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
