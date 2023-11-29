CREATE TABLE IF NOT EXISTS orders
(
    order_uid varchar(255) primary key,
    track_number varchar(255),
    entry varchar(255),
--  delivery
    name varchar(255),
    phone varchar(127),
    zip varchar(255),
    city varchar(127),
    address varchar(255),
    region varchar(127),
    email varchar(255),
--  payment
    transaction varchar(255),
    request_id varchar(127),
    currency varchar(127),
    provider varchar(127),
    amount int,
    payment_dt int,
    bank varchar(127),
    delivery_cost int,
    goods_total int,
    custom_fee int,
--  items по order_uid
    locale varchar(31),
    internal_signature varchar(127),
    customer_id varchar(127),
    delivery_service varchar(127),
    shardkey varchar(127),
    sm_id int,
    date_created varchar(127),
    oof_shard varchar(127)
);

CREATE TABLE IF NOT EXISTS items
(
    id serial primary key,
    order_id varchar(255),
    chrt_id int,
    track_number varchar(255),
    price int,
    rid varchar(255),
    name varchar(255),
    sale int,
    size varchar(127),
    total_price int,
    nm_id int,
    brand varchar(255),
    status int,
    foreign key (order_id) references orders (order_uid)
);