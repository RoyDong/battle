package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/roydong/potato"
    "github.com/roydong/battle/controller"
    "github.com/roydong/battle/model"
)

func init() {
    potato.Init()

    //init after potato
    model.Init()
}

func main() {
    //define template funcs
    potato.T.SetFuncs(map[string]interface{}{})

    //the map keys here must corresponds with
    //the controller configured in routes.yml
    potato.R.SetControllers(map[string]interface{}{
        "main": &controller.Main{},
        "map": &controller.Map{},
    })

    potato.Serve()
}
