package routes

import (
	h "github.com/alanphil2k01/SSMC/pkg/handlers"
	t "github.com/alanphil2k01/SSMC/pkg/types"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.RegisterUser).Methods("POST")
	router.HandleFunc("/login", h.LoginUser).Methods("POST")

	router.Handle("/product", h.CheckAuth(h.GetProducts, t.USER)).Methods("GET")
	router.Handle("/product", h.CheckAuth(h.PutProduct, t.ADIMINISTATOR)).Methods("POST")
	router.Handle("/product/{prod_id}", h.CheckAuth(h.GetProductById, t.USER)).Methods("GET")
	router.Handle("/product/search/{prod_name}", h.CheckAuth(h.GetProductsByName, t.USER)).Methods("GET")
	router.Handle("/product/{prod_id}", h.CheckAuth(h.UpdateProduct, t.ADIMINISTATOR)).Methods("PUT")
	router.Handle("/product/{prod_id}", h.CheckAuth(h.DeleteProduct, t.ADIMINISTATOR)).Methods("DELETE")

	router.Handle("/supplier", h.CheckAuth(h.GetSuppliers, t.STAFF)).Methods("GET")
	router.Handle("/supplier", h.CheckAuth(h.PutSupplier, t.ADIMINISTATOR)).Methods("POST")
	router.Handle("/supplier/{supplier_id}", h.CheckAuth(h.GetSupplierById, t.STAFF)).Methods("GET")
	router.Handle("/supplier/search/{s_name}", h.CheckAuth(h.GetSuppliersByName, t.STAFF)).Methods("GET")
	router.Handle("/supplier/{supplier_id}", h.CheckAuth(h.UpdateSupplier, t.ADIMINISTATOR)).Methods("PUT")
	router.Handle("/supplier/{supplier_id}", h.CheckAuth(h.DeleteSupplier, t.ADIMINISTATOR)).Methods("DELETE")

	router.Handle("/product_category", h.CheckAuth(h.GetProductCategories, t.USER)).Methods("GET")
	router.Handle("/product_category", h.CheckAuth(h.PutProductCategory, t.STAFF)).Methods("POST")
	router.Handle("/product_category/{cat_id}", h.CheckAuth(h.GetProductCategoryById, t.USER)).Methods("GET")
	router.Handle("/product_category/search/{cat_name}", h.CheckAuth(h.GetProductCategoryByName, t.USER)).Methods("GET")
	router.Handle("/product_category/{cat_id}", h.CheckAuth(h.UpdateProductCategory, t.STAFF)).Methods("PUT")
	router.Handle("/product_category/{cat_id}", h.CheckAuth(h.DeleteProductCategory, t.STAFF)).Methods("DELETE")

	router.Handle("/stock", h.CheckAuth(h.GetStocks, t.STAFF)).Methods("GET")
	router.Handle("/stock", h.CheckAuth(h.AddStock, t.STAFF)).Methods("POST")
	router.Handle("/stock/{stock_id}", h.CheckAuth(h.GetStockById, t.STAFF)).Methods("GET")
	router.Handle("/stock/{prod_id}/{qty}", h.CheckAuth(h.RemoveStocks, t.STAFF)).Methods("DELETE")

	router.Handle("/stock_log", h.CheckAuth(h.GetStockLogs, t.STAFF)).Methods("GET")
	router.Handle("/stock_log/all", h.CheckAuth(h.GetAllStockLogs, t.STAFF)).Methods("GET")
	router.Handle("/stock_log/{num}", h.CheckAuth(h.GetLastNStockLogs, t.STAFF)).Methods("GET")

	router.Handle("/remove_expired", h.CheckAuth(h.RemoveExpired, t.ADIMINISTATOR)).Methods("GET")
}
