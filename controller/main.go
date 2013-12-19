package controller

import (
    "time"
    "github.com/roydong/potato"
)

type Main struct {
    potato.Controller
}

func (c *Main) Index() {
    c.RenderJson("111")
    
    time.Sleep(time.Second * 10)
}