package model

import (
    "github.com/roydong/potato/orm"
    "time"
)

type Role struct {
    Id        int64     `column:"id"`
    Name      string    `column:"name"`
    Desc      string    `column:"desc"`
    CreatedAt time.Time `column:"created_at"`
    UpdatedAt time.Time `column:"updated_at"`
}

type roleModel struct {
    *orm.Model
}
