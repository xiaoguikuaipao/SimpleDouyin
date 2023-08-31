package api

import (
	"context"

	"SimpleDouyin/web/model"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterHandler(c context.Context, ctx *app.RequestContext) {
	registerInfo := &model.User{}
	err := ctx.BindAndValidate(registerInfo)
	if err != nil {
		ctx.JSON(consts.StatusBadRequest, utils.H{
			"message": "用户名或密码不符合要求",
		})
	}

}

func LoginHandler(ctx context.Context, c *app.RequestContext) {

}
