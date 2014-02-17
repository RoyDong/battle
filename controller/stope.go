package controller

import (
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.SetAction(func(r *pt.Request) *pt.Response {
        user, _ := r.Session.Get("user").(*model.User)
        if user == nil {
            return pt.ErrorResponse(403, "not signed in")
        }

        bid, _ := r.Int64("bid")
        if bid == 0 {
            return pt.ErrorResponse(400, "missing base id")
        }

        base := user.Bases()[bid]
        if base == nil {
            return pt.ErrorResponse(400, "error base id")
        }

        stope := model.NewStope(base)
        model.StopeModel.Save(stope)
        return pt.JsonResponse(stope)
    }, "/stope/new")
}
