CREATE TABLE shops (
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE products (
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE products_count (
    id serial not null unique,
    shop_id int references shops (id) on delete cascade not null,
    product_id int references products (id) on delete cascade not null,
    count int not null
);

INSERT INTO shops (id, name, description) VALUES
(1, 'Магазин 1', 'Магазин продуктов'),
(2, 'Магазин 2', 'Магазин iphones'),
(3, 'Магазин 3', 'Магазин техники');


INSERT INTO products (id, name, description) VALUES
(1, 'iphone 14 pro', 'Классный телефон'),
(2, 'Холодильник', 'Классный холодильник'),
(3, 'Арбуз', 'Классный арбуз');

INSERT INTO products_count (shop_id, product_id, count) VALUES
(1, 3, 10),
(2, 1, 8),
(3, 2, 5);

