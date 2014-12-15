package hare

import (
	"log"
	"net"
)

type Hare struct {
	db     DB
	server Server
}

func (hare *Hare) Start() (err error) {
	log.Println("Starting Hare")

	path := "hare.db"
	log.Println("Opening database", path)
	hare.db = DB{path: path}
	err = hare.db.Open()
	if err != nil {
		log.Println(err)
		return err
	}

	var laddr *net.TCPAddr
	laddr, err = net.ResolveTCPAddr("tcp", "127.0.0.1:4888")
	if err != nil {
		return err
	}
	log.Println("Starting server on", laddr)
	hare.server = Server{db: &hare.db, TCPAddr: *laddr}
	err = hare.server.Start()
	if err != nil {
		return err
	}

	return nil
}

func (hare *Hare) Stop() (err error) {
	log.Println("Stopping Hare")

	log.Println("Stopping server")
	hare.server.Stop()
	log.Println("Server stopped")

	log.Println("Closing database at", hare.db.Mark)
	err = hare.db.Close()
	if err != nil {
		return err
	}
	log.Println("Database closed")

	return nil
}
