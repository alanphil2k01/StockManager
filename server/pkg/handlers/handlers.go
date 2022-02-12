package handlers

import (
	"net/http"

	"github.com/alanphil2k01/SSMC/pkg/db"
)

func RemoveExpired(w http.ResponseWriter, r *http.Request) {
	err := db.RemoveExpired()
	if err != nil {
		responsMessage(w, r, "Error - failed to remove expired", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Removed expired products", http.StatusOK, nil)
}
