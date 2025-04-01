package gormrepo_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/iotassss/saitamarental/internal/repository/gormrepo"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	_ = godotenv.Load("../../../.env")
}

func setupTestDB(t *testing.T) *gorm.DB {
	dbName := mustGetEnv("MYSQL_DATABASE") + "_test"
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mustGetEnv("MYSQL_ROOT_PASSWORD"),
		mustGetEnv("DB_HOST"),
		mustGetEnv("DB_PORT"),
		dbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		// If the connection fails, try to connect to the local MySQL server
		dbHost := "127.0.0.1"
		dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mustGetEnv("MYSQL_ROOT_PASSWORD"),
			dbHost,
			mustGetEnv("DB_PORT"),
			dbName,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	require.NoError(t, err)
	require.NoError(t, db.Exec("DROP TABLE IF EXISTS cities").Error)
	require.NoError(t, db.Exec("DROP TABLE IF EXISTS average_rents").Error)
	require.NoError(t, db.AutoMigrate(
		&gormrepo.CityModel{},
		&gormrepo.AverageRentModel{},
	))
	return db
}

func mustGetEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	panic(fmt.Sprintf("Missing required environment variable: %s", key))
}
