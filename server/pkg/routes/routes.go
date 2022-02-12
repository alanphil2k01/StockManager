package routes

import (
	"github.com/alanphil2k01/SSMC/pkg/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/product", handlers.GetProduct).Methods("GET")
	router.HandleFunc("/product/{prod_id}", handlers.GetProductById).Methods("GET")
	router.HandleFunc("/product", handlers.PutProduct).Methods("POST")
	router.HandleFunc("/product/search/{prod_name}", handlers.GetProductByName).Methods("GET")

	// router.HandleFunc("/product", handlers.GetProduct).Methods("GET")
	// router.HandleFunc("/product/{id}", handlers.GetProduct).Methods("GET")
	// router.HandleFunc("/product", handlers.PutProduct).Methods("POST")
	//
	// router.HandleFunc("/product", handlers.GetProduct).Methods("GET")
	// router.HandleFunc("/product/{id}", handlers.GetProduct).Methods("GET")
	// router.HandleFunc("/product", handlers.PutProduct).Methods("POST")

	router.HandleFunc("/remove_expired", handlers.RemoveExpired).Methods("GET")
}
