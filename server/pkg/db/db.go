package db

import (
	"github.com/alanphil2k01/SSMC/pkg/config"
	"github.com/alanphil2k01/SSMC/pkg/types"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	db = config.GetDB()
}

func GetProductById() {
}

func GetProductByName() {
}

func InsertProduct(product *types.Product) {
}

func GetSupplier() {
}

func InsertSupplier() {
}
