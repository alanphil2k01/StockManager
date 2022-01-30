package config

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/alanphil2k01/SSMC/pkg/routes"
	"github.com/alanphil2k01/SSMC/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	// "gorm.io/driver/mysql"
)

var (
	//go:embed static
	content embed.FS
	address string
	port    string
)

func GetAddress() string {
	address = ":" + utils.GetEnv("SERVER_PORT", "8080")
	return address
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

func GetServer(_ ...func(http.Handler) http.Handler) *http.Server {
	router := mux.NewRouter()
	fsys, _ := fs.Sub(content, "static")
	routes.RegisterRoutes(router)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.FS(fsys))))

	return &http.Server{
		Handler:           router,
		Addr:              address,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}
}
