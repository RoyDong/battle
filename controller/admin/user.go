package admin

import (
    "fmt"
    "time"
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func IsGrantd(session *pt.Session, roles ...string) bool {
    user, ok := session.Get("user").(*model.User)
    return ok && user.IsGrantd(roles...)
}

func init() {
    pt.SetAction(func(r *pt.Request, p *pt.Response) *pt.Error {
        if !IsGrantd(r.Session, "RoleAdmin") {
            return pt.NewError(403, "permission denied")
        }

        key, _ := r.String("key")

        page, _ := r.Int64("page")
        if page <= 0 {
            page = 1
        }

        size, _ := r.Int64("size")
        if size <= 0 {
            size = 40
        }

        sort, _ := r.String("sort")
        if sort == "" {
            sort = "id"
        }

        order, _ := r.String("order")
        if order == "" {
            order = "desc"
        }

        m := map[string]string{sort: order}
        users := model.
            UserModel.
            Search(key, m, size, (page - 1) * size)

        p.Render("admin/user/list", users)
        return nil
    }, "/admin/user")

    pt.SetAction(func(r *pt.Request, p *pt.Response) *pt.Error {
        id, _ := r.Int64("$1")
        user := model.UserModel.User(id)
        if user == nil {
            return pt.NewError(404, "user not found")
        }

        p.Render("admin/user/detail", user)
        return nil
    }, `/admin/user/(\d+)`)

    pt.SetAction(func(r *pt.Request, p *pt.Response) *pt.Error {
        if r.Method != "post" {
            return pt.NewError(400, "must be post")
        }
        id, _ := r.Int64("$1")
        user := model.UserModel.User(id)
        if user == nil {
            return pt.NewError(404, "user not found")
        }

        desc, _ := r.String("desc")
        name, _ := r.String("role")
        if name == "" {
            return pt.NewError(400, "need a role name")
        }

        if user.IsGrantd(name) {
            p.Redirect(r, fmt.Sprintf("/admin/user/%d", id), 302)
            return nil
        }

        role := model.RoleModel.RoleByName(name)
        if role == nil {
            role = &model.Role{
                Name: name,
                Desc: desc,
                CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
            }
            if !model.RoleModel.Save(role) {
                return pt.NewError(500, "can't save role " + name)
            }
        }

        if user.AddRole(role) {
            p.Redirect(r, fmt.Sprintf("/admin/user/%d", id), 302)
            return nil
        }
        return pt.NewError(500, "can't add role " + name + " to user")
    }, `/admin/user/(\d+)/add_role`)
}
