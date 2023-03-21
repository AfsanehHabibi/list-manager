package main

import (
	"github.com/AfsanehHabibi/list-manager/handler"
	"github.com/labstack/echo/v4"
)

func main()  {
	e := echo.New()
	h := handler.NewHandler("list")
	h.Register(e.Group(""))
	e.Logger.Fatal(e.Start(":1323"))
}