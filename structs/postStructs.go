package structs

type PostItem struct {
	Item_name string
	Quantity  int
}

type PostSender struct {
	Sender_first_name string
	Sender_last_name  string
	Phone_number      string
	Email             string
	Address           string
	City_name         string
	Post_code         string
}

type PostReceiver struct {
	Receiver_first_name string
	Receiver_last_name  string
	PassportNumber      string
	Phone_number        string
	Address             string
	City_name           string
}

type PostOrder struct {
	Comment        string
	Sign_signature string
	Item           PostItem
	Sender         PostSender
	Receiver       PostReceiver
}

type PostLogin struct {
	Username string
	Password string
}

type PatchAcceptOrder struct {
	OrderId int
	Status  bool
}

type PatchOrder struct {
	OrderId int
	Message string
}
