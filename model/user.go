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

    roles []*Role
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

func (u *User) AddRole(r *Role) bool {
    return UserRoleModel.Save(u, r)
}

func (u *User) RemoveRole(r *Role) bool {
    return UserRoleModel.Remove(u, r)
}

func (u *User) Roles() []*Role {
    if u.roles == nil {
        rows, e := orm.NewStmt().
            Select("r.*").From("Role", "r").
            InnerJoin("UserRole", "ur", "r.id = ur.role_id").
            Where("ur.user_id = ?").
            Query(u.Id)

        roles := make([]*Role, 0)
        if e != nil {
            orm.L.Println(e)
            return roles
        }

        for rows.Next() {
            var r *Role
            rows.ScanEntity(&r)
            roles = append(roles, r)
        }

        u.roles = roles
    }

    return u.roles
}

func (u *User) IsGrantd(roles ...string) bool {
    for _, r := range u.Roles() {
        for _, name := range roles {
            if name == r.Name {
                return true
            }
        }
    }

    return false
}

func (u *User) IsGrantedAll(roles ...string) bool {
    tags := make([]bool, len(roles))
    for _, r := range u.Roles() {
        for i, name := range roles {
            if name == r.Name {
                tags[i] = true
            }
        }
    }

    for _, tag := range tags {
        if !tag {
            return false
        }
    }

    return true
}

/**
 * user model
 */
type userModel struct {
    *orm.Model
}

func (m *userModel) User(id int64) *User {
    var u *User
    rows, e := orm.NewStmt().Select("u.*").
        From("User", "u").Where("u.id = ?").Query(id)

    if e == nil {
        rows.ScanRow(&u)
    }

    return u
}

func (m *userModel) UserByEmail(email string) *User {
    var u *User
    rows, e := orm.NewStmt().Select("u.*").
        From("User", "u").Where("u.email = ?").Query(email)

    if e == nil {
        rows.ScanRow(&u)
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
