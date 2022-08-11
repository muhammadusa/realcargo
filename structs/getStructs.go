package structs

type Item struct {
	Item_name string
	Quantity  int
}

type Sender struct {
	Sender_first_name string
	Sender_last_name  string
	Phone_number      string
	Email             string
	Address           string
	City_name         string
	Post_code         string
}

type Receiver struct {
	Receiver_first_name string
	Receiver_last_name  string
	Passport_number     string
	Phone_number        string
	Address             string
	City_name           string
}

type Order struct {
	OrderId        int
	Comment        string
	Sign_signature string
	Date           string
	Item           Item
	Sender         Sender
	Receiver       Receiver
}
