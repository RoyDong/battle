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
    Id        int64     `column:"id" json:"id"`
    Passwd    string    `column:"passwd" json:"-"`
    Salt      string    `column:"salt" json:"-"`
    Name      string    `column:"name" json:"name"`
    Email     string    `column:"email" json:"email"`
    Gold      int64     `column:"gold" json:"gold"`
    CreatedAt time.Time `column:"created_at" json:"created_at"`
    UpdatedAt time.Time `column:"updated_at" json:"updated_at"`

    roles map[string]*Role
    bases map[int64]*Base
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
    ret := UserRoleModel.Save(u, r)
    if ret && len(u.roles) > 0 {
        u.roles[r.Name] = r
    }
    return ret
}

func (u *User) RemoveRole(r *Role) bool {
    ret := UserRoleModel.Remove(u, r)
    if ret && len(u.roles) > 0 {
        delete(u.roles, r.Name)
    }
    return ret
}

func (u *User) Roles() map[string]*Role {
    if u.roles == nil {
        rows, e := orm.NewStmt().
            Select("r.*").From("Role", "r").
            InnerJoin("UserRole", "ur", "r.id = ur.role_id").
            Where("ur.user_id = ?").
            Query(u.Id)
        if e != nil {
            orm.Logger.Println(e)
            return nil
        }
        u.roles = make(map[string]*Role)
        for rows.Next() {
            var r *Role
            rows.ScanEntity(&r)
            u.roles[r.Name] = r
        }
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

func (u *User) AddBase(base *Base) {
    if u.bases != nil {
        u.bases[base.Id] = base
    }
}

func (u *User) Bases() map[int64]*Base {
    if u.bases == nil {
        rows, e := orm.NewStmt().
            Select("b.*").
            From("Base", "b").
            Where("b.user_id = ?").
            Query(u.Id)
        if e != nil {
            orm.Logger.Println(e)
            return nil
        }
        u.bases = make(map[int64]*Base)
        for rows.Next() {
            var b *Base
            rows.ScanEntity(&b)
            u.bases[b.Id] = b
        }
    }
    return u.bases
}

func (u *User) Base(id int64) *Base{
    if bases := u.Bases(); bases != nil {
        return bases[id]
    }
    return nil
}

/**
 * user model
 */
type userModel struct {
    *orm.Model
}

func (m *userModel) Search(key string, order map[string]string, limit, offset int64) []*User {
    key = "%"+key+"%"
    stmt := orm.NewStmt().
        Select("u.*").From("User", "u").
        Where("u.name like ? or u.email like ?").
        Limit(limit).
        Offset(offset)
    for k, v := range order {
        stmt.OrderBy("u."+k, v)
    }
    rows, e := stmt.Query(key, key)
    if e != nil {
        orm.Logger.Println(e)
        return nil
    }
    users := make([]*User, 0, limit)
    for rows.Next() {
        var u *User
        rows.ScanEntity(&u)
        users = append(users, u)
    }
    return users
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

    if e != nil {
        orm.Logger.Println(e)
        return nil
    }

    rows.ScanRow(&u)
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
