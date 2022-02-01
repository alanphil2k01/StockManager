-- Create views
CREATE VIEW expired_products AS
    SELECT * FROM stocks WHERE expiry_date < CURDATE();
