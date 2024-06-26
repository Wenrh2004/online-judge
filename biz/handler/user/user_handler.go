// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"online-judge/biz/dal/mysql"
	"online-judge/biz/model/user"
	userService "online-judge/biz/service/user"
	"online-judge/biz/utils"
)

// Register .
// @router /api/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.Response(ctx, c, err)
		return
	}

	userIDByRegister, err := userService.RegisterService(&req)
	if err != nil {
		utils.Response(ctx, c, err)
		return
	}
	resp := new(user.RegisterResp)
	resp.Id = *userIDByRegister

	utils.Response(ctx, c, nil, resp)
}

// Login .
// @router /api/user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.Response(ctx, c, err)
		return
	}

	userByLogin, err := userService.LoginService(&req)
	if err != nil {
		utils.Response(ctx, c, err)
		return
	}

	resp := new(user.LoginResp)
	resp.Data = userByLogin

	utils.Response(ctx, c, nil, resp)
}

// CurrentUser .
// @router /api/user/current [GET]
func CurrentUser(ctx context.Context, c *app.RequestContext) {
	u := c.MustGet("user").(string)

	userInfo, err := mysql.QueryUserByID(u)
	if err != nil {
		utils.Response(ctx, c, err)
	}

	resp := user.CurrentUserResp{
		Id:       userInfo.ID,
		Username: userInfo.Username,
		Role:     userInfo.Role,
		Avatar:   userInfo.Avatar,
		Email:    userInfo.Email,
		Phone:    userInfo.Phone,
	}

	utils.Response(ctx, c, nil, resp)
}

// Logout .
// @router /api/user/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LogoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.Response(ctx, c, err)
		return
	}

	resp := new(user.LogoutResp)

	utils.Response(ctx, c, nil, resp)
}
