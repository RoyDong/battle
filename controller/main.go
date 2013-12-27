package controller

import (
    "log"
    "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
)

type Main struct {
    potato.Controller
}

func (c *Main) Index() {
    stmt := orm.NewStmt()
    stmt.Select("u.id, u.name, s.id").From("User", "u").LeftJoin("User", "s", "s.id = :sid").Where("u.id = :id")
    log.Println(1)
    c.Response.Write([]byte(stmt.String()))
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
