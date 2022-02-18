package db

import (
	"errors"

	"github.com/alanphil2k01/SSMC/pkg/types"
	"github.com/alanphil2k01/SSMC/pkg/utils"
)

func GetProducts() ([]types.ProductsList, error) {
	var products []types.ProductsList
	var p types.ProductsList
	resultSet, err := db.Query("SELECT prod_id, prod_name, rate, total_qty, max_capacity, supplier_id, s_name, cat_id, cat_name FROM products_list")
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&p.Prod_id,
			&p.Prod_name,
			&p.Rate,
			&p.Total_qty,
			&p.Max_capacity,
			&p.Supplier_id,
			&p.S_name,
			&p.Cat_id,
			&p.Cat_name)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProductById(prod_id string) (types.ProductsList, error) {
	var p types.ProductsList
	row := db.QueryRow("SELECT prod_id, prod_name, rate, total_qty, max_capacity, supplier_id, s_name, cat_id, cat_name FROM products_list WHERE prod_id = ?", prod_id)
	err := row.Scan(&p.Prod_id,
		&p.Prod_name,
		&p.Rate,
		&p.Total_qty,
		&p.Max_capacity,
		&p.Supplier_id,
		&p.S_name,
		&p.Cat_id,
		&p.Cat_name)
	if err != nil {
		return types.ProductsList{}, err
	}
	return p, nil
}

func GetProductsByName(prod_name string) ([]types.ProductsList, error) {
	var products []types.ProductsList
	var p types.ProductsList
	resultSet, err := db.Query("SELECT prod_id, prod_name, rate, total_qty, max_capacity, supplier_id, s_name, cat_id, cat_name FROM products_list WHERE prod_name LIKE ?", "%"+prod_name+"%")
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&p.Prod_id,
			&p.Prod_name,
			&p.Rate,
			&p.Total_qty,
			&p.Max_capacity,
			&p.Supplier_id,
			&p.S_name,
			&p.Cat_id,
			&p.Cat_name)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func PutProduct(p types.Products) error {
	stmt, err := db.Prepare("INSERT INTO products(prod_id, prod_name, rate, max_capacity, cat_id, supplier_id) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if p.Max_capacity == 0 {
		p.Max_capacity = 100
	}
	_, err = stmt.Exec(p.Prod_id,
		p.Prod_name,
		p.Rate,
		p.Max_capacity,
		p.Cat_id,
		p.Supplier_id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(prod_id string, p types.ProductsList) error {
	var product types.ProductsList
	row := db.QueryRow("SELECT prod_id, prod_name, rate, max_capacity, supplier_id, s_name, cat_id, cat_name FROM products_list WHERE prod_id = ?", prod_id)
	err := row.Scan(&product.Prod_id,
		&product.Prod_name,
		&product.Rate,
		&product.Max_capacity,
		&product.Supplier_id,
		&product.S_name,
		&product.Cat_id,
		&product.Cat_name)
	if err != nil {
		return err
	}
	if p.Prod_name == "" || !utils.ValidateNameWithNumbers(p.Prod_name){
		p.Prod_name = product.Prod_name
	}
	if p.Rate == 0 {
		p.Rate = product.Rate
	}
	if p.Max_capacity == 0 {
		p.Max_capacity = product.Max_capacity
	}
	if p.Cat_id == 0 {
		p.Cat_id = product.Cat_id
	}
	if p.Supplier_id == 0 {
		p.Supplier_id = product.Supplier_id
	}
	stmt, err := db.Prepare("UPDATE products SET prod_name=?, rate=?, max_capacity=?, cat_id=?, supplier_id=? WHERE prod_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Prod_name, p.Rate, p.Max_capacity, p.Cat_id, p.Supplier_id, prod_id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(prod_id string) error {
	var count uint
	row := db.QueryRow("SELECT total_qty FROM products WHERE prod_id = ?", prod_id)
	err := row.Scan(&count)
    if count != 0 {
        return errors.New("product has some quantity left")
    }
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("DELETE FROM products WHERE prod_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(prod_id)
	if err != nil {
		return err
	}
	return nil
}
