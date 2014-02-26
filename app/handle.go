package app


import (
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.AddHandler("run", func(args ...interface{}) {
        if pt.Env == "dev" {
            pt.AddHandler("action", func(args ...interface{}) {
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
