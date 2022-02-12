package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alanphil2k01/SSMC/pkg/db"
	"github.com/alanphil2k01/SSMC/pkg/types"
	"github.com/alanphil2k01/SSMC/pkg/utils"
	"github.com/gorilla/mux"
)

func responsMessage(w http.ResponseWriter, r *http.Request, msg string, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(types.ReponeMsg{Msg: msg, Data: data})
	log.Println(msg, " ", r.RemoteAddr)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetProducts()
	if err != nil {
		responsMessage(w, r, "Error - cannot get products", http.StatusInternalServerError, err)
		return
	}
	if products == nil {
		responsMessage(w, r, "Empty Result Set", http.StatusOK, err)
		return
	}
	responsMessage(w, r, "Got products", http.StatusOK, products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prod_id := vars["prod_id"]
	product, err := db.GetProductById(prod_id)
	if err != nil {
		responsMessage(w, r, "Error - cannot get product with that id", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Got product by id", http.StatusOK, product)
}

func GetProductByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prod_name := vars["prod_name"]
	products, err := db.GetProductByName(prod_name)
	if err != nil {
		responsMessage(w, r, "Error -  cannot get product by name", http.StatusInternalServerError, err)
		return
	}
	if products == nil {
		responsMessage(w, r, "Empty Result Set", http.StatusOK, nil)
		return
	}
	responsMessage(w, r, "Got product by id", http.StatusOK, products)
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	var product types.Products
	utils.ParseBody(r, &product)
	if product.Prod_id == "" || product.Prod_name == "" || product.Rate == 0 || product.Prod_category == 0 || product.Supplier_id == 0 {
		responsMessage(w, r, "Error - invalid input json", http.StatusBadRequest, nil)
		return
	}
	err := db.PutProduct(product)
	if err != nil {
		responsMessage(w, r, "Error - inserting product", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Inserted product", http.StatusOK, nil)
}

func RemoveExpired(w http.ResponseWriter, r *http.Request) {
	err := db.RemoveExpired()
	if err != nil {
		responsMessage(w, r, "Error - failed to remove expired", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Removed expired products", http.StatusOK, nil)
}
