package controller

import (
    "log"
    "github.com/roydong/potato"
    ws "code.google.com/p/go.net/websocket"
)

type Main struct {
    potato.Controller
}

func (c *Main) Index() {
    c.RenderJson("111")
}

func (c *Main) Ws() {
    conn := c.Request.WSConn
    var a string

    for {
        if err := ws.Message.Receive(conn, &a); err != nil {
            log.Println(err)
        }

        log.Println(a)
        ws.Message.Send(conn, a)
    }
}
