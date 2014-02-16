package model

import (
    "github.com/roydong/potato/orm"
)

type Lab struct {
    BaseUnit
    base *Base
}

type labModel struct {
    *orm.Model
}
