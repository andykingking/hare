package main

import (
  "fmt"
  "os"
  "os/signal"
  "github.com/captainpete/hare"
)

func main() {

  var hare = &hare.Hare{}
  var err error

  sigChan := make(chan os.Signal, 1)
  signal.Notify(sigChan, os.Interrupt)

  err = hare.Start()
  if err != nil {
    fmt.Println(err)
  }
  defer hare.Stop()

  <-sigChan
}
