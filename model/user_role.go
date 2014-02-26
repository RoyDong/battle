package model

import (
    "github.com/roydong/potato/orm"
    "time"
    "log"
)

type UserRole struct {
    UserId    int64     `column:"user_id"`
    RoleId    int64     `column:"role_id"`
    CreatedAt time.Time `column:"created_at"`
}

type userRoleModel struct {
    *orm.Model
}

func (m *userRoleModel) Save(u *User, r *Role) bool {
    _, e := orm.NewStmt("").
        Insert("UserRole", "user_id", "role_id", "created_at").
        Exec(u.Id, r.Id, time.Now())

    if e != nil {
        log.Println(e)
        return false
    }
    return true
}

func (m *userRoleModel) Remove(u *User, r *Role) bool {
    _, e := orm.NewStmt("").
        Delete("UserRole").
        Where("user_id = ? AND role_id = ?").
        Exec(u.Id, r.Id)

    if e != nil {
        log.Println(e)
        return false
    }

    return true
}
