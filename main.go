package main

import (
	"example.com/test1/db"
	"example.com/test1/handler"
	"example.com/test1/repository/repo_impl"
	"example.com/test1/router"
	"github.com/labstack/echo/v4"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "hoangthaifc01",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close()
	e := echo.New()
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()
	e.Logger.Fatal(e.Start(":3000"))
}
