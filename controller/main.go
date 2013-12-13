package controller

import (
    "github.com/roydong/potato"
)

type Main struct {
    potato.Controller
}

func (c *Main) Index() {
    c.RenderJson("111")
}