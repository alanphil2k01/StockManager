package types

const (
	USER = iota
	ADIMINISTATOR
)

type Product struct {
	prodCode string
	prodName string
	category string
	rate     uint
	maxQty   uint
}

type Batch struct {
	batchCode string
	producode string
	batchQty  uint
	currQty   uint
}

type Stocks struct {
	id        uint
	prodCode  string
	batchCode string
	qty       uint
}

type ProductCategory struct {
	category   string
	storageLoc string
}

// ID         uint `gorm:"primaryKey" json:"id"`
// Name       string
// SupplierId uint
// CreatedAt  time.Time
// UpdatedAt  time.Time
// DeletedAt  gorm.DeletedAt `gorm:"index"`

// type Supplier struct {
// }
//
// type User struct {
// 	Id       uint
// 	username string
// 	email    string
// 	password string
// }
