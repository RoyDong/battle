package controller

import (
    "fmt"
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
    "time"
)

func init() {
    pt.SetAction(func(r *pt.Request, p *pt.Response) {
        if r.Method == "POST" {
            name, _ := r.String("name")
            email, _ := r.String("email")
            if len(email) == 0 {
                panic("email is empty")
            }

            passwd, _ := r.String("passwd")
            if len(passwd) == 0 {
                panic("password is empty")
            }

            user := new(model.User)
            user.Name = name
            user.Email = email
            user.SetPasswd(passwd)
            user.UpdatedAt = time.Now()
            user.CreatedAt = time.Now()

            if model.UserModel.Save(user) {
                r.Session.Set("user", user, true)
                p.Redirect(r, "/user", 302)
            }

            panic("server bussy try later")
        }

        p.Render("user/signup", nil)
    }, "/signup")

    pt.SetAction(func(r *pt.Request, p *pt.Response) {
        if r.Method == "POST" {
            email, _ := r.String("email")
            if len(email) == 0 {
                panic("email is empty")
            }

            passwd, _ := r.String("passwd")
            if len(passwd) == 0 {
                panic("password is empty")
            }

            m := model.UserModel
            if user := m.UserByEmail(email); user != nil &&
                user.CheckPasswd(passwd) {
                r.Session.Set("user", user, true)
                p.Redirect(r, "/user", 302)
            }

            panic("email or password error")
        }

        p.Render("user/signin", nil)
    }, "/signin")

    pt.SetAction(func(r *pt.Request, p *pt.Response) {
        r.Session.Set("user", nil, true)
        p.RenderText("user/signout")
    }, "/signout")

    pt.SetAction(func(r *pt.Request, p *pt.Response) {
        id, _ := r.Int64("$1")
        var user *model.User
        if id == 0 {
            user, _ = r.Session.Get("user").(*model.User)
        } else {
            user = model.UserModel.User(id)
        }

        if user == nil {
            panic("user not found")
        }

        if id == 0 {
            p.Redirect(r, fmt.Sprintf("/user/%d", user.Id), 301)
        }

        p.Render("user/show", user)
    }, `/user/(\d+)`, "/user")
}
