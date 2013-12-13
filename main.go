package main


import (
    "github.com/roydong/potato"
    "github.com/roydong/battle/controller"
)

func init() {
    potato.Init()

    //define template funcs
    potato.T.Funcs(map[string]interface{} {
    })

    //the map keys here must corresponds with 
    //the controller configured in routes.yml
    potato.R.Controllers(map[string]interface{} {
        "main": new(controller.Main),
    })
}

func main() {
    potato.Serve()
}



