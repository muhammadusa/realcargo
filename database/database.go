package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"

	s "app/structs"
)

func GetCargos() []s.Order {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	rows, err := db.Query(SQL_GET_CARGO)

	defer rows.Close()

	if err != nil {
		panic(err)
	}

	var cargos []s.Order

	for rows.Next() {

		var cargo s.Order

		rows.Scan(
			&cargo.Sender.Sender_first_name,
			&cargo.Sender.Sender_last_name,
			&cargo.Sender.Phone_number,
			&cargo.Sender.Email,
			&cargo.Sender.Address,
			&cargo.Sender.City_name,
			&cargo.Sender.Post_code,
			&cargo.Receiver.Receiver_first_name,
			&cargo.Receiver.Receiver_last_name,
			&cargo.Receiver.Passport_number,
			&cargo.Receiver.Phone_number,
			&cargo.Receiver.Address,
			&cargo.Receiver.City_name,
			&cargo.Item.Item_name,
			&cargo.Item.Quantity,
			&cargo.OrderId,
			&cargo.Comment,
			&cargo.Sign_signature,
			&cargo.Date,
		)

		cargos = append(cargos, cargo)
	}

	return cargos
}

func PostCargos(body s.PostOrder) (newCargo s.Order) {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	err = db.QueryRow(
		SQL_POST_SENDER,
		body.Sender.Sender_first_name,
		body.Sender.Sender_last_name,
		body.Sender.Phone_number,
		body.Sender.Email,
		body.Sender.Address,
		body.Sender.City_name,
		body.Sender.Post_code,
	).Scan(
		&newCargo.Sender.Sender_first_name,
		&newCargo.Sender.Sender_last_name,
		&newCargo.Sender.Phone_number,
		&newCargo.Sender.Email,
		&newCargo.Sender.Address,
		&newCargo.Sender.City_name,
		&newCargo.Sender.Post_code,
	)

	if err != nil {
		panic(err)
	}

	err = db.QueryRow(
		SQL_POST_RECEIVER,
		body.Receiver.Receiver_first_name,
		body.Receiver.Receiver_last_name,
		body.Receiver.PassportNumber,
		body.Receiver.Phone_number,
		body.Receiver.Address,
		body.Receiver.City_name,
	).Scan(
		&newCargo.Receiver.Receiver_first_name,
		&newCargo.Receiver.Receiver_last_name,
		&newCargo.Receiver.Phone_number,
		&newCargo.Receiver.Address,
		&newCargo.Receiver.City_name,
	)

	if err != nil {
		panic(err)
	}

	err = db.QueryRow(
		SQL_POST_ITEM,
		body.Item.Item_name,
		body.Item.Quantity,
	).Scan(
		&newCargo.Item.Item_name,
		&newCargo.Item.Quantity,
	)

	if err != nil {
		panic(err)
	}

	err = db.QueryRow(
		SQL_POST_ORDER,
		body.Comment,
		body.Sign_signature,
	).Scan(
		&newCargo.OrderId,
		&newCargo.Comment,
		&newCargo.Sign_signature,
		&newCargo.Date,
	)

	if err != nil {
		panic(err)
	}

	return newCargo
}

func GetOrderId(id string) []s.Order {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	vars, _ := strconv.Atoi(id)

	rows, err := db.Query(SQL_GetOrderId_CARGO, vars)

	defer rows.Close()

	if err != nil {
		panic(err)
	}

	var orders []s.Order

	for rows.Next() {

		var order s.Order

		rows.Scan(
			&order.Sender.Sender_first_name,
			&order.Sender.Sender_last_name,
			&order.Sender.Phone_number,
			&order.Sender.Email,
			&order.Sender.Address,
			&order.Sender.City_name,
			&order.Sender.Post_code,
			&order.Receiver.Receiver_first_name,
			&order.Receiver.Receiver_last_name,
			&order.Receiver.Passport_number,
			&order.Receiver.Phone_number,
			&order.Receiver.Address,
			&order.Receiver.City_name,
			&order.Item.Item_name,
			&order.Item.Quantity,
			&order.OrderId,
			&order.Comment,
			&order.Sign_signature,
			&order.Date,
		)

		orders = append(orders, order)
	}

	return orders
}

func Login(body s.PostLogin) (newCargo s.PostItem) {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	row, err := db.Query(
		SQL_LOGIN,
		body.Username,
		body.Password,
	)

	if row.Next() {

		row.Scan(
			&newCargo.Item_name,
			&newCargo.Quantity,
		)
	}

	if err != nil {
		panic(err)
	}

	return newCargo
}

func PatchAcceptOrder(id string) []s.Order {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	vars, _ := strconv.Atoi(id)

	rows, err := db.Query(SQL_ACCEPT, vars)

	defer rows.Close()

	if err != nil {
		panic(err)
	}

	var orders []s.Order

	for rows.Next() {

		var order s.Order

		rows.Scan(
			&order.Sender.Sender_first_name,
			&order.Sender.Sender_last_name,
			&order.Sender.Phone_number,
			&order.Sender.Email,
			&order.Sender.Address,
			&order.Sender.City_name,
			&order.Sender.Post_code,
			&order.Receiver.Receiver_first_name,
			&order.Receiver.Receiver_last_name,
			&order.Receiver.Passport_number,
			&order.Receiver.Phone_number,
			&order.Receiver.Address,
			&order.Receiver.City_name,
			&order.Item.Item_name,
			&order.Item.Quantity,
			&order.OrderId,
			&order.Comment,
			&order.Sign_signature,
			&order.Date,
		)

		orders = append(orders, order)
	}

	return orders
}

func PatchAcceptOrders(body s.PatchAcceptOrder) int {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	sqlStatement := `update orders set status = $1 where order_id = $2 returning order_id`

	var id int

	err = db.QueryRow(sqlStatement, body.Status, body.OrderId).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query = %v", err)
	}

	fmt.Printf("Successfully updated order_id = %v", id)

	return id
}
