package hare

import (
  "net"
  "fmt"
  "bufio"
  "encoding/json"
)

type Server struct {
  addr        string
  listener    net.Listener
  running     bool
  db          DB
}

type Command struct {
  Name        string
  Args        map[string]string
}

type Result struct {
  Err         string
  Val         interface{}
}

func (server *Server) Start() error {

  listener, err := net.Listen("tcp", server.addr)
  if err != nil {
    return err
  }

  server.listener = listener
  go server.run()

  return nil
}

func (server *Server) Stop() error {
  server.running = false
  err := server.listener.Close()
  return err
}

func (server *Server) run() {
  server.running = true

  var err error
  for server.running {
    var conn net.Conn
    conn, err = server.listener.Accept()
    if err != nil {
      fmt.Println(err) // TODO handle error
      continue
    }
    go server.handleConnection(conn)
  }
}

func (server *Server) handleConnection(conn net.Conn) {
  defer conn.Close()

  var err error
  var cmd = &Command{}
  var res = &Result{}
  var jsonRes []byte
  reader := bufio.NewReader(conn)
  writer := bufio.NewWriter(conn)
  jsonIn := json.NewDecoder(reader)

  err = jsonIn.Decode(&cmd)
  if err != nil {
    res.Err = err.Error()
  } else {
    res = server.HandleCommand(cmd)
  }

  jsonRes, err = json.Marshal(&res)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(string(jsonRes))

  writer.Write(jsonRes)
  writer.Flush()
}
