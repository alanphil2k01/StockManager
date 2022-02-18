package db

import (
	"errors"

	"github.com/alanphil2k01/SSMC/pkg/types"
	"github.com/alanphil2k01/SSMC/pkg/utils"
)

func LoginUser(username, password string) (bool, uint, error) {
	var hash string
	var role, count uint
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username)
    if row.Scan(&count); count != 1 {
		return false, 0, errors.New("invalid credentials")
    }
	row = db.QueryRow("SELECT password, role FROM users WHERE username = ?", username)
	err := row.Scan(&hash, &role)
	if err != nil {
		return false, 0,err
	}
	return utils.CompareHashPass(hash, password), role, nil
}

func RegisterUser(user types.Users) error {
    var count uint
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", user.Email)
    if row.Scan(&count); count == 1 {
		return errors.New("email already exists")
    }
	stmt, err := db.Prepare("INSERT INTO users(username, password, email, name, role) VaLUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Username, user.Password, user.Email, user.Name, user.Role)
	if err != nil {
		return err
	}
	return nil
}
