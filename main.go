package main

import (
	"fmt"
	"net/http"

	"github.com/jchou8/sling/handlers"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func main() {
	//TODO: Set up DB connection
	var conn *gorm.DB

	srv := echo.New()
	srv.File("/", "public/index.html")
	srv.POST("api/register", handlers.NewUserHandler(conn))
	srv.POST("api/login", handlers.LoginHandler(conn))
	srv.GET("api/users/current", handlers.GetCurrentUserHandler(conn))
	srv.GET("api/users", handlers.GetUsersHandler(conn))

	fmt.Println("Listening at localhost:8000...")
	http.ListenAndServe(":8000", nil)
}
