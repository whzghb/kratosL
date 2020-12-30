package http

import (
	"fmt"
	"kratosL/app04/internal/service"
	"kratosL/middleWare"
	"net/http"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	pb "kratosL/app04/api"
	"kratosL/app04/internal/model"
)

var svc service.Service
//var svc pb.DemoServer

// New new a bm server.
func New(s pb.DemoServer) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	s = &svc
	engine = bm.DefaultServer(&cfg)
	pb.RegisterDemoBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	//全局使用
	//e.UseFunc(middleWare.AuthMiddleWare)
	g := e.Group("/app04")
	{
		g.GET("/start", howToStart)
	}
	l := e.Group("/kratosL", middleWare.AuthMiddleWare)
	//l.UseFunc(middleWare.AuthMiddleWare)
	{
		l.POST("/login", Login)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svc.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}

func Login(c *bm.Context)  {
	fmt.Println("YOU ARE LOGIN !!")
	c.JSON("ok", nil)
}