package db

import (
	"database/sql"
	"log"

	"github.com/alanphil2k01/SSMC/pkg/config"
	"github.com/alanphil2k01/SSMC/pkg/types"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db          *sql.DB
	getProd     *sql.Stmt
	getProdById *sql.Stmt
	insProduct  *sql.Stmt
)

func init() {
	var err error
	db, err = config.GetDB()
	if err != nil {
		log.Println("Cannot connect to mysql: ", err)
	} else {
		log.Println("Successfully connected to mysql")
	}
	getProd, err := db.Prepare("SELECT * FROM Products WHERE number = ?")
	if err != nil {
		log.Fatal("Error creating statemennt:", err)
	}
	getProdById, err = db.Prepare("INSERT INTO Products VALUES( ?, ? )")
	if err != nil {
		log.Fatal("Error creating statemennt:", err)
	}
	insProduct, err = db.Prepare("INSERT INTO Products VALUES( ?, ? )")
	if err != nil {
		log.Fatal("Error creating statemennt:", err)
	}

}

func Close() {
	db.Close()
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
