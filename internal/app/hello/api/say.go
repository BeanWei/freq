package api

import (
	"context"

	"github.com/BeanWei/freq/internal/pkg/biz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Say(ctx context.Context, c *app.RequestContext) {
	type SayReq struct {
		Ticket  string `query:"ticket,required" vd:"len($)>1 && len($)<3"`
		Ticket2 string `query:"ticket2,required"`
	}
	var req SayReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithError(consts.StatusBadRequest, biz.NewParamBindError(err))
		return
	}
	c.JSON(consts.StatusOK, biz.RespSuccess(utils.H{
		"msg": "你好",
	}))
}
