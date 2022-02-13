-- Create views
CREATE VIEW products_list AS
    SELECT p.prod_id, p.prod_name, p.rate, p.max_capacity, p.total_qty, p.supplier_id, s.s_name, p.cat_id, c.cat_name
        FROM products AS p, suppliers AS s, product_categories AS c
        WHERE p.supplier_id = s.supplier_id AND p.cat_id = c.cat_id;

CREATE VIEW stocks_list AS
    SELECT s.stock_id, s.expiry_date, s.curr_qty, s.prod_id, p.prod_name
        FROM products AS p, stocks AS s
        WHERE p.prod_id = s.prod_id;

CREATE VIEW expired_stocks AS
    SELECT * FROM stocks WHERE expiry_date < CURDATE();
