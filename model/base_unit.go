package model

import (
    "time"
)

const (
    UnitStateCreating  = 0
    UnitStateAvaliable = 1
    UnitStateUpgrading = 2
    UnitStateDestroyed = 3

    AttackModeNormal      = 0
    AttackModeExplosive   = 1
    AttackModePenetrative = 2
)

type BaseUnit struct {
    Id        int64     `column:"id"`
    Hp        int64     `column:"hp"`
    Armor     int64     `column:"armor"`
    Level     int64     `column:"level"`
    State     int       `column:"state"`
    X         int64     `column:"x"`
    Y         int64     `column:"y"`
    CreatedAt time.Time `column:"created_at"`
    UpdatedAt time.Time `column:"updated_at"`

    Location *Location
    Weapon   *Weapon
}

type Weapon struct {
    AttackMode int
    Interval   int
    Damage     int64
}
