package model

import (
    pt "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
)

var MapModel *mapModel

var UserModel *userModel

var RoleModel *roleModel

var UserRoleModel *userRoleModel

func init() {
    pt.AddHandler("after_orm_init", func(args ...interface{}) {
        MapModel = newMapModel()

        UserModel = &userModel{orm.NewModel("user", &User{})}

        RoleModel = &roleModel{orm.NewModel("role", &Role{})}

        UserRoleModel = &userRoleModel{orm.NewModel("user_role", &UserRole{})}
    })
}
