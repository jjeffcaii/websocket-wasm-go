package main

import (
  "time"

  "github.com/jjeffcaii/websocket-wasm-go"
)

func main() {
  ws, err := websocket.Connect("ws://127.0.0.1:8080/echo")
  if err != nil {
    panic(err)
  }
  _ = ws.Send([]byte("hello world!"))
  time.Sleep(3 * time.Second)

}
