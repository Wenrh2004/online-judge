package mysql

import (
	"online-judge/biz/dal/mysql/query"
	"online-judge/biz/model"
)

func QueryUserByUsername(username string) ([]*model.User, error) {
	userQuery := query.Use(DBHook).User
	users, err := userQuery.Where(userQuery.Username.Eq(username)).Find()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(user *model.User) (*string, error) {
	userCreate := query.Use(DBHook).User
	err := userCreate.Create(user)
	if err != nil {
		return nil, err
	}

	return &user.ID, nil
}

func QueryUserByID(userID string) (*model.User, error) {
	userQuery := query.Use(DBHook).User
	user, err := userQuery.Where(userQuery.ID.Eq(userID)).First()
	if err != nil {
		return nil, err
	}
	user.Password = ""

	return user, nil

}
