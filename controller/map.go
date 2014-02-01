package controller

import (
    "github.com/roydong/potato"
    "github.com/roydong/battle/model"
)

func init() {
    potato.SetAction(func(r *potato.Request, p *potato.Response) {
        p.Render("map/main", nil)
    }, "/map")

    potato.SetAction(func(r *potato.Request, p *potato.Response) {
        x, _ := r.Int64("x")
        y, _ := r.Int64("y")
        w, _ := r.Int64("w")
        h, _ := r.Int64("h")

        locs := model.MapModel.Rect(x, y, w, h)
        p.RenderJson(locs)
    }, "/map/rect")

    potato.SetAction(func(r *potato.Request, p *potato.Response) {
        m := model.MapModel
        p.RenderJson([]int64{m.Metal(), m.Energy(), int64(m.RefreshState)})
    }, "/map/sum")

    potato.SetAction(func(r *potato.Request, p *potato.Response) {
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
    }, "/map/refresh")
}
