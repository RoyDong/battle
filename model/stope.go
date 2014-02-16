package model

import (
    "github.com/roydong/potato/orm"
    "time"
)

type Stope struct {
    Id        int64     `column:"id"`
    BaseId    int64     `column:"base_id"`
    Hp        int64     `column:"hp"`
    Armor     int64     `column:"armor"`
    State     int       `column:"state"`
    CreatedAt time.Time `column:"created_at"`
    UpdatedAt time.Time `column:"updated_at"`

    base      *Base
    location  *Location
    rpm       int64
}

func NewStope(b *Base) *Stope {
    rpm, _ := Conf.Int64("stope.rpm")
    hp, _ := Conf.Int64("hp")
    armor, _ := Conf.Int64("armor")
    return &Stope{
        BaseId: b.Id,
        Hp: hp,
        Armor: armor,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),

        rpm: rpm,
        base: b,
    }
}

type stopeModel struct {
    *orm.Model
}

