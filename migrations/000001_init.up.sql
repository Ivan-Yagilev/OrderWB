CREATE TABLE IF NOT EXISTS items
(
    id varchar(255),
    chrt_id int primary key,
    track_number varchar(255),
    price int,
    rid varchar(255),
    name varchar(255),
    sale int,
    size varchar,
    total_price int,
    nm_id int,
    brand varchar(255),
    status int
);

CREATE TABLE IF NOT EXISTS payment
(
    transaction varchar(255) primary key,
    request_id varchar(127),
    currency varchar(127),
    provider varchar(127),
    amount int,
    payment_dt int,
    bank varchar(127),
    delivery_cost int,
    goods_total int,
    custom_fee int
);

CREATE TABLE IF NOT EXISTS delivery
(
    id serial primary key,
    name varchar(255),
    phone varchar(127),
    zip varchar(255),
    city varchar(127),
    address varchar(255),
    region varchar(127),
    email varchar(255)
);

CREATE TABLE IF NOT EXISTS orders
(
    order_uid varchar(255) primary key,
    track_number varchar(255),
    entry varchar(255),
    delivery_id int,
    locale varchar(31),
    internal_signature varchar(127),
    customer_id varchar(127),
    delivery_service varchar(127),
    shardkey varchar,
    sm_id int,
    date_created varchar,
    oof_shard varchar,
    foreign key (order_uid) references payment (transaction),
    foreign key (delivery_id) references delivery (id)
);

-- items по order_uid