package controller

import (
    "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
    "log"
    "time"
)

type Main struct {
    potato.Controller
}

type Topic struct {
    Id        int64     `column:"id" json:"id"`
    Title     string    `column:"title" json:"title"`
    Content   string    `column:"content" json:"content"`
    State     int       `column:"state" json:"state"`
    CreatedAt time.Time `column:"created_at" json:"created_at"`
    UpdatedAt time.Time `column:"updated_at" json:"updated_at"`
}

type topicModel struct {
    *orm.Model
}

var TopicModel = &topicModel{orm.NewModel("topic", new(Topic))}

func (c *Main) Index() {
    rows, e := orm.NewStmt().Select("t.*").From("Topic", "t").Desc("t.id").Query()
    if e != nil {
        log.Println(e, 1)
    }

    orm.NewStmt().Insert("Topic", "title", "content").Exec("aaa", "ddddd")
    orm.NewStmt().Update("Topic", "t", "title", "content").Where("id=?").Exec("aadd", "ccccccc", 13)

    topics := make([]*Topic, 0, 5)
    for rows.Next() {
        var t *Topic

        rows.ScanEntity(&t)

        t.State = 2
        if t.Id == 12 {
            t.Id = 103
            t.Title = potato.RandString(10)
            t.UpdatedAt = time.Now()
            log.Println(orm.Save(t))
        }

        topics = append(topics, t)
    }

    c.RenderJson(topics)
}

func (c *Main) Ws() {
    for {
        txt := c.WSReceive()

        if len(txt) == 0 {
            return
        }

        c.WSSendJson(map[string]interface{}{
            "message": "ok",
            "error":   0,
            "data":    txt,
        })
    }
}
