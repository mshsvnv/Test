CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    name TEXT,
    surname TEXT,
    email TEXT unique,
    password TEXT,
    role TEXT
);

CREATE TABLE IF NOT EXISTS racket (
    id SERIAL PRIMARY KEY,
    brand TEXT,
    weight FLOAT,
    balance FLOAT,
    head_size FLOAT,
    avaliable BOOLEAN,
    price FLOAT,
    quantity INT,
    image BYTEA
);

CREATE TABLE IF NOT EXISTS cart (
    user_id INT unique,
    quantity INT,
    total_price FLOAT,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "order" (
    id SERIAL PRIMARY KEY,
    user_id INT,
    status TEXT,
    total_price FLOAT,
    creation_date TIMESTAMP,
    delivery_date TIMESTAMP,
    address TEXT,
    recepient_name TEXT,
    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_racket (
    order_id INT REFERENCES "order"(id) ON DELETE CASCADE,
    racket_id INT REFERENCES racket(id) ON DELETE CASCADE,
    PRIMARY KEY (order_id, racket_id),
    quantity INT
);

CREATE TABLE IF NOT EXISTS cart_racket (
    cart_id INT REFERENCES cart(user_id) ON DELETE CASCADE,
    racket_id INT REFERENCES racket(id) ON DELETE CASCADE,
    PRIMARY KEY (cart_id, racket_id),
    quantity INT
);


INSERT INTO "user"(name, surname, email, password, role) 
VALUES('admin', 'admin', 'stepaha78@gmail.com', '$2a$10$TvqHCkE9vzttMir6vS.quOGAA9W3o/K/z7G8.EvQU8hQjLjKiY1qS', 'admin', 'Admin');