package controller

import (
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.SetAction(func(r *pt.Request, p *pt.Response) *pt.Error {
        p.Render("map/main", nil)
        return nil
    }, "/map")

    pt.SetAction(func(r *pt.Request, p *pt.Response) *pt.Error {
        x, _ := r.Int("x")
        y, _ := r.Int("y")
        w, _ := r.Int("w")
        h, _ := r.Int("h")

        locs := model.MapModel.Rect(x, y, w, h)
        p.RenderJson(locs)
        return nil
    }, "/map/rect")

    pt.SetAction(func(r *pt.Request, p *pt.Response) *pt.Error {
        m := model.MapModel
        p.RenderJson([]int64{m.Metal(), m.Energy(), int64(m.RefreshState)})
        return nil
    }, "/map/sum")

    pt.SetAction(func(r *pt.Request, p *pt.Response) *pt.Error {
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
