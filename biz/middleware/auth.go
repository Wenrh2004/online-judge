package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"online-judge/biz/model/entity"
	"online-judge/biz/utils"
	"online-judge/biz/utils/auth"
)

func Auth() []app.HandlerFunc {
	return []app.HandlerFunc{func(ctx context.Context, c *app.RequestContext) {
		accessToken := c.GetHeader("Authorization")

		token := entity.TokenPair{
			AccessToken: string(accessToken),
		}

		userClaims, err := auth.ParseTokenString(token.AccessToken)
		if err != nil {
			c.Abort()
			utils.Response(ctx, c, err)
		}

		c.Set("user", userClaims.UserID)
		c.Next(ctx)
	}}
}
