package controller

import (
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.SetAction(func(r *pt.Request) *pt.Response {
        user, _ := r.Session.Get("user").(*model.User)
        if user == nil {
            return r.ErrorResponse(403, "no permisson")
        }

        x, _ := r.Int("x")
        y, _ := r.Int("y")
        name, _ := r.String("name")

        if len(user.Bases()) > 0 {
            return r.ErrorResponse(403, "you are not a new player")
        }
        loc := model.MapModel.Location(x, y)
        if loc == nil {
            return r.ErrorResponse(400, "coordinates not exists")
        }
        if loc.Geo != model.MapGeoLand {
            return r.ErrorResponse(400, "can't build base here")
        }
        if loc.Base() != nil {
            return r.ErrorResponse(400, "a base has already built here")
        }
        if !loc.Lock() {
            return r.ErrorResponse(400, "location was locked by others")
        }
        defer loc.Unlock()

        base := model.NewBase(name, user, loc)
        if model.BaseModel.Save(base) {
            user.AddBase(base)
            return r.JsonResponse(base)
        }
        return r.ErrorResponse(500, "can't save base to db")
    }, "/base/born")
}
