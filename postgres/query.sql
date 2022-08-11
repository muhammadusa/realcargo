{
    "Comment": "Iloji boricha tezroq yetib borsin.",
    "Sign_signature": "Javlon",
    "Item": {
        "Item_name": "furnitures",
        "Quantity": 2
    },
    "Sender": {
        "Sender_first_name": "Javlon",
        "Sender_last_name": "Jurabekov",
        "Phone_number": "+998946803533",
        "Email": "javlon1@gmail.com",
        "Address": "Salomatlik 2/4",
        "City_name": "Tashkent",
        "Post_code": "110105"
    },
    "Receiver": {
        "Receiver_first_name": "Alex",
        "Receiver_last_name": "John",
        "PassportNumber": "AA9995575",
        "Phone_number": "+861568627",
        "Address": "Wall Street 2",
        "City_name": "London"
    }
}

var SQL_GET_CARGO = `select
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
    o.status,
    o.sign_signature,
    o.created_at
from
    orders as o
join senders as s on s.sender_id = o.sender_id
join receivers as r on r.receiver_id = o.receiver_id
join items as i on i.item_id = o.item_id
where o.order_id = $1
`

select 
    o.order_id, 
    s.sender_first_name || ' ' || s.sender_last_name,
    r.receiver_first_name || ' ' || r.receiver_last_name 
from 
    orders as o
join senders as s on s.sender_id = o.sender_id
join receivers as r on r.receiver_id = o.receiver_id 
where 
    o.status = true;

    select 
        s.sender_first_name || ' ' || s.sender_last_name,
        r.receiver_first_name || ' ' || r.receiver_last_name
    from
        orders as o
    join senders as s on s.sender_id = o.sender_id
    join receivers as r on r.receiver_id = o.receiver_id

update orders set status = true where order_id = 1;