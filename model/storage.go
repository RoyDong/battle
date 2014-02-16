package model

import (
    "github.com/roydong/potato/orm"
)

type Storage struct {
    BaseUnit
    base *Base
}

type storageModel struct {
    *orm.Model
}
