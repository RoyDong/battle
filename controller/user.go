package controller

import (
    "time"
    "github.com/roydong/potato"
    "github.com/roydong/battle/model"
)

func init() {
    potato.SetAction("signin", signin)

    potato.SetAction("signup", signup)

    potato.SetAction("signout", signout)
}

func signup(r *potato.Request, p *potato.Response) {
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
            p.Redirect(r, "/", 302)
        }

        panic("server bussy try later")
    }

    p.Render("user/signup", nil)
}

func signin(r *potato.Request, p *potato.Response) {
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
            p.Redirect(r, "/admin/setting", 302)
        }

        panic("email or password error")
    }

    p.Render("admin/user/signin", nil)
}

func signout(r *potato.Request, p *potato.Response) {
    p.RenderText("signout")
}
