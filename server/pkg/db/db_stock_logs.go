package db

import (
	"github.com/alanphil2k01/SSMC/pkg/types"
)

func GetAllStockLogs() ([]types.StockLogs, error) {
	var logs []types.StockLogs
	var s types.StockLogs
	resultSet, err := db.Query("SELECT log_id, stock_id, prod_id, qty, date_processed, expiry_date, action, status FROM stock_logs ORDER BY log_id DESC")
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&s.Log_id,
			&s.Stock_id,
			&s.Prod_id,
			&s.Qty,
			&s.Date_processed,
			&s.Expiry_date,
			&s.Action,
			&s.Status)
		if err != nil {
			return nil, err
		}
		logs = append(logs, s)
	}
	return logs, nil
}

func GetLastNStockLogs(num int) ([]types.StockLogs, error) {
	var logs []types.StockLogs
	var log types.StockLogs
	resultSet, err := db.Query("SELECT log_id, stock_id, prod_id, qty, date_processed, expiry_date, action, status FROM stock_logs ORDER BY log_id DESC LIMIT ?", num)
	if err != nil {
		return nil, err
	}
	defer resultSet.Close()
	for resultSet.Next() {
		err = resultSet.Scan(&log.Log_id,
			&log.Stock_id,
			&log.Prod_id,
			&log.Qty,
			&log.Date_processed,
			&log.Expiry_date,
			&log.Action,
			&log.Status)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
