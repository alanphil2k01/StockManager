package routes

import (
	"github.com/alanphil2k01/SSMC/pkg/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// mux.CORSMethodMiddleware(router)
	router.HandleFunc("/product", handlers.GetProduct).Methods("GET")
	router.HandleFunc("/remove_expired", handlers.RemoveExpired).Methods("GET")
	router.HandleFunc("/product", handlers.PutProduct).Methods("POST")
}
