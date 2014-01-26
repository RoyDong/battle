package controller

import (
    "github.com/roydong/potato"
    "github.com/roydong/battle/model"
)

type Map struct {
    potato.Controller
}


func (c *Map) Scan() {
    x, _ := c.Request.Int64("x")
    y, _ := c.Request.Int64("y")
    r, ok := c.Request.Int64("r")
    if !ok {
        r = 20
    }

    locs := model.MapModel.Rect(x, y, r)
    c.RenderJson(locs)
}

func (c *Map) Sum() {
    m := model.MapModel
    c.RenderJson([]int64{m.Metal(), m.Energy()})
}
