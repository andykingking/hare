package hare

import (
  "encoding/binary"
  "github.com/boltdb/bolt"
)

type DB struct {
  path    string
  opened  bool

  bolt    *bolt.DB
  seq     *Sequencer
}

func (db *DB) Open() error {
  var err error
  db.bolt, err = bolt.Open(db.path, 0600, nil)
  if err != nil {
    return err
  }
  db.opened = true

  db.seq = &Sequencer{index: db.loadSequenceIndex()}

  return nil
}

func (db *DB) Close() error {
  db.opened = false
  db.saveSequenceIndex(db.seq.index)
  return db.bolt.Close()
}

func (db *DB) loadSequenceIndex() uint64 {
  var seqIndex uint64

  db.bolt.Update(func(tx *bolt.Tx) error {
    bucket, _ := tx.CreateBucketIfNotExists([]byte("seq"))
    bSeqIndex := bucket.Get([]byte("seqindex"))
    if bSeqIndex == nil {
      seqIndex = 0
    } else {
      seqIndex  = binary.LittleEndian.Uint64(bSeqIndex[0:8])
    }

    return nil
  })

  return seqIndex
}

func (db *DB) saveSequenceIndex(seqIndex uint64) {
  db.bolt.Update(func(tx *bolt.Tx) error {
    bucket, _ := tx.CreateBucketIfNotExists([]byte("seq"))
    bSeqIndex := make([]byte, 8); binary.LittleEndian.PutUint64(bSeqIndex, seqIndex)
    bucket.Put([]byte("seqindex"), bSeqIndex)

    return nil
  })
}
