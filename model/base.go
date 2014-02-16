package model

import (
    "github.com/roydong/potato/orm"
    "fmt"
    "time"
)

type Base struct {
    Id        int64     `column:"id"`
    UserId    int64     `column:"user_id"`
    Star      int       `column:"star"`
    X         int       `column:"x"`
    Y         int       `column:"y"`
    Name      string    `column:"name"`
    CreatedAt time.Time `column:"created_at"`
    UpdatedAt time.Time `column:"updated_at"`

    SelfSupply  int64
    SelfStorage int64

    user     *User
    location *Location
    center   *Center
    lab      *Lab
    armory   *Armory
    supply   *Supply
    stope    *Stope
    storage  *Storage
}


func NewBase(n string, u *User, loc *Location) *Base {
    b := &Base{
        UserId: u.Id,
        Name: n,
        X: loc.X,
        Y: loc.Y,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        Star: 1,

        location: loc,
        user: u,
    }
    b.SelfSupply, _ = Conf.Int64("base.supply", )
    b.SelfStorage, _ = Conf.Int64("base.storage", )
    return b
}

type baseModel struct {
    *orm.Model
}


func baseQueryStmt() *orm.Stmt {
    return orm.NewStmt().
        Select("b.*, l.*, c.*, a.*, r.*, s.*, t.*, o.*").
        From("Base", "b").
        InnerJoin("Location", "l", "l.x = b.x AND l.y = b.y").
        LeftJoin("Center",    "c", "c.base_id = b.id").
        LeftJoin("Lab",       "a", "a.base_id = b.id").
        LeftJoin("Armory",    "r", "r.base_id = b.id").
        LeftJoin("Supply",    "s", "s.base_id = b.id").
        LeftJoin("Stope",     "t", "t.base_id = b.id").
        LeftJoin("Storage",   "o", "o.base_id = o.id")
}

func scanBaseEntity(rows *orm.Rows) *Base {
    var (
        b *Base
        l *Location
        c *Center
        a *Lab
        r *Armory
        s *Supply
        t *Stope
        o *Storage
    )
    e := rows.ScanRow(&b, &l, &c, &a, &r, &s, &t, &o)
    if e != nil {
        orm.Logger.Println(e)
        return nil
    }
    b.location = l
    l.base = b
    if c.Id > 0 {
        b.center = c
        c.base = b
    }
    if a.Id > 0 {
        b.lab = a
        a.base = b
    }
    if r.Id > 0 {
        b.armory = r
        a.base = b
    }
    if s.Id > 0 {
        b.supply = s
        s.base = b
    }
    if t.Id > 0 {
        b.stope = t
        t.base = b
    }
    if o.Id > 0 {
        b.storage = o
        o.base = b
    }
    conf := Conf.Tree(fmt.Sprintf("base.%d", b.Star))
    if conf != nil {
        b.SelfSupply, _ = conf.Int64("supply")
        b.SelfStorage, _ = conf.Int64("storage")
    }
    return b
}

func (m *baseModel) Base(id int64) *Base {
    rows, e := baseQueryStmt().Where("b.id = ?").Query(id)
    if e != nil {
        orm.Logger.Println(e)
        return nil
    }
    return scanBaseEntity(rows)
}

func (m *baseModel) BaseByXY(x, y int64) *Base {
    rows, e := baseQueryStmt().Where("b.x = ? AND b.y = ?").Query(x, y)
    if e != nil {
        orm.Logger.Println(e)
        return nil
    }
    return scanBaseEntity(rows)
}
