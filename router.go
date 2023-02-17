package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Args struct {
	Query      string   `query:"query"`
	QuerySlice []string `query:"q"`
	Path       string   `path:"path"`
	Headers    string   `header:"header"`
	Form       string   `form:"form"`
	Json       string   `json:"json"`
	Vd         int      `query:"vd" vd:"$==0||$==1"`
}

func RegisterPublishRoute(h *server.Hertz) {
	Publish := h.Group("/douyin/Publish")
	{
		Publish.POST("/action/")
		Publish.POST("/list/")
	}
}

func RegisterRelationsRoute(h *server.Hertz) {
	Relation := h.Group("/douyin/relation")
	{
		Relation.POST("/action/")
		Relation.POST("/follow/list/")
		Relation.POST("/follower/list/")
		Relation.POST("/friend/list/")
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

func RegisterParaRoute(h *server.Hertz) {
	h.GET("/hertz/:version", func(ctx context.Context, c *app.RequestContext) {
		version := c.Param("version")
		c.String(consts.StatusOK, "Hello %s", version)
	})

	h.GET("/hertz/:version/*action", func(ctx context.Context, c *app.RequestContext) {
		version := c.Param("version")
		action := c.Param("action")
		message := version + " is " + action
		c.String(consts.StatusOK, message)
	})

	h.POST("/hertz/:version/*action", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, c.FullPath())
	})

	h.GET("/hertz/test/v1?token&time", func(ctx context.Context, c *app.RequestContext) {
		v2 := c.Param("token")
		v3 := c.Param("time")
		message := v3 + " ? " + v2
		c.String(consts.StatusOK, message)
	})
}

func another(h *server.Hertz) {
	h.POST("v:path/bind", func(ctx context.Context, c *app.RequestContext) {
		var arg Args
		err := c.BindAndValidate(&arg)
		if err != nil {
			panic(err)
		}
		fmt.Println(arg)
	})
}
