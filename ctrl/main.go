package ctrl

import (
    "github.com/roydong/potato"
)

func init() {
    potato.SetAction("", func(r *potato.Request, p *potato.Response) {
        p.RenderText("aaa")
    })
}
