package app


import (
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.AddHandler("before_serve", func(args ...interface{}) {
        if pt.Env == "dev" {
            pt.AddHandler("before_action", func(args ...interface{}) {
                r := args[0].(*pt.Request)
                user, _ := r.Session.Get("user").(*model.User)
                if user == nil {
                    user = model.UserModel.UserByEmail("i@roydong.com")
                    r.Session.Set("user", user)
                }
            })
        }
    })
}
