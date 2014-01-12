package model

import (
    "github.com/roydong/potato/orm"
)

const (
    FortStateDestroyed    = 0
    FortStateConstructing = 1
    FortStateUpgrading    = 1
    FortStateAvaliable    = 2

    GeoLand = 0
    GeoSea  = 1
)

type Location struct {
    Lat float64 `column:"lat"`
    Lon float64 `column:"lon"`
    Geo int `column:"geo"`
    FortId int64 `column:"fort_id"`
}

type BaseUnit struct {
    Id int64 `column:"id"`
    Hp int64 `column:"hp"`
    Dp int64 `column:"dp"`
    Damage int64 `column:"damage"`
    Level int64 `column:"level"`
    State int `column:"state"`
    Lat float64 `column:"lat"`
    Lon float64 `column:"lon"`
    CreatedAt int64 `column:"created_at"`
    UpdatedAt int64 `column:"updated_at"`

    Location Location
}


type Fort struct {
    BaseUnit
}

type Army struct {
    BaseUnit
}



type Base struct {
    user_id int64 `column:"user_id"`
}


