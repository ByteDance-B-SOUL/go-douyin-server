package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterPublishRoute(h *server.Hertz) {
	apiGroup := h.Group("/douyin")
	{
		// 基础接口
		apiGroup.GET("/feed/")
		apiGroup.POST("/user/register/")
		apiGroup.POST("/user/login/")
		apiGroup.GET("/user/")
		apiGroup.POST("/publish/action/")
		apiGroup.POST("/publish/list/")
		// 互动接口
		apiGroup.POST("/favorite/action/")
		apiGroup.GET("/favorite/list/")
		apiGroup.POST("/comment/action/")
		apiGroup.GET("/comment/list/")
		// 社交接口
		apiGroup.POST("/relation/action/")
		apiGroup.GET("/relation/follow/list/")
		apiGroup.GET("/relation/follower/list/")
		apiGroup.GET("/relation/friend/list/")
		apiGroup.GET("/message/chat/")
		apiGroup.POST("/message/action/")
	}
}

func RegisterGroupRouteWithMiddleware(h *server.Hertz) {
	example1 := h.Group("/example1", basic_auth.BasicAuth(map[string]string{"test": "test"}))
	example1.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "ping")
	})

	example2 := h.Group("/example2/v1")
	example2.Use(basic_auth.BasicAuth(map[string]string{"test": "test"}))
	example2.GET("/ping/", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "ping")
	})
}
