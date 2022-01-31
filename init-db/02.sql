ALTER TABLE products ADD FOREIGN KEY (supplier_id) REFERENCES suppliers(supplier_id);
ALTER TABLE products ADD FOREIGN KEY (prod_category) REFERENCES product_categories(cat_id);
ALTER TABLE stocks ADD FOREIGN KEY (prod_id) REFERENCES products(prod_id);
ALTER TABLE stock_logs ADD FOREIGN KEY (prod_id) REFERENCES products(prod_id);
ALTER TABLE stock_logs ADD FOREIGN KEY (stock_id) REFERENCES stocks(stock_id);
