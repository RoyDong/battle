package model

import (
    "fmt"
    "github.com/roydong/potato/orm"
    "math/rand"
    "strings"
    "sync"
    "time"
    "log"
)

const (
    MapGeoSea  = 0
    MapGeoLand = 1

    MapRefreshInterval = time.Second

    RefreshStateOff = 0
    RefreshStateOn  = 1
)

var (
    ResourceRate = 5
    EnergyRate   = 30
    ResourceMax = 5000
    ResourceMin = 1000
)

type Location struct {
    X      int   `column:"x" json:"x"`
    Y      int   `column:"y" json:"y"`
    Geo    int   `column:"geo" json:"geo"`

    Metal  int64 `column:"metal" json:"metal"`
    Energy int64 `column:"energy" json:"energy"`

    UpdatedAt time.Time `column:"updated_at" json:"updated_at"`
    RefreshAt time.Time `column:"refresh_at" json:"refresh_at"`

    base *Base
}

func (loc *Location) Key() string {
    return fmt.Sprintf("%d,%d", loc.X, loc.Y)
}

var locLocker = &sync.Mutex{}
var lockedLocs = make(map[string]*Location)
func (loc *Location) Lock() bool {
    locLocker.Lock()
    defer locLocker.Unlock()
    if _, has := lockedLocs[loc.Key()]; has {
        return false
    }
    lockedLocs[loc.Key()] = loc
    return true
}

func (loc *Location) Unlock() {
    delete(lockedLocs, loc.Key())
}

func (loc *Location) Base() *Base {
    if loc.base == nil {
        rows, e := orm.NewStmt("").
            Select("b.*").
            From("Base", "b").
            Where("b.x = ? AND b.y = ?").
            Query(loc.X, loc.Y)

        if e == nil {
            rows.ScanRow(&loc.base)
        }
    }
    return loc.base
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
        Model:        orm.NewModel("map", &Location{}),
        mlocker:      &sync.Mutex{},
        elocker:      &sync.Mutex{},
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

        rows, e := orm.NewStmt("").
            Select("l.*").From("Location", "l").
            Where("l.base_id = 0").
            Asc("l.refresh_at").Asc("l.x").Asc("l.y").
            Limit(10).Query()

        if e != nil {
            log.Println(e)
            continue
        }

        for rows.Next() {
            var loc *Location
            if rows.ScanEntity(&loc) != nil {
                continue
            }

            metal, energy := loc.Metal, loc.Energy
            if rand.Intn(100) < ResourceRate {
                n := rand.Intn(100)
                total := int64(ResourceMin +
                    (ResourceMax - ResourceMin) * n / 100)

                if rand.Intn(100) < EnergyRate {
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

func (m *mapModel) Rect(x, y, w, h int) []*Location {
    rows, e := orm.NewStmt("").
        Select("l.*").From("Location", "l").
        Where("l.x BETWEEN ? AND ? AND l.y BETWEEN ? AND ?").
        Query(x, x+w, y, y+h)

    locs := make([]*Location, 0)
    if e != nil {
        log.Println(e)
        return locs
    }

    for rows.Next() {
        var loc *Location
        rows.ScanEntity(&loc)
        locs = append(locs, loc)
    }

    return locs
}

func (m *mapModel) Location(x, y int) *Location {
    var loc *Location
    var base *Base
    rows, e := orm.NewStmt("").
        Select("l.*,b.*").
        From("Location", "l").
        LeftJoin("Base", "b", "b.x = l.x AND b.y = l.y").
        Where("l.x = ? AND l.y = ?").
        Query(x, y)
    if e != nil {
        log.Println(e)
        return nil
    }
    rows.ScanRow(&loc, &base)
    if base.Id > 0 {
        loc.base = base
    }
    return loc
}

func (m *mapModel) SaveResource(loc *Location) bool {
    _, e := orm.NewStmt("").
        Update("Location", "l", "metal", "energy", "refresh_at").
        Where("l.x = ? AND l.y = ?").
        Exec(loc.Metal, loc.Energy, loc.RefreshAt, loc.X, loc.Y)

    if e != nil {
        log.Println(e)
        return false
    }

    return true
}

func (m *mapModel) Sum(cols []string, nums ...interface{}) error {
    sums := make([]string, 0, len(cols))
    for _, col := range cols {
        sums = append(sums, fmt.Sprintf("SUM(%s)", col))
    }

    row := orm.SqlDB("").QueryRow(
        "SELECT " + strings.Join(sums, ", ") + " FROM map")
    return row.Scan(nums...)
}
