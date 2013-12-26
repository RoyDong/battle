package model

import (
    "io"
    "time"
    "crypto/rand"
    "crypto/md5"
    "crypto/sha512"
    "encoding/hex"
    "github.com/roydong/potato"
)

type User struct {
    id int64
    passwd, salt string

    Name, Email string
    CreatedAt, UpdatedAt time.Time
}

func (u *User) Id() int64 {
    return u.id
}

/**
 * set a hash password
 */
func (u *User) SetPasswd(passwd string) {
    rnd := make([]byte, 32)
    if _,e := io.ReadFull(rand.Reader, rnd); e != nil {
        panic("could not generate random salt")
    }

    hash := md5.New()
    if _, e := hash.Write(rnd); e != nil {
        panic("could not hash salt")
    }

    u.salt = hex.EncodeToString(hash.Sum(nil))
    u.passwd = UserModel.HashPasswd(passwd, u.salt)
}

func (u *User) CheckPasswd(passwd string) bool {
    return UserModel.HashPasswd(passwd, u.salt) == u.passwd
}


type userModel struct {Model}

var UserModel = &userModel{
    Model{"user", []string{
            "id","email","name","passwd","salt","created_at","updated_at"}}
}

func (m *userModel) User(id int64) *User {
    row := m.FindOne(map[string]interface{"id": id}, nil)
    return m.loadData(row)
}

func (m *userModel) UserByEmail(email string) *User {
    row := m.FindOne(map[string]interface{"email": email}, nil)
    return m.loadData(row)
}

func (m *userModel) loadData(row potato.Scanner) *User {
    u := new(User)
    var ct, ut int64
    if e := row.Scan(&u.id, &u.Email, &u.Name, &u.passwd, &u.salt, &ct, &ut); e != nil {
        return nil
    }

    u.CreatedAt = time.Unix(0, ct)
    u.UpdatedAt = time.Unix(0, ut)
    return u
}

func (m *userModel) Save(u *User) bool {
    data := map[string]interface{} {
        "email": u.Email,
        "name": u.Name,
        "passwd": u.passwd,
        "salt": u.salt,
        "created_at": u.CreatedAt.UnixNano(),
        "updated_at": u.UpdatedAt.UnixNano(),
    }

    if u.Id() > 0 {
        return m.Update(data, map[string]interface{}{"id": u.id}) > 0
    }

    u.id = m.Insert(data)
    return u.id > 0
}

func (m *userModel) HashPasswd(passwd string, salt string) string {
    hash := sha512.New()

    if _, e := hash.Write([]byte(passwd + salt)); e != nil {
        panic("could not hash password")
    }

    return hex.EncodeToString(hash.Sum(nil))
}