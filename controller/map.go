package controller

import (
    "github.com/roydong/potato"
    "github.com/roydong/battle/model"
)

func init() {
    potato.SetAction("map", map_main)

    potato.SetAction("map/rect", map_rect)

    potato.SetAction("map/sum", map_sum)

    potato.SetAction("map/refresh", map_refresh)
}

func map_main(r *potato.Request, p *potato.Response) {
    p.Render("map/main", nil)
}

func map_rect(r *potato.Request, p *potato.Response) {
    x, _ := r.Int64("x")
    y, _ := r.Int64("y")
    w, _ := r.Int64("w")
    h, _ := r.Int64("h")

    locs := model.MapModel.Rect(x, y, w, h)
    p.RenderJson(locs)
}

func map_sum(r *potato.Request, p *potato.Response) {
    m := model.MapModel
    p.RenderJson([]int64{m.Metal(), m.Energy(), int64(m.RefreshState)})
}

func map_refresh(r *potato.Request, p *potato.Response) {
    state, has := r.Int("state")
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
    p.RenderText("done")
}
