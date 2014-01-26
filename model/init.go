package model

import (
    "github.com/roydong/potato/orm"
)

var MapModel *mapModel
var UserModel *userModel

func Init() {
    MapModel = newMapModel()

    UserModel = &userModel{orm.NewModel("user", new(User))}
}
