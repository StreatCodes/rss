package db

import (
	"os"

	bolt "go.etcd.io/bbolt"
)

type DB struct {
	raw *bolt.DB
}

func New(path string, mode os.FileMode, options *bolt.Options) (*DB, error) {
	db, err := bolt.Open(path, mode, options)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(feedsBucket)
		return nil
	})

	return &DB{raw: db}, err
}
