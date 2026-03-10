CREATE TABLE IF NOT EXISTS category (
    category_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    category_name VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS discount (
    discount_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    discount_rate INT,
    description VARCHAR(500),
    is_flash_sale BOOLEAN
);

CREATE TABLE IF NOT EXISTS products (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_name VARCHAR(30),
    product_desc TEXT,
    price INT,
    quantity INT,
    product_category INT,
    discount INT,
    CONSTRAINT fk_category FOREIGN KEY (product_category) REFERENCES category(category_id) ON DELETE SET NULL,
    CONSTRAINT fk_discount FOREIGN KEY (discount) REFERENCES discount(discount_id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS product_variant (
    product_variant_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_id INT,
    variant_name VARCHAR(50),
    add_price INT,
    CONSTRAINT fk_variant FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS product_size (
    product_size_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_id INT,
    size_name VARCHAR(50),
    size_price INT,
    CONSTRAINT fk_size FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS product_images (
    product_images_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_id INT,
    path VARCHAR(255),
    CONSTRAINT fk_images FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    full_name VARCHAR(30),
    email VARCHAR(30),
    password TEXT,
    address TEXT,
    phone VARCHAR(30),
    pictures VARCHAR(100)
);


CREATE TABLE IF NOT EXISTS reviews (
    review_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT,
    messages TEXT,
    ratings INT,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);


CREATE TABLE  IF NOT EXISTS"orders" (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    delivery_method VARCHAR(30),
    full_name VARCHAR(100),
    email VARCHAR(100),
    address TEXT,
    sub_total INT,
    delivery_fee INT,
    tax INT,
    total INT,
    date TIMESTAMP,
    status VARCHAR(30),
    payment_method INT
);


ALTER TABLE product_variant 
ADD CONSTRAINT unique_variant_per_product 
UNIQUE (product_id, variant_name);

ALTER TABLE product_size
ADD CONSTRAINT unique_size_per_product
UNIQUE (product_id, size_name);
