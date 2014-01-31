package model

import (
    "github.com/roydong/potato/orm"
    "math/rand"
    "strings"
    "fmt"
    "sync"
    "time"
)

const (
    MapGeoLand = 0
    MapGeoSea  = 1

    MapRefreshInterval = time.Second

    RefreshStateOff = 0
    RefreshStateOn  = 1
)

var (
    LocResourceRate = 5
    LocEnergyRate   = 30

    LocResourceMax = 5000
    LocResourceMin = 1000
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

    RefreshState int
}

func newMapModel() *mapModel {
    m := &mapModel{
        Model:   orm.NewModel("map", &Location{}),
        mlocker: &sync.Mutex{},
        elocker: &sync.Mutex{},
        RefreshState: RefreshStateOff,
    }

    m.Sum([]string{"metal", "energy"}, &m.metal, &m.energy)
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
    for now := range time.Tick(MapRefreshInterval) {
        if m.RefreshState == RefreshStateOff {
            continue
        }

        rows, e := orm.NewStmt().
            Select("l.*").From("Location", "l").
            Where("l.base_id = 0").
            Asc("l.refresh_at").Asc("l.x").Asc("l.y").
            Limit(10).Query()

        if e != nil {
            orm.L.Println(e)
            continue
        }

        for rows.Next() {
            var loc *Location
            if rows.ScanEntity(&loc) != nil {
                continue
            }

            metal, energy := loc.Metal, loc.Energy
            if rand.Intn(100) < LocResourceRate {
                n := rand.Intn(100)
                total := int64(LocResourceMin +
                    (LocResourceMax - LocResourceMin) * n / 100)

                if rand.Intn(100) < LocEnergyRate {
                    loc.Energy = total * (int64(rand.Intn(10)) + 1) / 10
                } else {
                    loc.Energy = 0
                }

                loc.Metal = total - loc.Energy
            } else {
                loc.Metal = 0
                loc.Energy = 0
            }

            m.IncrMetal(loc.Metal - metal)
            m.IncrEnergy(loc.Energy - energy)
            loc.RefreshAt = now
            m.SaveResource(loc)
        }
    }
}

func (m *mapModel) Rect(x, y, w, h int64) []*Location {
    rows, e := orm.NewStmt().
        Select("l.*").From("Location", "l").
        Where("l.x BETWEEN ? AND ? AND l.y BETWEEN ? AND ?").
        Query(x, x + w, y, y + h)

    locs := make([]*Location, 0)
    if e != nil {
        orm.L.Println(e)
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

func (m *mapModel) SaveResource(loc *Location) bool {
    _, e := orm.NewStmt().
        Update("Location", "l", "metal", "energy", "refresh_at").
        Where("l.x = ? AND l.y = ?").
        Exec(loc.Metal, loc.Energy, loc.RefreshAt, loc.X, loc.Y)

    if e != nil {
        orm.L.Println(e)
        return false
    }

    return true
}

func (m *mapModel) Sum(cols []string, nums ...interface{}) error {
    sums := make([]string, 0, len(cols))
    for _, col := range cols {
        sums = append(sums, fmt.Sprintf("SUM(%s)", col))
    }

    row := orm.D.QueryRow(
        "SELECT " + strings.Join(sums, ", ") + " FROM map")
    return row.Scan(nums...)
}
