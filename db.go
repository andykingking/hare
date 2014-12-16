package hare

import (
	"bytes"
	"errors"
	"github.com/boltdb/bolt"
	"log"
)

type DB struct {
	path string

	*Sequencer
	bolt *bolt.DB
}

func (db *DB) Open() (err error) {
	db.bolt, err = bolt.Open(db.path, 0600, nil)
	if err != nil {
		return err
	}

	err = db.loadSequencer()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Close() (err error) {
	err = db.saveSequencer()
	return err
}

func (db *DB) Load(obj Recorded) (err error) {
	bKey := obj.Key()
	err = db.bolt.Update(func(tx *bolt.Tx) error {
		bucketName := obj.BucketName()
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}
		bObj := bucket.Get(bKey)
		if bObj == nil {
			return errors.New("Object not found")
		} else {
			obj.Load(bytes.NewReader(bObj))
			// log.Println("LOAD", bucketName, bKey, bObj)
			obj.SetSaved(true)
		}
		return nil
	})
	return err
}

func (db *DB) Save(obj Recorded) (err error) {
	bKey := obj.Key()
	err = db.bolt.Update(func(tx *bolt.Tx) error {
		bucketName := obj.BucketName()
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}

		var bbObj bytes.Buffer
		obj.Save(&bbObj)
		bObj := bbObj.Bytes()

		err = bucket.Put(bKey, bObj)
		if err != nil {
			return err
		}
		// log.Println("SAVE", bucketName, bKey, bObj)
		obj.SetSaved(true)
		return nil
	})
	return err
}

func (db *DB) Delete(obj Recorded) (err error) {
	bKey := obj.Key()
	err = db.bolt.Update(func(tx *bolt.Tx) error {
		bucketName := obj.BucketName()
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}
		// log.Println("DELETE", bucketName, bKey)
		return bucket.Delete(bKey)
	})
	return err
}

func (db *DB) loadSequencer() (err error) {
	db.Sequencer = &Sequencer{}
	db.Load(db.Sequencer)
	if !db.Sequencer.Saved() {
		log.Println("No sequencer found, creating")
		err = db.Save(db.Sequencer)
		if err != nil {
			return err
		}
	} else {
		log.Println("Sequencer found")
	}
	log.Println("Sequencer", db.Sequencer)

	return nil
}

func (db *DB) saveSequencer() (err error) {
	err = db.Save(db.Sequencer)
	log.Println(db.Sequencer)

	return err
}

func (db *DB) NextKey() uint64 {
	// TODO write new sequence number to db
	return db.Sequencer.NextKey()
}
