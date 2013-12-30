package main


import (
    "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
    "github.com/roydong/battle/controller"
)

func init() {
    potato.Init()
    orm.Init(potato.D, potato.L)

    //define template funcs
    potato.T.SetFuncs(map[string]interface{} {})

    //the map keys here must corresponds with 
    //the controller configured in routes.yml
    potato.R.SetControllers(map[string]interface{} {
        "main": new(controller.Main),
    })
}

func main() {
    potato.Serve()
}
