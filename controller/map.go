package controller

import (
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.SetAction(func(r *pt.Request) *pt.Response {
        return r.HtmlResponse("map/main", nil)
    }, "/map")

    pt.SetAction(func(r *pt.Request) *pt.Response {
        x, _ := r.Int("x")
        y, _ := r.Int("y")
        w, _ := r.Int("w")
        h, _ := r.Int("h")

        locs := model.MapModel.Rect(x, y, w, h)
        return r.JsonResponse(locs)
    }, "/map/rect")

    pt.SetAction(func(r *pt.Request) *pt.Response {
        m := model.MapModel
        return r.JsonResponse([]int64{m.Metal(), m.Energy(), int64(m.RefreshState)})
    }, "/map/sum")

    pt.SetAction(func(r *pt.Request) *pt.Response {
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
        return r.TextResponse("done")
    }, "/map/refresh")
}
