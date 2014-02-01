package controller

import (
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.SetAction(func(r *pt.Request, p *pt.Response) {
        user, _ := r.Session.Get("user").(*model.User)
        p.Render("main", user)
    }, "")
}
