package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jchou8/sling/handlers"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host         = "localhost"
	port         = "5432"
	user         = "jcho"
	password     = "jcho"
	database     = "sling"
	ssl          = "sslmode=disable"
	migrationDir = "file://./migrations/"
)

var pgAddr = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", user, password, host, port, database, ssl)

func main() {
	conn, err := gorm.Open("postgres", pgAddr)
	if err != nil {
		log.Fatalf("failed to open DB conn: %s", err.Error())
	}

	srv := echo.New()

	srv.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "HTTP[${time_rfc3339}] ${method} ${path} status=${status} latency=${latency_human}\n",
		Output: io.MultiWriter(os.Stdout),
	}))

	srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	srv.File("/", "public/index.html")
	srv.POST("/api/register", handlers.NewUserHandler(conn))
	srv.POST("/api/login", handlers.LoginHandler(conn))

	users := srv.Group("api/users")
	users.Use(handlers.NewTokenAuthMiddleware(conn))
	users.GET("/", handlers.GetUsersHandler(conn))
	users.GET("/current", handlers.GetCurrentUserHandler(conn))

	fmt.Println("Listening at localhost:8000...")
	if err := srv.Start(":8000"); err != nil {
		log.Fatalf(err.Error())
	}
}
