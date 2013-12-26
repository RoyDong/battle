package controller

import (
    "log"
    "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
)

type Main struct {
    potato.Controller
}

type A struct {
    a string `sql:"a"`
}
type B struct {
    b string `sql:"b"`
}

func (c *Main) Index() {
    var a A
    var b B

    r := new(orm.Rows)
    r.Scan(&a, &b)
    log.Println(a, b)
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
