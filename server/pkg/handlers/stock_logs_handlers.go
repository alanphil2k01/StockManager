package handlers

import (
	"net/http"
	"strconv"

	"github.com/alanphil2k01/SSMC/pkg/db"
	"github.com/gorilla/mux"
)

func GetLastNStockLogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num, _ := strconv.Atoi(vars["num"])
	logs, err := db.GetLastNStockLogs(num)
	if err != nil {
		responsMessage(w, r, "Error - cannot get stock logs", http.StatusInternalServerError, err)
		return
	}
	if logs == nil {
		responsMessage(w, r, "Empty Result Set", http.StatusOK, err)
		return
	}
	responsMessage(w, r, "Got stock logs", http.StatusOK, logs)
}

func GetAllStockLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := db.GetAllStockLogs()
	if err != nil {
		responsMessage(w, r, "Error - cannot get stock logs", http.StatusInternalServerError, err)
		return
	}
	if logs == nil {
		responsMessage(w, r, "Empty Result Set", http.StatusOK, err)
		return
	}
	responsMessage(w, r, "Got stock logs", http.StatusOK, logs)
}

func GetStockLogs(w http.ResponseWriter, r *http.Request) {
	num := 10
	logs, err := db.GetLastNStockLogs(num)
	if err != nil {
		responsMessage(w, r, "Error - cannot get stock logs", http.StatusInternalServerError, err)
		return
	}
	if logs == nil {
		responsMessage(w, r, "Empty Result Set", http.StatusOK, err)
		return
	}
	responsMessage(w, r, "Got stock logs", http.StatusOK, logs)
}
