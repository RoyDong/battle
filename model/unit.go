package model

import (
    "github.com/roydong/potato"
)

const (
    UnitStateDestroyed = 0
    UnitStateConstructing = 1
    UnitStateAvaliable = 2
)

type Unit struct {
    id int64
    level int64
    State int
    CreatedAt, UpdatedAt int64
}


