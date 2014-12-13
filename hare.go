package hare

import (
  "fmt"
)

type Hare struct {
  db        DB
  server    Server
}

func (hare *Hare) Start() error {

  var err error

  path := "hare.db"
  fmt.Printf("Opening database %s\n", path)
  hare.db = DB{path: path}
  err = hare.db.Open()
  if err != nil {
    fmt.Println(err)
    return err
  }
  fmt.Printf("Sequence %v\n", hare.db.seq.index)

  addr := "127.0.0.1:4888"
  fmt.Printf("Starting server on %s\n", addr)
  hare.server = Server{addr: addr, db: hare.db}
  err = hare.server.Start()
  if err != nil {
    fmt.Println(err)
    return err
  }

  return nil
}

func (hare *Hare) Stop() {

  var err error

  fmt.Println("Stopping server")
  err = hare.server.Stop()
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println("Closing database")
  fmt.Printf("Sequence %v\n", hare.db.seq.index)
  err = hare.db.Close()
  if err != nil {
    fmt.Println(err)
  }

}
