package controller

import (
    "github.com/roydong/potato"
    "github.com/roydong/battle/model"
)

func init() {
    potato.SetAction("", func(r *potato.Request, p *potato.Response) {
        user, _ := r.Session.Get("user").(*model.User)
        p.Render("main", user)
    })
}
