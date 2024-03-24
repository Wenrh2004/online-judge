package user

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"online-judge/biz/dal/mysql"
	"online-judge/biz/model"
	"online-judge/biz/model/entity"
	"online-judge/biz/model/user"
	"online-judge/biz/utils/auth"
	"strings"
)

func RegisterService(register *user.RegisterReq) (*string, error) {
	// 验证密码是否相等
	if strings.Compare(register.Password, register.CheckPassword) != 0 {
		return nil, errors.New("password not equal")
	}

	// 查询用户是否被注册
	usersByUsername, err := mysql.QueryUserByUsername(register.Username)
	if err != nil {
		return nil, err
	}
	if len(usersByUsername) > 0 {
		return nil, errors.New("user was registered")
	}

	// 创建用户
	// 密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	userRegister := model.User{
		ID:       uuid.New().String(),
		Username: register.Username,
		Password: string(hashPassword),
	}
	createUser, err := mysql.CreateUser(&userRegister)
	if err != nil {
		return nil, err
	}

	return createUser, nil
}

func LoginService(user *user.LoginReq) (*entity.TokenPair, error) {
	// 查询用户是否存在
	usersByUsername, err := mysql.QueryUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	if len(usersByUsername) == 0 {
		return nil, errors.New("user not found")
	}

	// 密码验证
	if err = bcrypt.CompareHashAndPassword([]byte(usersByUsername[0].Password), []byte(user.Password)); err != nil {
		return nil, errors.New("password error")
	}

	tokenPair, err := auth.CreateToken(usersByUsername[0].ID)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}
