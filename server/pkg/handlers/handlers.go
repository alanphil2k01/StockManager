package handlers

import (
	"log"
	"net/http"

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
