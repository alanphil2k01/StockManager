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

func GetSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := db.GetSuppliers()
	if err != nil {
		responsMessage(w, r, "Error - cannot get suppliers", http.StatusInternalServerError, err)
		return
	}
	if suppliers == nil {
		responsMessage(w, r, "Empty Result Set", http.StatusOK, err)
		return
	}
	responsMessage(w, r, "Got suppliers", http.StatusOK, suppliers)
}

func GetSupplierById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	supplier_id, _ := strconv.Atoi(vars["supplier_id"])
	product, err := db.GetSupplierById(supplier_id)
	if err != nil {
		responsMessage(w, r, "Error - cannot get supplier with id "+strconv.Itoa(supplier_id), http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Got supplier by id", http.StatusOK, product)
}

func GetSuppliersByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s_name := vars["s_name"]
	suppliers, err := db.GetSuppliersByName(s_name)
	if err != nil {
		responsMessage(w, r, "Error -  cannot get supplier by name", http.StatusInternalServerError, err)
		return
	}
	if suppliers == nil {
		responsMessage(w, r, "Empty Result Set", http.StatusOK, nil)
		return
	}
	responsMessage(w, r, "Got supplier by name", http.StatusOK, suppliers)
}

func PutSupplier(w http.ResponseWriter, r *http.Request) {
	var s types.Suppliers
	utils.ParseBody(r, &s)
	if s.S_name == "" || s.S_email == "" || s.Manager == "" || s.Address == "" || s.Phone_no == "" {
		responsMessage(w, r, "Error - invalid input json", http.StatusBadRequest, nil)
		return
	}
	if !utils.ValidateName(s.S_name) || !utils.ValidateName(s.Manager) || !utils.ValidateEmail(s.S_email) || !utils.ValidatePhoneNo(s.Phone_no) || !utils.ValidateAddress(s.Address){
		responsMessage(w, r, "Error - invalid input format", http.StatusBadRequest, nil)
		return
	}
	err := db.PutSupplier(s)
	if err != nil {
		responsMessage(w, r, "Error - inserting suppliers", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Inserted supplier", http.StatusOK, nil)
}

func UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier types.Suppliers
	vars := mux.Vars(r)
	supplier_id, _ := strconv.Atoi(vars["supplier_id"])
	utils.ParseBody(r, &supplier)
	err := db.UpdateSupplier(supplier_id, supplier)
	if err != nil {
		responsMessage(w, r, "Error - updating supplier", http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Updated supplier", http.StatusCreated, nil)
}

func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	supplier_id, _ := strconv.Atoi(vars["supplier_id"])
	log.Println(supplier_id)
	err := db.DeleteSupplier(supplier_id)
	if err != nil {
		responsMessage(w, r, "Error - cannot delete supplier with id "+strconv.Itoa(supplier_id), http.StatusInternalServerError, err)
		return
	}
	responsMessage(w, r, "Deleted supplier", http.StatusOK, nil)
}
