package db

import (
	"database/sql"
	"log"

	"github.com/alanphil2k01/SSMC/pkg/config"
	"github.com/alanphil2k01/SSMC/pkg/types"
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

func GetProducts() ([]types.ProductsList, error) {
	var products []types.ProductsList
	var product types.ProductsList
	resultSet, err := db.Query("SELECT prod_id, prod_name, rate, total_qty, s_name, cat_name FROM products_list")
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&product.Prod_id, &product.Prod_name, &product.Rate, &product.Total_qty, &product.S_name, &product.Cat_name)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetProductById(prod_id string) (types.ProductsList, error) {
	var product types.ProductsList
	row := db.QueryRow("SELECT prod_id, prod_name, rate, total_qty, s_name, cat_name FROM products_list WHERE prod_id = ?", prod_id)
	err := row.Scan(&product.Prod_id, &product.Prod_name, &product.Rate, &product.Total_qty, &product.S_name, &product.Cat_name)
	if err != nil {
		return types.ProductsList{}, err
	}
	return product, nil
}

func GetProductByName(prod_name string) ([]types.ProductsList, error) {
	var products []types.ProductsList
	var product types.ProductsList
	resultSet, err := db.Query("SELECT prod_id, prod_name, rate, total_qty, s_name, cat_name FROM products_list WHERE prod_name LIKE ?", "%"+prod_name+"%")
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&product.Prod_id, &product.Prod_name, &product.Rate, &product.Total_qty, &product.S_name, &product.Cat_name)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func PutProduct(p types.Products) error {
	stmt, err := db.Prepare("INSERT INTO products(prod_id, prod_name, rate, max_capacity, prod_category, supplier_id) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if p.Max_capacity == 0 {
		p.Max_capacity = 100
	}
	_, err = stmt.Exec(p.Prod_id, p.Prod_name, p.Rate, p.Max_capacity, p.Prod_category, p.Supplier_id)
	if err != nil {
		return err
	}
	return nil
}

func GetSupplier() {
}

func InsertSupplier() {
}
