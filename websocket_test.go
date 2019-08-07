package websocket_test

import (
  "log"
  "testing"
  "time"

  "github.com/jjeffcaii/websocket-wasm-go"
)

func TestName(t *testing.T) {
  c, err := websocket.Connect("ws://127.0.0.1:8080/echo")
  if err != nil {
    t.Error(err)
  }
  log.Printf("%v\n", c)
  time.Sleep(5 * time.Second)
}
