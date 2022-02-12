-- Create tables
CREATE TABLE products (
    prod_id         VARCHAR(10)     NOT NULL,
    prod_name       VARCHAR(30)     NOT NULL,
    rate            INT,
    total_qty       INT             DEFAULT(0),
    max_capacity    INT             DEFAULT(100),
    cat_id          INT             NOT NULL,
    supplier_id     INT             NOT NULL,
    PRIMARY KEY     (prod_id)
);

CREATE TABLE stocks (
    stock_id        VARCHAR(30)     NOT NULL,
    expiry_date     DATE,
    curr_qty        INT             DEFAULT(0),
    prod_id         VARCHAR(30)     NOT NULL,
    PRIMARY KEY    (stock_id)
);

CREATE TABLE stock_logs (
    log_id          INT             NOT NULL AUTO_INCREMENT,
    stock_id        VARCHAR(30)     NOT NULL,
    prod_id         VARCHAR(30)     NOT NULL,
    qty             INT             NOT NULL,
    date_processed  DATE            NOT NULL,
    expiry_date     DATE,
    action          VARCHAR(30)     NOT NULL,
    status          VARCHAR(30)     NOT NULL,
    PRIMARY KEY     (log_id)
);

CREATE TABLE suppliers (
    supplier_id     INT             NOT NULL,
    s_name          VARCHAR(30)     NOT NULL,
    s_email         VARCHAR(30)     NOT NULL,
    manager         VARCHAR(30)     NOT NULL,
    address         VARCHAR(100)    NOT NULL,
    phone_no        VARCHAR(15)     NOT NULL,
    PRIMARY KEY     (supplier_id)
);

CREATE TABLE product_categories (
    cat_id          INT             NOT NULL,
    cat_name        VARCHAR(30)     NOT NULL,
    warehouse_loc   VARCHAR(30)     DEFAULT('UNASSIGNED'),
    PRIMARY KEY     (cat_id)
);

CREATE TABLE users (
    username        VARCHAR(30),
    password        VARCHAR(30),
    email           VARCHAR(30),
    name            VARCHAR(30),
    role            CHAR             DEFAULT('U'),
    PRIMARY KEY     (username)
);

-- Add required foreign keys
ALTER TABLE products ADD FOREIGN KEY (supplier_id) REFERENCES suppliers(supplier_id);
ALTER TABLE products ADD FOREIGN KEY (cat_id) REFERENCES product_categories(cat_id);
ALTER TABLE stocks ADD FOREIGN KEY (prod_id) REFERENCES products(prod_id);
