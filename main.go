package main

import (
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/roydong/battle/ctrl"
    "github.com/roydong/potato"
)

func init() {
    potato.Init()
}

func main() {
    potato.T.SetFuncs(map[string]interface{}{})
    potato.Serve()
}
