INSERT INTO product_categories (cat_id, cat_name, warehouse_loc) VALUES
    (1, 'Stationary', 'A'),
    (2, 'Snacks', 'B'),
    (3, 'Drinks', 'C');

INSERT INTO suppliers(supplier_id, s_name, s_email, manager, address, phone_no) VALUES
    (1, 'ABC', 'abc@abc.com', 'A.B. C', 'abc lane', '8882223331'),
    (2, 'DEF', 'def@def.fom', 'D.E. F', 'def ldne', '8892443332');

INSERT INTO products(prod_id, prod_name, rate, max_capacity, prod_category, supplier_id) VALUES
    ('1', 'Pen1', 300, 200, 1, 1),
    ('2', 'Pen2', 200, 400, 1, 2),
    ('3', 'Biscuit1', 500, 500, 2, 1),
    ('4', 'Biscuit2', 600, 600, 2, 2),
    ('5', 'Biscuit3', 100, 200, 2, 2),
    ('6', 'Drink1', 250, 400, 3, 1),
    ('7', 'Drink2', 330, 200, 3, 1);

-- CREATE TABLE stocks (stock_id, expiry_date, curr_qty, prod_id) VALUES
--     ('1', '2020-02-01', 10, 1),
--     ('2', '2020-02-01', 20, 2),
--     ('3', '2022-03-01', 30, 3),
--     ('4', '2022-01-01', 10, 1),
--     ('5', '2023-02-01', 20, 3),
--     ('6', '2020-02-01', 30, 7),
--     ('7', '2022-02-01', 50, 2),
--     ('8', '2020-02-01', 70, 6),
--     ('9', '2020-02-01', 90, 4),
--     ('10', '2020-02-01', 80, 3);


call add_stock('1', '80', 80, '2021-03-03');
call add_stock('1', '2', 80, '2021-03-03');
call add_stock('1', '2', 80, '2023-03-03');
call add_stock('1', '2', 80, '2023-03-03');
call add_stock('1', '2', 80, '2023-03-03');
call add_stock('1', '2', 80, '2023-03-03');
call add_stock('1', '2', 80, '2023-03-03');
call add_stock('1', '2', 80, '2023-03-03');
call add_stock('1', '5', 300, '2023-03-03');
call add_stock('1', '5', 100, '2023-03-03');