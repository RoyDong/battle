package model

import (
    "github.com/roydong/potato/orm"
    "math/rand"
    "sync"
    "time"
)

const (
    GeoLand = 0
    GeoSea  = 1

    MapSizeX = 1000
    MapSizeY = 1000

    MapRefreshInterval = 0.5 * time.Second
)

var (
    LocResourceRate = 5
    LocMetalRate    = 70
    LocEnergyRate   = 30

    LocMetalMax = 500
    LocMetalMin = 100

    LocEnergyMax = 500
    LocEnergyMin = 100
)

type Location struct {
    X      int64 `column:"x"`
    Y      int64 `column:"y"`
    Geo    int   `column:"geo"`
    BaseId int64 `column:"base_id"`

    Metal  int64 `column:"metal"`
    Energy int64 `column:"energy"`

    CreatedAt time.Time `column:"created_at"`
    UpdatedAt time.Time `column:"updated_at"`
    RefreshAt time.Time `column:"refresh_at"`
}

/**
 * map model
 */
type mapModel struct {
    *orm.Model

    metal   int64
    mlocker *symc.Mutex

    energy  int64
    elocker *symc.Mutex
}

var MapModel = &mapModel{orm.NewModel("map", &Location{})}

func newMapModel(model *orm.Model) *mapModel {
    m := &mapModel{
        Model:   model,
        mlocker: &sync.Mutex{},
        elocker: &sync.Mutex{},
    }

    m.metal = m.Sum("metal")
    m.energy = m.Sum("energy")
    go m.Refresh()
    return m
}

func (m *mapModel) IncrMetal(n int64) {
    m.mlocker.Lock()
    m.metal += n
    m.mlocker.Unlock()
}

func (m *mapModel) Metal() int64 {
    return m.metal
}

func (m *mapModel) IncrEnergy(n int64) {
    m.elocker.Lock()
    m.energy += n
    m.elocker.Unlock()
}

func (m *mapModel) Energy() int64 {
    return m.energy
}

func (m *mapModel) Resource() int64 {
    return m.metal + m.energy
}

func (m *mapModel) Refresh() {
    rate := LocResourceRate
    num := 0

    for now := range time.Tick(MapRefreshInterval) {
        rows, e := orm.NewStmt().
            Select("l.*").From("Location", "l").
            Where("l.base_id = 0").
            Asc("l.refresh_at").Asc("l.x").Asc("l.y").
            Limit(1).Query()

        var loc *Location
        if e == nil && rows.Next() {
            rows.ScanEntity(&loc)
        } else {
            continue
        }

        metal, energy := loc.Metal, loc.Energy
        if rand.Intn(100) < rate {
            rate = LocResourceRate
            num = 0

            if rand.Intn(100) < LocMetalRate {
                n := rand.Intn(100)
                loc.Metal = LocMetalMin +
                    (LocMetalMax-LocMetalMin)*n/100
            }

            if rand.Intn(100) < LocEnergyRate {
                n := rand.Intn(100)
                loc.Energy = LocEnergyMin +
                    (LocEnergyMax-LocEnergyMin)*n/100
            }

        } else {
            num++
            if num%20 == 0 {
                rate++
            }

            loc.Metal = 0
            loc.Energy = 0
        }

        m.IncrMetal(loc.Metal - metal)
        m.IncrEnergy(loc.Energy - energy)
        loc.RefreshAt = time.Now()
        m.Save(loc)
    }
}

func (m *mapModel) Location(x, y int64) *Location {
    var loc *Location
    rows, e := orm.NewStmt().
        Select("l.*").
        From("Location", "l").
        Where("l.x = ? AND l.y = ?").
        Query(x, y)

    if e == nil && rows.Next() {
        rows.ScanEntity(&loc)
    }

    return loc
}

func (m *mapModel) Save(loc *Location) bool {
    _, e := orm.NewStmt().
        Update("Location", "l",
        "geo", "base_id", "metal", "energy",
        "created_at", "updated_at", "refresh_at").
        Where("l.x = ? AND l.y = ?").
        Exec(loc.Geo, loc.BaseId,
        loc.Metal, loc.Energy,
        loc.CreatedAt, loc.UpdatedAt, loc.RefreshAt)

    return e == nil
}

func (m *mapModel) Sum(col) int64 {
    rows, e := orm.D.Query("SELECT SUM(" + col + ") n FROM map")
    var n int64
    if e == nil && rows.Next() {
        rows.Scan(&n)
    }
    return n
}
