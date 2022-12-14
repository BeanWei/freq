package middleware

import (
	"context"

	"github.com/BeanWei/freq/internal/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func ErrorHandler() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)

		if err := c.Errors.Last(); err != nil {
			if bizErr, ok := err.Meta.(*biz.BizError); ok {
				c.JSON(biz.Code2HttpCode(bizErr.HttpCode), biz.RespFail(
					bizErr.BizCode, bizErr.Msg,
				))
			} else {
				c.JSON(consts.StatusInternalServerError, biz.RespFail(
					biz.CodeServerError,
				))
			}
		}
	}
}
