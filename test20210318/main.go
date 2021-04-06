package main

import (
	"encoding/json"

	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/server/xecho"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/labstack/echo/v4"
)

func main() {
	eng := NewEngine()
	if err := eng.Run(); err != nil {
		xlog.Error(err.Error())
	}
}

type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.serveHTTP,
	); err != nil {
		xlog.Panic("startup", xlog.Any("err", err))
	}
	return eng
}

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

// HTTP地址
func (eng *Engine) serveHTTP() error {
	server := xecho.StdConfig("http").Build()
	user1 := &User{
		UserName: "user1",
		NickName: "上课看似",
		Age:      18,
		Birthday: "2008/8/8",
		Sex:      "男",
		Email:    "mahuateng@qq.com",
		Phone:    "110",
	}

	data, _ := json.Marshal(user1)

	server.GET("/hello", func(ctx echo.Context) error {
		return ctx.JSON(200, string(data))
	})
	return eng.Serve(server)
}
