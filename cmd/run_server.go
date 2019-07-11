package cmd

import (
	"context"
	"fmt"
	"github.com/calvinfeng/sling/handler"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	// Postgres database driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// RunServerCmd is the command used to run the server.
var RunServerCmd = &cobra.Command{
	Use:   "runserver",
	Short: "run user authentication server",
	RunE:  runServer,
}

func runServer(cmd *cobra.Command, args []string) error {
	conn, err := gorm.Open("postgres", pgAddr)
	if err != nil {
		log.Fatalf("failed to open DB conn: %s", err.Error())
	}

	setupBroker(conn)

	srv := echo.New()

	srv.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "HTTP[${time_rfc3339}] ${method} ${path} status=${status} latency=${latency_human}\n",
		Output: io.MultiWriter(os.Stdout),
	}))

	srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// srv.File("/", "public/index.html")
	// srv.Static("/assets", "public/assets")

	srv.File("/", "frontend/build/index.html")
	srv.Static("/static", "frontend/build/static")

	srv.POST("/api/register", handler.NewUserHandler(conn))
	srv.POST("/api/login", handler.LoginHandler(conn))

	users := srv.Group("api/users")
	users.Use(handler.NewTokenAuthMiddleware(conn))
	users.GET("/", handler.GetUsersHandler(conn))
	users.GET("/current", handler.GetCurrentUserHandler(conn))

	srv.GET("/api/rooms", handler.GetRoomsHandler(conn), handler.NewTokenAuthMiddleware(conn))

	messageStreamHandler := handler.GetMessageStreamHandler(&websocket.Upgrader{})
	actionStreamHandler := handler.GetActionStreamHandler(&websocket.Upgrader{})

	streams := srv.Group("api/stream")

	streams.GET("/messages", messageStreamHandler)
	streams.GET("/actions", actionStreamHandler)

	fmt.Println("Listening at localhost:8888...")
	if err := srv.Start(":8888"); err != nil {
		return err
	}

	return nil
}

func setupBroker(conn *gorm.DB) {
	handler.RunBroker(context.Background(), conn)
}
