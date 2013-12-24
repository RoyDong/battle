package controller

import (
    "github.com/roydong/potato"
    "github.com/roydong/potato/db"
)

type Main struct {
    potato.Controller
}

func (c *Main) Index() {
    stmt := new(db.Stmt)
    stmt.Select("*").
        From("user", "u").
        Where(stmt.And("u.id = :id")).
        Order("u.updated_at", "desc").
        Offset(1).
        Limit(2)

    c.RenderJson(stmt.String())

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
