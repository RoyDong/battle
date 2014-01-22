package model

import (
    "github.com/roydong/potato/orm"
    "sync"
    "time"
    "math/rand"
)

const (
    ResouceRate = 5
    GeoLand = 0
    GeoSea  = 1

    MapSizeX = 1000
    MapSizeY = 1000
)

type Location struct {
    X      int64 `column:"x"`
    Y      int64 `column:"y"`
    Geo    int   `column:"geo"`
    BaseId int64 `column:"base_id"`

    Metal     int64 `column:"metal"`
    Energy    int64 `column:"energy"`

    CreatedAt time.Time `column:"created_at"`
    UpdatedAt time.Time `column:"updated_at"`
    RefreshAt time.Time `column:"refresh_at"`
}


/**
 * map model
 */
type mapModel struct {
    *orm.Model

    mutex symc.Mutex

    rsNum int
    metalNum int64
    energyNum int64

    lastRefreshLoc *Location

    rate int
    missNum int
}

var MapModel = &mapModel{orm.NewModel("map", &Location{})}

func newMapModel(model *orm.Model) *mapModel {
    m := &mapModel {
        Model: model,
        mutex: &sync.Mutex{},
    }

    return m
}

func (m *mapModel) IncrRsNum(n int64) {
    m.mutex.Lock()
    m.rsNum = m.rsNum + n
    m.mutex.Unlock()
}

func (m *mapModel) IncrMetalNum(n int64) {
    m.mutex.Lock()
    m.metalNum = m.metalNum + n
    m.mutex.Unlock()
}

func (m *mapModel) IncrEnergyNum(n int64) {
    m.mutex.Lock()
    m.energyNum = m.energyNum + n
    m.mutex.Unlock()
}

func (m *mapModel) Refresh() {
    loc := m.LastRefreshLoc()

    //hit rate
    if rand.Intn(100) < rate {
        m.rate = ResouceRate
        missNum = 0
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


func (m *mapModel) lastRefreshLoc() *Location {
    if m.lastRefreshLoc == nil {
        rows, e := orm.NewStmt().
            Select("l.*").
            From("Location", "l").
            Desc("l.refresh_at").
            Limit(1).
            Query()

        if e == nil && rows.Next() {
            rows.ScanEntity(&lastRefreshLoc)
        }
    }

    return m.lastRefreshLoc
}
