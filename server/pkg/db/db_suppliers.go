package db

import (
	"log"

	"github.com/alanphil2k01/SSMC/pkg/types"
)

func GetSuppliers() ([]types.Suppliers, error) {
	var suppliers []types.Suppliers
	var s types.Suppliers
	resultSet, err := db.Query("SELECT supplier_id, s_name, s_email, manager, address, phone_no FROM suppliers")
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&s.Supplier_id,
			&s.S_name,
			&s.S_email,
			&s.Manager,
			&s.Address,
			&s.Phone_no)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		suppliers = append(suppliers, s)
	}
	return suppliers, nil
}

func GetSupplierById(supplier_id int) (types.Suppliers, error) {
	var s types.Suppliers
	row := db.QueryRow("SELECT supplier_id, s_name, s_email, manager, address, phone_no FROM suppliers WHERE supplier_id = ?", supplier_id)
	err := row.Scan(&s.Supplier_id,
		&s.S_name,
		&s.S_email,
		&s.Manager,
		&s.Address,
		&s.Phone_no)
	if err != nil {
		log.Println(err)
		return types.Suppliers{}, err
	}
	return s, nil
}

func GetSuppliersByName(s_name string) ([]types.Suppliers, error) {
	var suppliers []types.Suppliers
	var s types.Suppliers
	resultSet, err := db.Query("SELECT supplier_id, s_name, s_email, manager, address, phone_no FROM suppliers WHERE s_name LIKE ?", "%"+s_name+"%")
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&s.Supplier_id,
			&s.S_name,
			&s.S_email,
			&s.Manager,
			&s.Address,
			&s.Phone_no)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, s)
	}
	return suppliers, nil
}

func PutSupplier(s types.Suppliers) error {
	stmt, err := db.Prepare("INSERT INTO suppliers(s_name, s_email, manager, address, phone_no) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&s.S_name,
		&s.S_email,
		&s.Manager,
		&s.Address,
		&s.Phone_no)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSupplier(supplier_id int, s types.Suppliers) error {
	var supplier types.Suppliers
	row := db.QueryRow("SELECT s_name, s_email, manager, address, phone_no FROM suppliers WHERE supplier_id = ?", supplier_id)
	err := row.Scan(&supplier.S_name,
		&supplier.S_email,
		&supplier.Manager,
		&supplier.Address,
		&supplier.Phone_no)
	if err != nil {
		return err
	}
	if s.S_name == "" {
		s.S_name = supplier.S_name
	}
	if s.S_email == "" {
		s.S_email = supplier.S_email
	}
	if s.Manager == "" {
		s.Manager = supplier.Manager
	}
	if s.Address == "" {
		s.Address = supplier.Address
	}
	if s.Phone_no == "" {
		s.Phone_no = supplier.Phone_no
	}
	stmt, err := db.Prepare("UPDATE suppliers SET s_name=?, s_email=?, manager=?, address=?, phone_no=? WHERE supplier_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.S_name, s.S_email, s.Manager, s.Address, s.Phone_no, supplier_id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSupplier(supplier_id int) error {
	stmt, err := db.Prepare("DELETE FROM suppliers WHERE supplier_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(supplier_id)
	if err != nil {
		return err
	}
	return nil
}
