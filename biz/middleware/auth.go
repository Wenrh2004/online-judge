package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"online-judge/biz/model/entity"
	"online-judge/biz/utils/auth"
)

func Auth() []app.HandlerFunc {
	return []app.HandlerFunc{func(ctx context.Context, c *app.RequestContext) {
		accessToken := c.GetHeader("Authorization")
		// refreshToken := c.GetHeader("RefreshToken")

		token := entity.TokenPair{
			AccessToken: string(accessToken),
			// RefreshToken: string(refreshToken),
		}

		userClaims, err := auth.ParseTokenString(token.AccessToken)
		if err != nil {
			c.Abort()
			c.JSON(consts.StatusBadRequest, err.Error())
		}

		// userInfo, err := mysql.QueryUserByID(userClaims.UserID)
		// if err != nil {
		// 	c.Abort()
		// 	c.JSON(consts.StatusBadRequest, err.Error())
		// }

		c.Set("user", userClaims.UserID)
		c.Next(ctx)
	}}
}
