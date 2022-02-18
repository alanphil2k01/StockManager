package handlers

import (
	"net/http"

	"github.com/alanphil2k01/SSMC/pkg/db"
	"github.com/alanphil2k01/SSMC/pkg/types"
	"github.com/alanphil2k01/SSMC/pkg/utils"
	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
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
		responsMessage(w, r, "Error - cannot get product with id "+prod_id, http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Got product by id", http.StatusOK, product)
}

func GetProductsByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prod_name := vars["prod_name"]
	products, err := db.GetProductsByName(prod_name)
	if err != nil {
		responsMessage(w, r, "Error -  cannot get product by name", http.StatusInternalServerError, err)
		return
	}
	if products == nil {
		responsMessage(w, r, "Empty Result Set", http.StatusOK, nil)
		return
	}
	responsMessage(w, r, "Got product by name", http.StatusOK, products)
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	var p types.Products
	utils.ParseBody(r, &p)
	if p.Prod_id == "" || p.Prod_name == "" || p.Rate == 0 || p.Cat_id == 0 || p.Supplier_id == 0 {
		responsMessage(w, r, "Error - invalid input json", http.StatusBadRequest, nil)
		return
	}
	if !utils.ValidateNameWithNumbers(p.Prod_name) || !utils.ValidateStrID(p.Prod_id) {
		responsMessage(w, r, "Error - invalid input format", http.StatusBadRequest, nil)
		return
	}
	err := db.PutProduct(p)
	if err != nil {
		responsMessage(w, r, "Error - inserting product", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Inserted product", http.StatusOK, nil)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product types.ProductsList
	vars := mux.Vars(r)
	prod_id := vars["prod_id"]
	utils.ParseBody(r, &product)
	err := db.UpdateProduct(prod_id, product)
	if err != nil {
		responsMessage(w, r, "Error - updating product", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Updated product", http.StatusCreated, nil)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prod_id := vars["prod_id"]
	err := db.DeleteProduct(prod_id)
	if err != nil {
		responsMessage(w, r, "Error - cannot delete product with id "+prod_id, http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Deleted product", http.StatusOK, nil)
}
