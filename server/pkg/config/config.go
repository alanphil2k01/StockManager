package config

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/alanphil2k01/SSMC/pkg/routes"
	"github.com/alanphil2k01/SSMC/pkg/utils"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func GetDB() *gorm.DB {
	database := utils.GetEnv("MYSQL_DATABASE", "db")
	user := utils.GetEnv("MYSQL_USER", "user")
	password := utils.GetEnv("MYSQL_PASSWORD", "password")
	connAddr := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", user, password, database)
	db, err := gorm.Open(mysql.Open(connAddr), &gorm.Config{})
	if err != nil {
		log.Println("Cannot connect to mysql ", err)
	} else {
		log.Println("Successfully userconnected to mysql")
	}
	return db
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
