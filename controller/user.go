package controller

import (
    "time"
    "fmt"
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.SetAction(func(r *pt.Request, p *pt.Response) *pt.Error {
        if r.Method == "POST" {
            name, _ := r.String("name")
            email, _ := r.String("email")
            if len(email) == 0 {
                return pt.NewError(400, "email is empty")
            }

            passwd, _ := r.String("passwd")
            if len(passwd) == 0 {
                return pt.NewError(400, "password is empty")
            }

            user := &model.User{
                Name:      name,
                Email:     email,
                UpdatedAt: time.Now(),
                CreatedAt: time.Now(),
            }
            user.SetPasswd(passwd)
            if model.UserModel.Save(user) {
                r.Session.Set("user", user)
                p.Redirect(r, "/user", 302)
                return nil
            }

            return pt.NewError(500, "server bussy try later")
        }

        p.Render("user/signup", nil)
        return nil
    }, "/signup")

    pt.SetAction(func(r *pt.Request, p *pt.Response) *pt.Error {
        if r.Method == "POST" {
            email, _ := r.String("email")
            if len(email) == 0 {
                return pt.NewError(400, "email is empty")
            }

            passwd, _ := r.String("passwd")
            if len(passwd) == 0 {
                return pt.NewError(400, "password is empty")
            }

            m := model.UserModel
            if user := m.UserByEmail(email); user != nil &&
                user.CheckPasswd(passwd) {
                r.Session.Set("user", user)
                p.Redirect(r, "/user", 302)
                return nil
            }

            return pt.NewError(400, "email or password error")
        }

        p.Render("user/signin", nil)
        return nil
    }, "/signin")

    pt.SetAction(func(r *pt.Request, p *pt.Response) *pt.Error {
        r.Session.Set("user", nil)
        p.RenderText("user/signout")
        return nil
    }, "/signout")

    pt.SetAction(func(r *pt.Request, p *pt.Response) *pt.Error {
        id, _ := r.Int64("$1")
        var user *model.User
        if id == 0 {
            user, _ = r.Session.Get("user").(*model.User)
        } else {
            user = model.UserModel.User(id)
        }

        if id == 0 && user != nil {
            p.Redirect(r, fmt.Sprintf("/user/%d", user.Id), 302)
            return nil
        }

        p.Render("user/show", user)
        return nil
    }, `/user/(\d+)`, "/user", "/")
}
