package lib

import (
    "github.com/roydong/potato"
    "github.com/roydong/battle/model"
)

type Auth struct {
    user *model.User
}

func NewAuth(s *potato.Session) *Auth {
    auth := &Auth{}
    if u, ok := s.Value("user").(*model.User); ok {
        auth.SetUser(u)
    }
    return auth
}

func (a *Auth) User() *model.User {
    return a.user
}

func (a *Auth) SetUser(u *model.User) {
    a.user = u
}

