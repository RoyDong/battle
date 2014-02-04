package admin

import (
    _"errors"
    "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.SetAction(func(r *pt.Request, p *pt.Response) error {
        /*
        user, ok := r.Session.Get("user").(*model.User)
        if !ok || !user.IsGrantd("RoleAdmin") {
            return errors.New("permission denied")
        }
        */

        key, _ := r.String("key")
        page, _ := r.Int64("page")
        if page <= 0 {
            page = 1
        }
        size, _ := r.Int64("size")
        if size <= 0 {
            size = 40
        }

        users := model.UserModel.
            Search(key, "created_at desc", size, (page - 1) * size)

        p.Render("admin/user/list", users)
        return nil
    }, "/admin/user" )
}
