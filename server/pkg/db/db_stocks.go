package db

import (
	"github.com/alanphil2k01/SSMC/pkg/types"
)

func GetStocks() ([]types.StocksList, error) {
	var stocks []types.StocksList
	var s types.StocksList
	resultSet, err := db.Query("SELECT stock_id, expiry_date, curr_qty, prod_id, prod_name FROM stocks_list")
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&s.Stock_id,
			&s.Expiry_date,
			&s.Curr_qty,
			&s.Prod_id,
			&s.Prod_name)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, s)
	}
	return stocks, nil
}

func GetStockById(stock_id int) (types.StocksList, error) {
	var s types.StocksList
	row := db.QueryRow("SELECT stock_id, expiry_date, curr_qty, prod_id, prod_name FROM stocks_list WHERE stock_id = ?", stock_id)
	err := row.Scan(&s.Stock_id,
		&s.Expiry_date,
		&s.Curr_qty,
		&s.Prod_id,
		&s.Prod_name)
	if err != nil {
		return types.StocksList{}, err
	}
	return s, nil
}

func AddStock(s types.Stocks) error {
	stmt, err := db.Prepare("call add_stock(?,?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&s.Stock_id,
		&s.Prod_id,
		&s.Curr_qty,
		&s.Expiry_date)
	if err != nil {
		return err
	}
	return nil
}

func RemoveStocks(prod_id string, qty int) error {
	stmt, err := db.Prepare("call remove_stock(?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(prod_id, qty)
	if err != nil {
		return err
	}
	return nil
}
