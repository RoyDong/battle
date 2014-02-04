package controller

import (
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.SetAction(func(r *pt.Request, p *pt.Response) error {
        p.Render("map/main", nil)
        return nil
    }, "/map")

    pt.SetAction(func(r *pt.Request, p *pt.Response) error {
        x, _ := r.Int64("x")
        y, _ := r.Int64("y")
        w, _ := r.Int64("w")
        h, _ := r.Int64("h")

        locs := model.MapModel.Rect(x, y, w, h)
        p.RenderJson(locs)
        return nil
    }, "/map/rect")

    pt.SetAction(func(r *pt.Request, p *pt.Response) error {
        m := model.MapModel
        p.RenderJson([]int64{m.Metal(), m.Energy(), int64(m.RefreshState)})
        return nil
    }, "/map/sum")

    pt.SetAction(func(r *pt.Request, p *pt.Response) error {
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
        return nil
    }, "/map/refresh")
}
