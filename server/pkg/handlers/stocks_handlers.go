package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alanphil2k01/SSMC/pkg/db"
	"github.com/alanphil2k01/SSMC/pkg/types"
	"github.com/alanphil2k01/SSMC/pkg/utils"
	"github.com/gorilla/mux"
)

func GetStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := db.GetStocks()
	if err != nil {
		responsMessage(w, r, "Error - cannot get stocks", http.StatusInternalServerError, err)
		return
	}
	if stocks == nil {
		responsMessage(w, r, "Empty Result Set", http.StatusOK, err)
		return
	}
	responsMessage(w, r, "Got stocks", http.StatusOK, stocks)
}

func GetStockById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stock_id, _ := strconv.Atoi(vars["stock_id"])
	stock, err := db.GetStockById(stock_id)
	if err != nil {
		responsMessage(w, r, "Error - cannot get stock with id "+strconv.Itoa(stock_id), http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Got stock by id", http.StatusOK, stock)
}

func AddStock(w http.ResponseWriter, r *http.Request) {
	var stock types.Stocks
	utils.ParseBody(r, &stock)
	if stock.Stock_id == "" || stock.Prod_id == "" || stock.Expiry_date == "" || stock.Curr_qty == 0 {
		responsMessage(w, r, "Error - invalid input json", http.StatusBadRequest, nil)
		return
	}
	err := db.AddStock(stock)
	if err != nil {
		responsMessage(w, r, "Error - inserting suppliers", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Inserted supplier", http.StatusOK, nil)
}

func RemoveStocks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prod_id := vars["prod_id"]
	qty, _ := strconv.Atoi(vars["qty"])
	log.Println(prod_id)
	err := db.RemoveStocks(prod_id, qty)
	if err != nil {
		responsMessage(w, r, "Error - cannot remove "+strconv.Itoa(qty)+" stocks with product id "+prod_id, http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Removed "+strconv.Itoa(qty)+" of stocks of product id "+prod_id, http.StatusOK, nil)
}
