package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	apihander "github.com/iotassss/saitamarental/internal/handler/api"
	"github.com/iotassss/saitamarental/internal/usecase"
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
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	if mysqlDatabase == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "MYSQL_DATABASE"))
		return
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "MYSQL_USER"))
		return
	}
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "MYSQL_PASSWORD"))
		return
	}
	dblHost := os.Getenv("DB_HOST")
	if dblHost == "" {
		slog.Error("Missing required environment variables", slog.Any("error", "DB_HOST"))
		return
	}

	// database
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, dblHost, mysqlDatabase)
	if env == "development" {
		slog.Info("connecting to database", slog.Any("dsn", dbDSN))
	}
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	err = db.AutoMigrate(
	// &repository.User{},
	)
	if err != nil {
		slog.Error("failed to migrate database", slog.Any("error", err))
		return
	}

	// usecase
	chintaiSoubaHikakuInteractor := usecase.NewChintaiSoubaHikakuInteractor()

	// handler
	chintaiSoubaHikakuHandler := apihander.NewChintaiSoubaHikakuTokyoVsSaitamaHandler(db, chintaiSoubaHikakuInteractor)

	// router
	router := gin.Default()

	// routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	router.GET("/chintai/souba/hikaku/tokyo-vs-saitama", func(c *gin.Context) {
		chintaiSoubaHikakuHandler.Execute(c)
	})

	// start server
	if err := router.Run(fmt.Sprintf(":%s", appPort)); err != nil {
		slog.Error("failed to start server", slog.Any("error", err))
		return
	}
}
