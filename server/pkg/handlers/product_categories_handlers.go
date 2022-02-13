package handlers

import (
	"net/http"
	"strconv"

	"github.com/alanphil2k01/SSMC/pkg/db"
	"github.com/alanphil2k01/SSMC/pkg/types"
	"github.com/alanphil2k01/SSMC/pkg/utils"
	"github.com/gorilla/mux"
)

func GetProductCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := db.GetProductCategories()
	if err != nil {
		responsMessage(w, r, "Error - cannot get product categories", http.StatusInternalServerError, err)
		return
	}
	if categories == nil {
		responsMessage(w, r, "Empty Result Set", http.StatusOK, err)
		return
	}
	responsMessage(w, r, "Got product categories", http.StatusOK, categories)
}

func GetProductCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cat_id, _ := strconv.Atoi(vars["cat_id"])
	category, err := db.GetProductCategoryById(cat_id)
	if err != nil {
		responsMessage(w, r, "Error - cannot get product category with id "+strconv.Itoa(cat_id), http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Got product category by id", http.StatusOK, category)
}

func GetProductCategoryByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cat_name := vars["cat_name"]
	categories, err := db.GetProductCategoriesByName(cat_name)
	if err != nil {
		responsMessage(w, r, "Error -  cannot get product categories by name", http.StatusInternalServerError, err)
		return
	}
	if categories == nil {
		responsMessage(w, r, "Empty Result Set", http.StatusOK, nil)
		return
	}
	responsMessage(w, r, "Got product categories by name", http.StatusOK, categories)
}

func PutProductCategory(w http.ResponseWriter, r *http.Request) {
	var category types.ProductCategories
	utils.ParseBody(r, &category)
	if category.Cat_name == "" {
		responsMessage(w, r, "Error - invalid input json", http.StatusBadRequest, nil)
		return
	}
	if category.Warehouse_loc == "" {
		category.Warehouse_loc = "UNASSIGNED"
	}
	err := db.PutProductCategory(category)
	if err != nil {
		responsMessage(w, r, "Error - inserting product category", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Inserted product category", http.StatusOK, nil)
}

func UpdateProductCategory(w http.ResponseWriter, r *http.Request) {
	var category types.ProductCategories
	vars := mux.Vars(r)
	cat_id, _ := strconv.Atoi(vars["cat_id"])
	utils.ParseBody(r, &category)
	err := db.UpdateProductCategory(cat_id, category)
	if err != nil {
		responsMessage(w, r, "Error - updating product category", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Updated product category", http.StatusCreated, nil)
}

func DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cat_id, _ := strconv.Atoi(vars["cat_id"])
	err := db.DeleteProductCategory(cat_id)
	if err != nil {
		responsMessage(w, r, "Error - cannot delete product category with id "+strconv.Itoa(cat_id), http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Deleted product category", http.StatusOK, nil)
}
