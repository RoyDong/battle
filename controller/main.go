package controller

import (
    _ "github.com/roydong/battle/model"
    pt "github.com/roydong/potato"
)

func init() {
    pt.SetAction(func(r *pt.Request, p *pt.Response) error {
        for {
            txt := r.WSReceive()

            if txt != "" {
                r.WSSend(txt)
            } else {
                break
            }
        }

        return nil
    }, "/ws")
}
