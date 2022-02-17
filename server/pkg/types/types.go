package types

import "github.com/golang-jwt/jwt"

const (
	USER = iota
	STAFF
	ADIMINISTATOR
)

type ReponeMsg struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type Products struct {
	Prod_id      string `json:"prod_id"`
	Prod_name    string `json:"prod_name"`
	Rate         uint   `json:"rate"`
	Total_qty    uint   `json:"total_qty"`
	Max_capacity uint   `json:"max_capacity"`
	Cat_id       uint   `json:"cat_id"`
	Supplier_id  uint   `json:"supplier_id"`
}

type Stocks struct {
	Stock_id    string `json:"stock_id"`
	Expiry_date string `json:"expiry_date"`
	Curr_qty    uint   `json:"curr_qty"`
	Prod_id     string `json:"prod_id"`
}

type StockLogs struct {
	Log_id         uint   `json:"log_id"`
	Stock_id       string `json:"stock_id"`
	Prod_id        string `json:"prod_id"`
	Qty            uint   `json:"qty"`
	Date_processed string `json:"date_processed"`
	Expiry_date    string `json:"expiry_date"`
	Action         string `json:"action"`
	Status         string `json:"status"`
}

type Suppliers struct {
	Supplier_id uint   `json:"supplier_id"`
	S_name      string `json:"s_name"`
	S_email     string `json:"s_email"`
	Manager     string `json:"manager"`
	Address     string `json:"address"`
	Phone_no    string `json:"phone_no"`
}

type ProductCategories struct {
	Cat_id        uint   `json:"cat_id"`
	Cat_name      string `json:"cat_name"`
	Warehouse_loc string `json:"warehouse_loc"`
}

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     uint   `json:"role"`
}

type UserClaims struct {
    Username string `json:"username"`
    Role     uint
    jwt.StandardClaims
}

type ProductsList struct {
	Prod_id      string `json:"prod_id"`
	Prod_name    string `json:"prod_name"`
	Rate         uint   `json:"rate"`
	Total_qty    uint   `json:"total_qty"`
	Max_capacity uint   `json:"max_capacity"`
	Supplier_id  uint   `json:"supplier_id"`
	S_name       string `json:"s_name"`
	Cat_id       uint   `json:"cat_id"`
	Cat_name     string `json:"cat_name"`
}

type StocksList struct {
	Stock_id    string `json:"stock_id"`
	Expiry_date string `json:"expiry_date"`
	Curr_qty    uint   `json:"curr_qty"`
	Prod_id     string `json:"prod_id"`
	Prod_name   string `json:"prod_name"`
}
