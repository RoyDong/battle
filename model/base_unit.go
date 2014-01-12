package model

import (
    "time"
    "github.com/roydong/potato/orm"
)

const (
    UnitStateCreating  = 0
    UnitStateAvaliable = 1
    UnitStateUpgrading = 2
    UnitStateDestroyed = 3
)

type BaseUnit struct {
    Id int64 `column:"id"`
    Hp int64 `column:"hp"`
    Dp int64 `column:"dp"`
    Damage int64 `column:"damage"`
    Level int64 `column:"level"`
    State int `column:"state"`
    Lat float64 `column:"lat"`
    Lon float64 `column:"lon"`
    CreatedAt time.Time `column:"created_at"`
    UpdatedAt time.Time `column:"updated_at"`

    Location Location
}


