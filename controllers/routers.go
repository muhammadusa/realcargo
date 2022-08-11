package controllers

import (
	// "log"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"

	d "app/database"
	s "app/structs"
)

func GetCargosCtrl(w http.ResponseWriter, r *http.Request) {

	e := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	e.Encode(d.GetCargos())
}

func PostCargosCtrl(w http.ResponseWriter, r *http.Request) {

	e := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)

	var newCargo s.PostOrder

	json.Unmarshal(body, &newCargo)

	e.Encode(d.PostCargos(newCargo))
}

func GetOrderIdCtrl(w http.ResponseWriter, r *http.Request) {

	e := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")

	v := mux.Vars(r)

	e.Encode(d.GetOrderId(v["order_id"]))
}

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	e := json.NewEncoder(w)

	body, _ := ioutil.ReadAll(r.Body)

	var loginAdmin s.PostLogin

	json.Unmarshal(body, &loginAdmin)

	admin := d.Login(loginAdmin)

	if admin.Item_name == "" {

		e.Encode("User not found")

	} else {

		e.Encode(admin)
	}
}

func PatchAcceptOrderCtrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	e := json.NewEncoder(w)

	body, _ := ioutil.ReadAll(r.Body)

	var newCargo s.PatchAcceptOrder

	json.Unmarshal(body, &newCargo)

	newCargo.Status = true

	e.Encode(d.PatchAcceptOrders(newCargo))
}

// func GetAcceptedOrdersCtrl(w http.ResponseWriter, r *http.Request) {

// 	e := json.NewEncoder(w)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	e.Encode(d.GetOrders())
// }
