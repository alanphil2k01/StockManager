-- Create procedures
DELIMITER //
CREATE PROCEDURE add_log(
    IN stock_id VARCHAR(30),
    IN prod_id VARCHAR(30),
    IN qty INT,
    IN date_added DATE,
    IN expiry_date DATE,
    IN action  VARCHAR(30),
    IN status  VARCHAR(30)
)
BEGIN
    INSERT INTO stock_logs(stock_id, prod_id, qty, date_processed, expiry_date, action, status) VALUES
        (stock_id, prod_id, qty, date_added, expiry_date, action, status);
END //
DELIMITER ;

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
    DECLARE pid VARCHAR(30);
    SET action = "ADD";
    START TRANSACTION;
    IF expiry < curr_date THEN
        SET status = "REJECTED - Expired stock";
    ELSE
        SELECT COUNT(*) INTO prod_count
        FROM products AS p
        WHERE p.prod_id = prod_id;
        IF prod_count = 0 THEN
            SET STATUS = "REJECTED - Unknown product";
        ELSE
            SELECT COUNT(*) INTO stock_count
            FROM stocks AS s
            WHERE s.stock_id = stock_id AND s.prod_id = prod_id;
            IF stock_count = 0 THEN
                INSERT INTO stocks(stock_id, expiry_date, prod_id) VALUES
                (stock_id, expiry, prod_id);
            END IF;
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
                    WHERE s.stock_id = stock_id AND s.prod_id = prod_id;
                SET status = "ACCEPTED";
            END IF;
        END IF;
    END IF;
    CALL add_log(stock_id, prod_id, qty, curr_date, expiry, action, status);
    COMMIT;
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE remove_stock(
    IN prod_id VARCHAR(30),
    IN qty INT)
BEGIN
    DECLARE stock_id VARCHAR(30);
    DECLARE status, action VARCHAR(30);
    DECLARE stock_qty, total_qty INT;
    DECLARE expiry DATE;
    DECLARE finished INTEGER DEFAULT 0;
    DECLARE stock_cur CURSOR FOR
        SELECT s.stock_id, s.curr_qty, s.expiry_date FROM stocks as s
        WHERE s.prod_id = prod_id
        ORDER BY s.expiry_date;
    DECLARE CONTINUE HANDLER
        FOR NOT FOUND SET finished = 1;
    SELECT p.total_qty INTO total_qty
        FROM products AS p
        WHERE p.prod_id = prod_id;
    IF total_qty >= qty THEN
        START TRANSACTION;
        OPEN stock_cur;
        UPDATE products AS p SET p.total_qty = p.total_qty - qty WHERE p.prod_id = prod_id;
        removeStocks: LOOP
            FETCH stock_cur INTO stock_id, stock_qty, expiry;
            IF finished = 1 THEN
                LEAVE removeStocks;
            END IF;
            SELECT stock_id AS stock_id, prod_id AS produ_id;
            IF stock_qty > qty THEN
                UPDATE stocks AS s SET s.curr_qty = s.curr_qty - qty
                    WHERE s.stock_id = stock_id AND s.prod_id = prod_id;
                CALL add_log(stock_id, prod_id, qty, CURDATE(), expiry, "REMOVE", "MOVED TO STORE");
                LEAVE removeStocks;
            ELSE
                DELETE FROM stocks AS s WHERE s.stock_id = stock_id AND s.prod_id = prod_id;
                SET qty = qty - stock_qty;
                CALL add_log(stock_id, prod_id, stock_qty, CURDATE(), expiry, "REMOVE", "OUT OF STOCK");
                IF qty = 0 THEN
                    LEAVE removeStocks;
                END IF;
            END IF;
        END LOOP removeStocks;
        CLOSE stock_cur;
        COMMIT;
    END IF;
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE remove_expired()
BEGIN
    DECLARE stock_id, prod_id VARCHAR(30);
    DECLARE status, action VARCHAR(30);
    DECLARE stock_qty INT;
    DECLARE expiry DATE;
    DECLARE finished INTEGER DEFAULT 0;
    DECLARE stock_cur CURSOR FOR
        SELECT e.stock_id, e.prod_id, e.curr_qty, e.expiry_date FROM expired_stocks as e;
    DECLARE CONTINUE HANDLER
        FOR NOT FOUND SET finished = 1;
    START TRANSACTION;
    OPEN stock_cur;
    removeExpiredStocks: LOOP
        FETCH stock_cur INTO stock_id, prod_id, stock_qty, expiry;
        IF finished = 1 THEN
            LEAVE removeExpiredStocks;
        END IF;
        UPDATE products AS p SET p.total_qty = p.total_qty - stock_qty
            WHERE p.prod_id = prod_id;
        DELETE FROM stocks AS s WHERE s.stock_id = stock_id;
        CALL add_log(stock_id, prod_id, stock_qty, CURDATE(), expiry, "REMOVE", "EXPIRED");
    END LOOP removeExpiredStocks;
    CLOSE stock_cur;
    COMMIT;
END //
DELIMITER ;
