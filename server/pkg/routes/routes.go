package routes

import (
	"github.com/alanphil2k01/SSMC/pkg/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/product", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/product", handlers.PutProduct).Methods("POST")
	router.HandleFunc("/product/{prod_id}", handlers.GetProductById).Methods("GET")
	router.HandleFunc("/product/search/{prod_name}", handlers.GetProductsByName).Methods("GET")
	router.HandleFunc("/product/{prod_id}", handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{prod_id}", handlers.DeleteProduct).Methods("DELETE")

	router.HandleFunc("/supplier", handlers.GetSuppliers).Methods("GET")
	router.HandleFunc("/supplier", handlers.PutSupplier).Methods("POST")
	router.HandleFunc("/supplier/{supplier_id}", handlers.GetSupplierById).Methods("GET")
	router.HandleFunc("/supplier/search/{s_name}", handlers.GetSuppliersByName).Methods("GET")
	router.HandleFunc("/supplier/{supplier_id}", handlers.UpdateSupplier).Methods("PUT")
	router.HandleFunc("/supplier/{supplier_id}", handlers.DeleteSupplier).Methods("DELETE")

	router.HandleFunc("/product_category", handlers.GetProductCategories).Methods("GET")
	router.HandleFunc("/product_category", handlers.PutProductCategory).Methods("POST")
	router.HandleFunc("/product_category/{cat_id}", handlers.GetProductCategoryById).Methods("GET")
	router.HandleFunc("/product_category/search/{cat_name}", handlers.GetProductCategoryByName).Methods("GET")
	router.HandleFunc("/product_category/{cat_id}", handlers.UpdateProductCategory).Methods("PUT")
	router.HandleFunc("/product_category/{cat_id}", handlers.DeleteProductCategory).Methods("DELETE")

	router.HandleFunc("/stock", handlers.GetStocks).Methods("GET")
	router.HandleFunc("/stock", handlers.AddStock).Methods("POST")
	router.HandleFunc("/stock/{stock_id}", handlers.GetStockById).Methods("GET")
	router.HandleFunc("/stock/{prod_id}/{qty}", handlers.RemoveStocks).Methods("DELETE")

	router.HandleFunc("/stock_log", handlers.GetStockLogs).Methods("GET")
	router.HandleFunc("/stock_log/all", handlers.GetAllStockLogs).Methods("GET")
	router.HandleFunc("/stock_log/{num}", handlers.GetLastNStockLogs).Methods("GET")

	router.HandleFunc("/remove_expired", handlers.RemoveExpired).Methods("GET")
}
