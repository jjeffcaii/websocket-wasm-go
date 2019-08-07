// +build js,wasm

package websocket

import (
  "fmt"
  "syscall/js"
)

type Socket struct {
  ws js.Value
}

func (p *Socket) Send(b []byte) (err error) {
  ret := p.ws.Call("send", js.TypedArrayOf(b))
  fmt.Printf("send result: %v\n", ret)
  return
}

func Connect(url string) (s *Socket, err error) {

  // https://www.runoob.com/html/html5-websocket.html
  // https://developer.mozilla.org/en-US/docs/Web/API/WebSocket

  //"ws://localhost:8080/ws"
  ws := js.Global().Get("WebSocket").New(url)

  success := make(chan struct{})
  fail := make(chan error)

  ws.Call("addEventListener", "open", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    success <- struct{}{}
    return nil
  }))

  ws.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    for _, value := range args {
      fmt.Printf("message: %s\n", value.JSValue().Get("data").Type().String())
    }
    return nil
  }))

  select {
  case <-success:
    s = &Socket{ws: ws}
  case err = <-fail:
  }
  return
}
