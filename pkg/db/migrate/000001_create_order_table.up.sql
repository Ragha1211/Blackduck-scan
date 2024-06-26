CREATE TABLE IF NOT EXISTS orders (
   order_id serial PRIMARY KEY,
   product_name VARCHAR (50)  NOT NULL,
   order_type VARCHAR (50)  NOT NULL,
   order_price integer NOT NULL,
   order_quantity integer  NOT NULL
);