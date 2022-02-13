package db

import (
	"log"

	"github.com/alanphil2k01/SSMC/pkg/types"
)

func GetProductCategories() ([]types.ProductCategories, error) {
	var categories []types.ProductCategories
	var c types.ProductCategories
	resultSet, err := db.Query("SELECT cat_id, cat_name, warehouse_loc FROM product_categories")
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&c.Cat_id,
			&c.Cat_name,
			&c.Warehouse_loc)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func GetProductCategoryById(cat_id int) (types.ProductCategories, error) {
	var c types.ProductCategories
	row := db.QueryRow("SELECT cat_id, cat_name, warehouse_loc FROM product_categories WHERE cat_id = ?", cat_id)
	err := row.Scan(&c.Cat_id,
		&c.Cat_name,
		&c.Warehouse_loc)
	if err != nil {
		log.Println(err)
		return types.ProductCategories{}, err
	}
	return c, nil
}

func GetProductCategoriesByName(cat_name string) ([]types.ProductCategories, error) {
	var categories []types.ProductCategories
	var c types.ProductCategories
	resultSet, err := db.Query("SELECT cat_id, cat_name, warehouse_loc FROM product_categories WHERE cat_name LIKE ?", "%"+cat_name+"%")
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&c.Cat_id,
			&c.Cat_name,
			&c.Warehouse_loc)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func PutProductCategory(c types.ProductCategories) error {
	stmt, err := db.Prepare("INSERT INTO product_categories(cat_name, warehouse_loc) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&c.Cat_name,
		&c.Warehouse_loc)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProductCategory(cat_id int, c types.ProductCategories) error {
	var category types.ProductCategories
	row := db.QueryRow("SELECT cat_name, warehouse_loc FROM product_categories WHERE cat_id = ?", cat_id)
	err := row.Scan(&category.Cat_name,
		&category.Warehouse_loc)
	if err != nil {
		return err
	}
	if c.Cat_name == "" {
		c.Cat_name = category.Cat_name
	}
	if c.Warehouse_loc == "" {
		c.Warehouse_loc = category.Warehouse_loc
	}
	stmt, err := db.Prepare("UPDATE product_categories SET cat_name=?, warehouse_loc=? WHERE cat_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.Cat_name, c.Warehouse_loc, cat_id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProductCategory(cat_id int) error {
	stmt, err := db.Prepare("DELETE FROM product_categories WHERE cat_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(cat_id)
	if err != nil {
		return err
	}
	return nil
}
