package config

import (
	"database/sql"
	"fmt"

	"github.com/alanphil2k01/SSMC/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

func GetAddress() string {
	return ":" + utils.GetEnv("SERVER_PORT", "8080")
}

func GetDB() (*sql.DB, error) {
	database := utils.GetEnv("MYSQL_DATABASE", "db")
	user := utils.GetEnv("MYSQL_USER", "user")
	password := utils.GetEnv("MYSQL_PASSWORD", "password")
	connAddr := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", user, password, database)
	// db, err := gorm.Open(mysql.Open(connAddr), &gorm.Config{})
	db, err := sql.Open("mysql", connAddr)
	return db, err
}
