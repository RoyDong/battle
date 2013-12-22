package controller

import (
    "github.com/roydong/potato"
)

type Main struct {
    potato.Controller
}

func (c *Main) Index() {
    c.RenderJson("111")
}

func (c *Main) Ws() {
    for {
        txt := c.WSReceive()

        if len(txt) == 0 {
            return
        }

        c.WSSendJson(map[string]interface{} {
            "message": "ok",
            "error": 0,
            "data": txt,
        })
    }
}
