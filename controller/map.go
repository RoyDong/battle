package controller

import (
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.SetAction(func(r *pt.Request, p *pt.Response) {
        p.Render("map/main", nil)
    }, "/map")

    pt.SetAction(func(r *pt.Request, p *pt.Response) {
        x, _ := r.Int64("x")
        y, _ := r.Int64("y")
        w, _ := r.Int64("w")
        h, _ := r.Int64("h")

        locs := model.MapModel.Rect(x, y, w, h)
        p.RenderJson(locs)
    }, "/map/rect")

    pt.SetAction(func(r *pt.Request, p *pt.Response) {
        m := model.MapModel
        p.RenderJson([]int64{m.Metal(), m.Energy(), int64(m.RefreshState)})
    }, "/map/sum")

    pt.SetAction(func(r *pt.Request, p *pt.Response) {
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
