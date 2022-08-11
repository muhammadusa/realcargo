package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	c "app/controllers"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/realcargo", c.GetCargosCtrl).Methods("GET")
	r.HandleFunc("/realcargo", c.PostCargosCtrl).Methods("POST")

	r.HandleFunc("/orders/{order_id}", c.GetOrderIdCtrl).Methods("GET")

	r.HandleFunc("/login", c.Login).Methods("POST")

	r.HandleFunc("/accept", c.PatchAcceptOrderCtrl).Methods("POST")
	// r.HandleFunc("/acceptedOrders", c.GetAcceptedOrdersCtrl).Methods("GET")

	log.Println("Listening on the localhost: 8080")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
