package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo"

	"academy/internal/app"
	"academy/internal/handler"
)

func main() {
	state := flag.String("state", "dev", "program environment")
	flag.Parse()

	c := app.NewConfig(*state)
	if err := c.Init(); err != nil {
		fmt.Println("config init error", err)
		return
	}

	e := echo.New()
	if err := handler.NewRouter(e, c); err != nil {
		fmt.Println("new router error", err)
		return
	}
	e.Start(":8080")

	fmt.Println("hello world")
}
