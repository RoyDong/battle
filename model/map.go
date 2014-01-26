package model

import (
    "github.com/roydong/potato/orm"
    "math/rand"
    "sync"
    "time"
)

const (
    MapGeoLand = 0
    MapGeoSea  = 1

    MapSizeX = 1000
    MapSizeY = 1000

    MapRefreshInterval = 500 * time.Millisecond
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
    X      int64 `column:"x" json:"x"`
    Y      int64 `column:"y" json:"y"`
    Geo    int   `column:"geo" json:"geo"`
    BaseId int64 `column:"base_id" json:"base_id"`

    Metal  int64 `column:"metal" json:"metal"`
    Energy int64 `column:"energy" json:"energy"`

    UpdatedAt time.Time `column:"updated_at" json:"updated_at"`
    RefreshAt time.Time `column:"refresh_at" json:"refresh_at"`
}

/**
 * map model
 */
type mapModel struct {
    *orm.Model

    metal   int64
    mlocker *sync.Mutex

    energy  int64
    elocker *sync.Mutex
}

func newMapModel() *mapModel {
    m := &mapModel{
        Model:   orm.NewModel("map", &Location{}),
        mlocker: &sync.Mutex{},
        elocker: &sync.Mutex{},
    }

    m.metal = m.Sum("metal")
    m.energy = m.Sum("energy")
    go m.refresh()
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

func (m *mapModel) refresh() {
    rate := LocResourceRate
    num := 0

    for now := range time.Tick(MapRefreshInterval) {
        rows, e := orm.NewStmt().
            Select("l.*").From("Location", "l").
            Where("l.base_id = 0").
            Asc("l.refresh_at").Asc("l.x").Asc("l.y").
            Limit(1).Query()

        var loc *Location
        if e != nil || rows.ScanRow(&loc) != nil {
            continue
        }

        metal, energy := loc.Metal, loc.Energy
        if rand.Intn(100) < rate {
            rate = LocResourceRate
            num = 0

            if rand.Intn(100) < LocMetalRate {
                n := rand.Intn(100)
                loc.Metal = int64(LocMetalMin +
                    (LocMetalMax - LocMetalMin) * n / 100)
            }

            if rand.Intn(100) < LocEnergyRate {
                n := rand.Intn(100)
                loc.Energy = int64(LocEnergyMin +
                    (LocEnergyMax - LocEnergyMin) * n / 100)
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
        loc.RefreshAt = now
        m.Save(loc)
    }
}

func (m *mapModel) Rect(x, y, r int64) []*Location {
    rows, e := orm.NewStmt().
        Select("l.*").From("Location", "l").
        Where("l.x >= ? AND l.x <= ? AND l.y >= ? AND l.y <= ?").
        Query(x - r, x + r, y - r, y + r)

    locs := make([]*Location, 0)
    if e != nil {
        return locs
    }

    for rows.Next() {
        var loc *Location
        rows.ScanEntity(&loc)
        locs = append(locs, loc)
    }

    return locs
}

func (m *mapModel) Location(x, y int64) *Location {
    var loc *Location
    rows, e := orm.NewStmt().
        Select("l.*").
        From("Location", "l").
        Where("l.x = ? AND l.y = ?").
        Query(x, y)

    if e != nil {
        rows.ScanRow(&loc)
    }

    return loc
}

func (m *mapModel) Save(loc *Location) bool {
    _, e := orm.NewStmt().
        Update("Location", "l", "geo", "base_id",
        "metal", "energy", "updated_at", "refresh_at").
        Where("l.x = ? AND l.y = ?").
        Exec(loc.Geo, loc.BaseId, loc.Metal, loc.Energy,
        loc.UpdatedAt, loc.RefreshAt, loc.X, loc.Y)

    if e != nil {
        orm.L.Println(e)
        return false
    }

    return true
}

func (m *mapModel) Sum(col string) int64 {
    row := orm.D.QueryRow("SELECT SUM(" + col + ") n FROM map")
    var n int64
    row.Scan(&n)
    return n
}
