package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// logger file
	loggerFile, err := os.OpenFile("log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("failed to open log file", slog.Any("error", err))
		return
	}
	defer loggerFile.Close()

	handler := slog.NewJSONHandler(loggerFile, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// load .env
	_ = godotenv.Load(".env")
	env := os.Getenv("APP_ENV")
	if env == "" {
		slog.Error(" environment variables", slog.Any("error", "APP_ENV"))
		return
	}
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "APP_PORT"))
		return
	}
	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "DB_DSN"))
		return
	}

	// database
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	err = db.AutoMigrate(
	// &repository.User{},
	)
	if err != nil {
		slog.Error("failed to migrate database", slog.Any("error", err))
		return
	}

	// usecase

	// handler

	// router
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")

	// routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Home",
		})
	})

	// start server
	if err := router.Run(fmt.Sprintf(":%s", appPort)); err != nil {
		slog.Error("failed to start server", slog.Any("error", err))
		return
	}
}
