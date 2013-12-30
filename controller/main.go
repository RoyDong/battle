package controller

import (
    "log"
    "time"
    "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
)

type Main struct {
    potato.Controller
}

type Topic struct {
    Id int64 `id`
    Title string `title`
    Content string `content`
    State int `state`
    CreatedAt time.Time `created_at`
    UpdatedAt time.Time `updated_at`
}

type topicModel struct {
    *orm.Model
}

var TopicModel = &topicModel{orm.NewModel("topic", new(Topic))}

func (c *Main) Index() {
    stmt := orm.NewStmt()
    rows := stmt.Select("t.id, t.title, t.content, t.state, t.created_at, t.updated_at").From("Topic", "t").Where("t.id = :id").Query(map[string]interface{} {"id": 1})
    var t Topic

    e := rows.Scan(&t)

    log.Println(t, e)
    c.Response.Write([]byte(stmt.String()))
}

func (c *Main) Ws() {
    for {
        txt := c.WSReceive()

        if len(txt) == 0 {
            return
        }

        c.WSSendJson(map[string]interface{} {
            "message": "ok",
            "error": 0,
            "data": txt,
        })
    }
}
