package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/middleware"
	"waysbucks/pkg/mysql"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func ToppingRoutes(r *mux.Router) {
	toppingRepository := repositories.RepositoryTopping(mysql.DB)
	h := handlers.HandlerTopping(toppingRepository)

	r.HandleFunc("/toppings", h.FindToppings).Methods("GET")
	r.HandleFunc("/topping/{id}", h.GetTopping).Methods("GET")
	r.HandleFunc("/topping", middleware.Auth(middleware.UploadFile(h.AddTopping))).Methods("POST")
	r.HandleFunc("/topping/{id}", middleware.Auth(middleware.UploadFile(h.EditTopping))).Methods("PATCH")
	r.HandleFunc("/topping/{id}", middleware.Auth(middleware.UploadFile(h.DeleteTopping))).Methods("DELETE")
}
