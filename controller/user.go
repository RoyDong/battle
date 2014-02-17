package controller

import (
    "time"
    "fmt"
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.SetAction(func(r *pt.Request) *pt.Response {
        if r.Method == "POST" {
            name, _ := r.String("name")
            email, _ := r.String("email")
            if len(email) == 0 {
                return pt.ErrorResponse(400, "email is empty")
            }

            passwd, _ := r.String("passwd")
            if len(passwd) == 0 {
                return pt.ErrorResponse(400, "password is empty")
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
                return pt.RedirectResponse("/user", 302)
            }
            return pt.ErrorResponse(500, "server bussy try later")
        }
        return pt.HtmlResponse("user/signup", nil)
    }, "/signup")

    pt.SetAction(func(r *pt.Request) *pt.Response {
        if r.Method == "POST" {
            email, _ := r.String("email")
            if len(email) == 0 {
                return pt.ErrorResponse(400, "email is empty")
            }

            passwd, _ := r.String("passwd")
            if len(passwd) == 0 {
                return pt.ErrorResponse(400, "password is empty")
            }

            m := model.UserModel
            if user := m.UserByEmail(email); user != nil &&
                user.CheckPasswd(passwd) {
                r.Session.Set("user", user)
                return pt.RedirectResponse("/user", 302)
            }

            return pt.ErrorResponse(400, "email or password error")
        }

        return pt.HtmlResponse("user/signin", nil)
    }, "/signin")

    pt.SetAction(func(r *pt.Request) *pt.Response {
        r.Session.Set("user", nil)
        return pt.TextResponse("done")
    }, "/signout")

    pt.SetAction(func(r *pt.Request) *pt.Response {
        id, _ := r.Int64("$1")
        var user *model.User
        if id == 0 {
            user, _ = r.Session.Get("user").(*model.User)
        } else {
            user = model.UserModel.User(id)
        }

        if id == 0 && user != nil {
            return pt.RedirectResponse(fmt.Sprintf("/user/%d", user.Id), 302)
        }

        return pt.HtmlResponse("user/show", user)
    }, `/user/(\d+)`, "/user", "/")
}
