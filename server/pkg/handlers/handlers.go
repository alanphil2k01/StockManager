package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alanphil2k01/SSMC/pkg/db"
	"github.com/alanphil2k01/SSMC/pkg/types"
)

func responsMessage(w http.ResponseWriter, r *http.Request, msg string, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(types.ReponeMsg{Msg: msg, Data: data})
	log.Println(msg, " ", r.RemoteAddr)
}

func RemoveExpired(w http.ResponseWriter, r *http.Request) {
	err := db.RemoveExpired()
	if err != nil {
		responsMessage(w, r, "Error - failed to remove expired", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Removed expired products", http.StatusOK, nil)
}
