package model

import (
    "crypto/rand"
    "crypto/sha512"
    "encoding/hex"
    "github.com/roydong/potato/orm"
    "io"
    "time"
)

type User struct {
    Id        int64     `column:"id"`
    Passwd    string    `column:"passwd"`
    Salt      string    `column:"salt"`
    Name      string    `column:"name"`
    Email     string    `column:"email"`
    CreatedAt time.Time `column:"created_at"`
    UpdatedAt time.Time `column:"updated_at"`
}

func (u *User) SetPasswd(passwd string) {
    rnd := make([]byte, 32)
    if _, e := io.ReadFull(rand.Reader, rnd); e != nil {
        panic("could not generate random salt")
    }

    hash := sha512.New()
    if _, e := hash.Write(rnd); e != nil {
        panic("could not hash salt")
    }

    u.Salt = hex.EncodeToString(hash.Sum(nil))[:32]
    u.Passwd = UserModel.HashPasswd(passwd, u.Salt)
}

func (u *User) CheckPasswd(passwd string) bool {
    return UserModel.HashPasswd(passwd, u.Salt) == u.Passwd
}

/**
 * user model
 */
type userModel struct {
    *orm.Model
}

var UserModel = &userModel{orm.NewModel("user", new(User))}

func (m *userModel) User(id int64) *User {
    var u *User
    rows, e := orm.NewStmt().Select("u.*").
        From("User", "u").Where("u.id = ?").Query(id)

    if e == nil && rows.Next() {
        rows.ScanEntity(&u)
    }

    return u
}

func (m *userModel) User(id int64) *User {
    var u *User
    rows, e := orm.NewStmt().Select("u.*").
        From("User", "u").Where("u.id = ?").Query(id)

    if e == nil && rows.Next() {
        rows.ScanEntity(&u)
    }

    return u
}

func (m *userModel) UserByEmail(email string) *User {
    var u *User
    rows, e := orm.NewStmt().Select("u.*").
        From("User", "u").Where("u.email = ?").Query(email)

    if e == nil && rows.Next() {
        rows.ScanEntity(&u)
    }

    return u
}

func (m *userModel) Exists(email string) bool {
    n, _ := orm.NewStmt().Count("User", "u").
        Where("u.email = ?").Exec(email)

    return n > 0
}

func (m *userModel) HashPasswd(passwd string, salt string) string {
    hash := sha512.New()

    if _, e := hash.Write([]byte(passwd + salt)); e != nil {
        panic("could not hash password")
    }

    return hex.EncodeToString(hash.Sum(nil))
}
