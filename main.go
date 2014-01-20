package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/roydong/battle/controller"
    "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
)

func init() {
    potato.Init()
    orm.InitDefault()
}

func main() {
    //define template funcs
    potato.T.SetFuncs(map[string]interface{}{})

    //the map keys here must corresponds with
    //the controller configured in routes.yml
    potato.R.SetControllers(map[string]interface{}{
        "main": new(controller.Main),
    })

    potato.Serve()
}
