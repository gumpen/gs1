package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/gumpen/gs1/handler"
	"github.com/gumpen/gs1/infrastructure"
)

func main() {
	repositories := infrastructure.NewRepositories()
	handler := handler.NewHandler(repositories.UserRepository)

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/users", handler.GetUsers)

	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
