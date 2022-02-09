package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alanphil2k01/SSMC/pkg/types"
	"github.com/alanphil2k01/SSMC/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db          *sql.DB
	getProd     *sql.Stmt
	getProdById *sql.Stmt
	insProduct  *sql.Stmt
)

func GetDB() (*sql.DB, error) {
	database := utils.GetEnv("MYSQL_DATABASE", "db")
	user := utils.GetEnv("MYSQL_USER", "user")
	password := utils.GetEnv("MYSQL_PASSWORD", "password")
	connAddr := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", user, password, database)
	// db, err := gorm.Open(mysql.Open(connAddr), &gorm.Config{})
	db, err := sql.Open("mysql", connAddr)
	return db, err
}
func init() {
	var err error
	db, err = GetDB()
	if err != nil {
		log.Println("Cannot connect to mysql: ", err)
	} else {
		log.Println("Successfully connected to mysql")
	}
	getProd, err = db.Prepare("SELECT * FROM Products WHERE number = ?")
	if err != nil {
		log.Println("Error creating statemennt:", err)
	}
	getProdById, err = db.Prepare("INSERT INTO Products VALUES( ?, ? )")
	if err != nil {
		log.Println("Error creating statemennt:", err)
	}
	insProduct, err = db.Prepare("INSERT INTO Products VALUES( ?, ? )")
	if err != nil {
		log.Println("Error creating statemennt:", err)
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

func GetProductById() {
}

func GetProductByName() {
}

func InsertProduct(product *types.Product) {

}

func GetSupplier() {
}

func InsertSupplier() {
}
