package model

import (
    "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
)

var MapModel *mapModel
var UserModel *userModel
var RoleModel *roleModel
var UserRoleModel *userRoleModel

func init() {
    potato.E.AddEventHandler("orm_init_done", func(args ...interface{}) {
        MapModel = newMapModel()

        UserModel = &userModel{orm.NewModel("user", &User{})}

        RoleModel = &roleModel{orm.NewModel("role", &Role{})}

        UserRoleModel = &userRoleModel{orm.NewModel("user_role", &UserRole{})}
    })
}
