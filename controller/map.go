package controller

import (
    "github.com/roydong/potato"
    "github.com/roydong/battle/model"
)

type Map struct {
    potato.Controller
}


func (c *Map) Main() {

    c.Render("map/main", nil)
}

func (c *Map) Scan() {
    x, _ := c.Request.Int64("x")
    y, _ := c.Request.Int64("y")
    r, _ := c.Request.Int64("r")

    locs := model.MapModel.Rect(x - r, y - r, 2 * r, 2 * r)
    c.RenderJson(locs)
}

func (c *Map) Rect() {
    r := c.Request
    x, _ := r.Int64("x")
    y, _ := r.Int64("y")
    w, _ := r.Int64("w")
    h, _ := r.Int64("h")

    locs := model.MapModel.Rect(x, y, w, h)
    c.RenderJson(locs)
}

func (c *Map) Sum() {
    m := model.MapModel
    c.RenderJson([]int64{m.Metal(), m.Energy(), int64(m.RefreshState)})
}

func (c *Map) Refresh() {
    state, has := c.Request.Int("state")
    m := model.MapModel

    if !has {
        state = model.RefreshStateOn
        if m.RefreshState == model.RefreshStateOn {
            state = model.RefreshStateOff
        } else {
            state = model.RefreshStateOn
        }
    }

    m.RefreshState = state
    c.RenderText("done")
}
