-- Create procedures
DELIMITER //
CREATE PROCEDURE add_stock (
    IN stock_id VARCHAR(30),
    IN prod_id VARCHAR(30),
    IN qty INT,
    IN expiry DATE)
        BEGIN
            DECLARE curr_date DATE DEFAULT CURDATE();
            DECLARE status, action VARCHAR(30);
            DECLARE prod_count, stock_count, curr_qty, max_qty INT;
            DECLARE pid INT;
            SET action = "ADD";
            START TRANSACTION;
            SELECT COUNT(*) INTO prod_count
                FROM products AS p
                WHERE p.prod_id = prod_id;
            IF prod_count = 0 THEN
                SET STATUS = "REJECTED - Unknown product";
            ELSE
                SELECT COUNT(*) INTO stock_count
                    FROM stocks AS s
                    WHERE s.stock_id = stock_id;
                IF stock_count = 0 THEN
                    INSERT INTO stocks(stock_id, expiry_date, prod_id) VALUES
                        (stock_id, expiry, prod_id);
                END IF;
                SELECT s.prod_id INTO pid
                    FROM stocks AS s
                    WHERE s.stock_id = stock_id;
                IF pid <> prod_id THEN
                    SET STATUS = "REJECTED - Mismatch in prod_id";
                ELSE
                    IF expiry < curr_date THEN
                        SET status = "REJECTED - Expired stock";
                    ELSE
                        SELECT total_qty, max_capacity INTO curr_qty, max_qty
                            FROM products AS p
                            WHERE p.prod_id = prod_id;
                        IF curr_qty + qty > max_qty THEN
                            SET status = "REJECTED - Over max capacity";
                        ELSE
                            UPDATE products AS p
                                SET p.total_qty = p.total_qty + qty
                                WHERE p.prod_id = prod_id;
                            UPDATE stocks AS s
                                SET s.curr_qty = s.curr_qty + qty
                                WHERE s.stock_id = stock_id;
                            SET status = "ACCEPTED";
                        END IF;
                    END IF;
                END IF;
            END IF;
            INSERT INTO stock_logs(stock_id, prod_id, qty, date_arrived, expiry_date, action, status) VALUES
                (stock_id, prod_id, qty, curr_date, expiry, action, status);
            COMMIT;
        END //
DELIMITER ;
