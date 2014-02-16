package model

import (
    "github.com/roydong/potato/orm"
)

type Armory struct {
    BaseUnit
    base *Base
}

type armoryModel struct {
    *orm.Model
}
