package model

import (
    "github.com/roydong/potato/orm"
)

type Supply struct {
    BaseUnit
    base *Base
}

type supplyModel struct {
    *orm.Model
}
