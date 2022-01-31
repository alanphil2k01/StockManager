CREATE TABLE products (
    prod_id         VARCHAR(20)     NOT NULL,
    rate            INT,
    total_qty       INT             DEFAULT(0),
    max_capacity    INT             DEFAULT(500),
    prod_category   INT             NOT NULL,
    supplier_id     INT             NOT NULL,
    PRIMARY KEY     (prod_id)
);

CREATE TABLE stocks (
    stock_id        VARCHAR(20)     NOT NULL,
    expiy_date      DATE,
    curr_qty        INT             DEFAULT(0),
    prod_id         VARCHAR(20)     NOT NULL,
    PRIMARY KEY    (stock_id)
);

CREATE TABLE stock_logs (
    log_id          INT             NOT NULL AUTO_INCREMENT,
    stock_id        VARCHAR(20)     NOT NULL,
    prod_id         VARCHAR(20)     NOT NULL,
    qty             INT             NOT NULL,
    date_arrived    DATE            NOT NULL,
    PRIMARY KEY     (log_id)
);

CREATE TABLE suppliers (
    supplier_id     INT             NOT NULL,
    s_name          VARCHAR(20)     NOT NULL,
    s_email         VARCHAR(20)     NOT NULL,
    manager         VARCHAR(20)     NOT NULL,
    address         VARCHAR(100)    NOT NULL,
    phone_no        VARCHAR(10)     NOT NULL,
    PRIMARY KEY     (supplier_id)
);

CREATE TABLE product_categories (
    cat_id          INT             NOT NULL,
    cat_name        VARCHAR(20)     NOT NULL,
    warehouse_loc   VARCHAR(20)     DEFAULT('UNASSIGNED'),
    PRIMARY KEY     (cat_id)
);

CREATE TABLE users (
    username        VARCHAR(20),
    password        VARCHAR(30),
    email           VARCHAR(30),
    name            VARCHAR(30),
    role            CHAR             DEFAULT('U'),
    PRIMARY KEY     (username)
);
