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
    Id int64 `column:"id" type:"int" json:"id"`
    Title string `column:"title" type:"string" json:"title"`
    Content string `column:"content" type:"string" json:"content"`
    State int `column:"state" type:"int" json:"state"`
    CreatedAt time.Time `column:"created_at" type:"time" json:"created_at"`
    UpdatedAt time.Time `column:"updated_at" type:"time" json"updated_at"`
}

type topicModel struct {
    *orm.Model
}

var TopicModel = &topicModel{orm.NewModel("topic", new(Topic))}

func (c *Main) Index() {
    stmt := orm.NewStmt()
    rows, e := stmt.Select("t.*").From("Topic", "t").Desc("id").Limit(2).Query(nil)

    if e != nil {
        log.Println(e)
    }

    topics := make([]*Topic, 0, 2)
    for rows.Next() {
        var t *Topic
        rows.ScanStruct(&t)
        topics = append(topics, t)
    }

    n := orm.NewStmt().Count("Topic", "t").QueryNum(nil)

    log.Println(n)

    c.RenderJson(topics)
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
