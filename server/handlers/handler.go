package handlers

import (
	"dashboard-api/service/user"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

type ServerConf struct {
	Address       string
}

type HTTPServer struct {
	ServerConf ServerConf
}

func (h * HTTPServer)  Run() {
	e := echo.New()
	e.Use(mw.Recover())
	e.Use(mw.Logger())
	e.Use(mw.BodyLimit("2M"))

	userG := &user.Getter{}
	userH := UserHandler{
		get: userG.Run,
	}
	e.GET("user/:userId", userH.Get)
	
	e.Logger.Fatal(e.Start(h.ServerConf.Address))
}