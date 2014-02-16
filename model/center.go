package model

import (
    "github.com/roydong/potato/orm"
    "fmt"
)

type Center struct {
    BaseUnit
    base *Base
}


func NewCenter(level int64) *Center {
    center := &Center{}
    conf := Conf.Tree(fmt.Sprintf("center.%d", level))
    if conf != nil {
        center.Hp, _ = conf.Int64("hp")
        center.Armor, _ = conf.Int64("armor")
    }
    return center
}

type centerModel struct {
    *orm.Model
}
