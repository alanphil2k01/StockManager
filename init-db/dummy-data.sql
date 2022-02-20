INSERT INTO product_categories (cat_name, warehouse_loc) VALUES
    ('Stationary', 'A'),
    ('Snacks', 'B'),
    ('Drinks', 'C');

INSERT INTO suppliers(s_name, s_email, manager, address, phone_no) VALUES
    ('ABC', 'abc@abc.com', 'A.B. C', 'abc lane', '8882223331'),
    ('DEF', 'def@def.fom', 'D.E. F', 'def ldne', '8892443332');

INSERT INTO products(prod_id, prod_name, rate, max_capacity, cat_id, supplier_id, total_qty) VALUES
    ('P1', 'Pen1', 300, 200, 1, 1, 0),
    ('P2', 'Pen2', 200, 400, 1, 2, 0),
    ('B1', 'Biscuit1', 500, 500, 2, 1, 80),
    ('B2', 'Biscuit2', 600, 600, 2, 2, 0),
    ('B3', 'Biscuit3', 100, 200, 2, 2, 0),
    ('D1', 'Drink1', 250, 400, 3, 1, 0),
    ('D2', 'Drink2', 330, 200, 3, 1, 0);

INSERT INTO stocks (stock_id, expiry_date, curr_qty, prod_id) VALUES
    ('FEB20', '2020-02-01', 0, 'P2'),
    ('MAR20', '2020-03-01', 0, 'B1'),
    ('JAN20', '2020-01-01', 0, 'P1'),
    ('FEB20', '2020-02-01', 0, 'B1'),
    ('FEB20', '2020-02-01', 0, 'D2'),
    ('FEB20', '2020-02-01', 0, 'D1'),
    ('FEB20', '2020-02-01', 0, 'B2');


call add_stock('MAR21', 'P1', 80, '2021-03-03');
call add_stock('MAR21', 'P1', 80, '2021-03-03');
call add_stock('FEB22', 'P1', 80, '2023-03-03');
call add_stock('FEB22', 'P1', 80, '2023-03-03');
call add_stock('FEB22', 'P1', 60, '2023-03-03');
call add_stock('FEB22', 'P1', 60, '2023-03-03');
call add_stock('FEB22', 'P1', 80, '2023-03-03');
call add_stock('FEB22', 'P1', 80, '2023-03-03');
call add_stock('FEB22', 'P1', 300, '2023-03-03');
call add_stock('FEB22', 'P1', 100, '2023-03-03');

call remove_expired();

call add_stock('FEB22', 'P2', 30, '2023-03-03');
call add_stock('FEB22', 'B1', 40, '2023-03-03');
call add_stock('FEB22', 'B2', 10, '2023-03-03');
call add_stock('FEB22', 'B3', 80, '2023-03-03');
call add_stock('FEB22', 'D1', 120, '2023-03-03');
call add_stock('FEB22', 'D2', 90, '2023-03-03');

call remove_stock('P2', 200);
call remove_stock('B2', 10);
call remove_stock('P2', 30);
call remove_stock('D2', 85);
