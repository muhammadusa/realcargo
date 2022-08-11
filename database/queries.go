package database

var SQL_GET_CARGO = `select
	s.sender_first_name,
	s.sender_last_name,
	s.phone_number,
	s.email,
	s.address,
	s.post_code,
	s.city_name,
	r.receiver_first_name,
	r.receiver_last_name,
	r.passport_number,
	r.phone_number,
	r.address,
	r.city_name,
	i.item_name,
	i.quantity,
	o.order_id,
	o.comment,
	o.sign_signature,
	o.created_at
from
	orders as o
join senders as s on s.sender_id = o.sender_id
join receivers as r on r.receiver_id = o.receiver_id
join items as i on i.item_id = o.item_id
`

var SQL_POST_SENDER = `
	insert into
		senders (sender_first_name, sender_last_name, phone_number, email, address, city_name, post_code)
	values
		($1, $2, $3, $4, $5, $6, $7)
	returning
		sender_first_name,
		sender_last_name,
		phone_number,
		email,
		address,
		city_name,
		post_code
`

var SQL_POST_RECEIVER = `
	insert into receivers (receiver_first_name, receiver_last_name, passport_number, phone_number, address, city_name) values ($1, $2, $3, $4, $5, $6)
	returning
		receiver_first_name,
		receiver_last_name,
		phone_number,
		address,
		city_name
`

var SQL_POST_ITEM = `
	insert into items (item_name, quantity) values ($1, $2)
	returning
		item_name,
		quantity
`

var SQL_POST_ORDER = `
	insert into
		orders (comment, sign_signature, item_id, sender_id, receiver_id)
	values ($1, $2, (select max(item_id) from items), (select max(sender_id) from senders), (select max(receiver_id) from receivers))
	returning
		order_id,
		comment,
		sign_signature,
		created_at	
`

var SQL_GetOrderId_CARGO = `select
    s.sender_first_name,
    s.sender_last_name,
    s.phone_number,
    s.email,
    s.address,
    s.city_name,
    s.post_code,
    r.receiver_first_name,
    r.receiver_last_name,
    r.passport_number,
    r.phone_number,
    r.address,
    r.city_name,
    i.item_name,
    i.quantity,
    o.order_id,
    o.comment,
    o.sign_signature,
    o.created_at
from
    orders as o
join senders as s on s.sender_id = o.sender_id
join receivers as r on r.receiver_id = o.receiver_id
join items as i on i.item_id = o.item_id
where o.order_id = $1
`

var SQL_LOGIN = `
	select 
		username, password
	from
		admins
	where
		username = $1 and password = crypt($2, password)
`

var SQL_ACCEPT = `
	update 
		orders 
	set 
		status = $1 
	where 
		order_id = $2
	returning
		order_id,
		comment,
		sign_signature,
		created_at	
`

// update orders set status = false where order_id = 1 returning order_id, comment, sign_signature, created_at;

// select order_id,  from orders as o where o.status = true;
