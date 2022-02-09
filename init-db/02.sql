-- Create views
CREATE VIEW expired_stocks AS
    SELECT * FROM stocks WHERE expiry_date < CURDATE();
