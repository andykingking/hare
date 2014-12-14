package hare

import (
	"bytes"
	"errors"
	"github.com/boltdb/bolt"
	"log"
)

type DB struct {
	path string

	Sequencer
	bolt *bolt.DB
}

func (db *DB) Open() error {
	var err error
	db.bolt, err = bolt.Open(db.path, 0600, nil)
	if err != nil {
		return err
	}

	db.Sequencer.Key.Id = 0
	db.Sequencer.Index = 65535
	db.Load(&db.Sequencer)
	log.Println("Sequencer", db.Sequencer)

	return nil
}

func (db *DB) Close() (err error) {
	err = db.Save(&db.Sequencer)
	return err
}

func (db *DB) NewKey() *Key {
	return &Key{Id: db.Next()}
}

func (db *DB) Load(obj Keyed) (err error) {
	log.Println("Loading", obj.IdBytes())
	err = db.bolt.Update(func(tx *bolt.Tx) error {
		bucket, _ := tx.CreateBucketIfNotExists([]byte("hare"))
		bObj := bucket.Get(obj.IdBytes())
		if bObj == nil {
			return errors.New("Object not found")
		} else {
			obj.Load(bytes.NewReader(bObj))
			obj.SetSaved(true)
		}
		return nil
	})
	return err
}

func (db *DB) Save(obj Keyed) (err error) {
	log.Println("Saving", obj.IdBytes())
	err = db.bolt.Update(func(tx *bolt.Tx) error {
		bucket, _ := tx.CreateBucketIfNotExists([]byte("hare"))
		var bObj bytes.Buffer
		obj.Save(&bObj)
		bucket.Put(obj.IdBytes(), bObj.Bytes())
		obj.SetSaved(true)
		return nil
	})
	return err
}

func (db *DB) Delete(obj Keyed) (err error) {
	log.Println("Deleting", obj.IdBytes())
	err = db.bolt.Update(func(tx *bolt.Tx) error {
		bucket, _ := tx.CreateBucketIfNotExists([]byte("hare"))
		return bucket.Delete(obj.IdBytes())
	})
	return err
}
