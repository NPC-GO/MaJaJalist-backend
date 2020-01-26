package database

import (
	"fmt"
	"github.com/NPC-GO/MaJaJalist-backend/model"
	"github.com/go-pg/pg/v9"
)

type User struct {
	DB *pg.DB
}

func (u *User) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}
func (u *User) GetUserByToken(token string) (*model.User, error) {
	return u.GetUserByField("token", token)
}
