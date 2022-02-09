package handlers

import (
	"log"
	"net/http"

	"github.com/alanphil2k01/SSMC/pkg/db"
	"github.com/alanphil2k01/SSMC/pkg/types"
	"github.com/alanphil2k01/SSMC/pkg/utils"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("/static -> ", r.Host)
	var product types.Product
	utils.ParseBody(r, &product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Got Product"))
	if err != nil {
		log.Println("Failed to write response")
	}
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	var product types.Product
	utils.ParseBody(r, &product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Got Product"))
	if err != nil {
		log.Println("Failed to write response")
	}
}

func RemoveExpired(w http.ResponseWriter, _ *http.Request) {
	err := db.RemoveExpired()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		log.Println("Failed to remove expired")
		_, err = w.Write([]byte("failed to remove expired products"))
		if err != nil {
			log.Println("Failed to write response")
		}
		return
	}
	_, err = w.Write([]byte("Successfully removed expired products\n"))
	if err != nil {
		log.Println("Successfully removed expired")
	}
}
