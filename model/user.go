package model

import "time"

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	NickName      string `json:"nickName" pg:"nickname"`
	Avatar        string `json:"avatar"`
	UserLevel     Level  `json:"userLevel" pg:"userlevel"`
	Special       bool   `json:"special"`
	EmailVerified bool   `json:"emailVerified" pg:"emailverified"`
	PassWord      string `pg:"password"`
	Token         string
	LastLogin     time.Time `pg:"lastlogin"`
}

type CreateUserInput struct {
	Email    string `json:"email"`
	ID       string `json:"id"`
	NickName string `json:"nickName" pg:"nickname"`
	Password string `json:"password" pg:"password"`
}
