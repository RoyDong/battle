package model

import (
    "github.com/roydong/potato/orm"
)

const (
    UnitStateDestroyed = 0
    UnitStateConstructing = 1
    UnitStateAvaliable = 2
)

type Unit struct {
    Id int64 `column:"id"`
    Level int64 `column:"level"`
    State int `column:"state"`
    CreatedAt int64 `column:"created_at"`
    UpdatedAt int64 `column:"updated_at"`
}


