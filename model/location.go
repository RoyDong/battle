package model

const (
    GeoLand = 0
    GeoSea  = 1
)

type Location struct {
    X int64 `column:"x"`
    Y int64 `column:"y"`
    Geo int `column:"geo"`
    FortId int64 `column:"fort_id"`
}
