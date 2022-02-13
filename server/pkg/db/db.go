package db

import (
	"database/sql"
	"log"

	"github.com/alanphil2k01/SSMC/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = config.GetDB()
	if err != nil {
		log.Println("Cannot connect to mysql: ", err)
	} else {
		log.Println("Successfully connected to mysql")
	}
}

func Close() {
	db.Close()
}

func RemoveExpired() error {
	remExpiredProc, err := db.Prepare("call remove_expired()")
	if err != nil {
		return err
	}
	defer remExpiredProc.Close()
	_, err = remExpiredProc.Exec()
	if err != nil {
		return err
	}
	log.Println("Remove expired stocks")
	return nil
}

func GetSupplier() {
}

func InsertSupplier() {
}
