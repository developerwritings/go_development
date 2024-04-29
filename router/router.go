// router/router.go
package router

import (
	"github.com/gorilla/mux"
	"github.com/yourusername/ecommerce-app/handler"
)

// NewRouter initializes a new router
func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	setupProductRoutes(r, db)
	setupUserRoutes(r, db)
	return r
}

func setupProductRoutes(r *mux.Router, db *sql.DB) {
	productHandler := handler.NewProductHandler(db)
	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
}

func setupUserRoutes(r *mux.Router, db *sql.DB) {
	userHandler := handler.NewUserHandler(db)
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
}
