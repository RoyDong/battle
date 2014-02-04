package main

import (
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/roydong/battle/controller"
    _ "github.com/roydong/battle/controller/admin"
    "github.com/roydong/potato"
)

func main() {
    potato.Init()
    potato.Serve()
}
