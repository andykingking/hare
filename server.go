package hare

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"
)

type Server struct {
	db *DB

	net.TCPAddr
	listener  *net.TCPListener
	signal    chan bool
	waitGroup *sync.WaitGroup
}

type Command struct {
	Name string
	Args map[string]string
}

type Result struct {
	Err error
	Val interface{}
}

func (server *Server) Start() (err error) {
	server.waitGroup = &sync.WaitGroup{}
	server.signal = make(chan bool)
	server.listener, err = net.ListenTCP("tcp4", &server.TCPAddr)
	if err != nil {
		return err
	}

	go server.run()

	return nil
}

func (server *Server) Stop() {
	close(server.signal) // sends zero-valued signal
	server.waitGroup.Wait()
}

func (server *Server) run() {
	server.waitGroup.Add(1)
	defer server.waitGroup.Done()

	var err error
	for {
		select {
		case <-server.signal:
			log.Println("Closing listener")
			server.listener.Close()
			return
		default:
		}

		server.listener.SetDeadline(time.Now().Add(1e9))
		var conn *net.TCPConn
		conn, err = server.listener.AcceptTCP()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			}
			log.Println(err)
		}

		log.Println(conn.RemoteAddr(), "connected")
		server.waitGroup.Add(1)
		go server.handleConnection(conn)
	}
}

func (server *Server) handleConnection(conn *net.TCPConn) {
	defer conn.Close()
	defer server.waitGroup.Done()

	var err error
	var cmd = &Command{}
	var res = &Result{}
	var jsonRes []byte
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	jsonIn := json.NewDecoder(reader)

	err = jsonIn.Decode(&cmd)
	if err != nil {
		res.Err = err
	} else {
		res = server.HandleCommand(cmd)
	}

	jsonRes, err = json.Marshal(&res)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(jsonRes)

	writer.Write(jsonRes)
	writer.Flush()
}
