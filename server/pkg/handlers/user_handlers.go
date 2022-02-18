package handlers

import (
	"net/http"
	"strings"

	"github.com/alanphil2k01/SSMC/pkg/db"
	"github.com/alanphil2k01/SSMC/pkg/types"
	"github.com/alanphil2k01/SSMC/pkg/utils"
)

type secretKey struct {
    Secret string `json:"secret"`
}


func CheckAuth(next http.HandlerFunc, role uint) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
            responsMessage(w, r, "Error - invalid header format", http.StatusUnauthorized, nil)
            return
        }
        jwtToken := authHeader[1]
        err := utils.ValidateToken(jwtToken, role)
        if err != nil {
            responsMessage(w, r, "Error - unauthoried", http.StatusUnauthorized, err)
        } else {
            next.ServeHTTP(w, r)
        }
	})
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user types.Users
	var secret secretKey
	utils.ParseBody(r, &user)
	utils.ParseBody(r, &secret)
	if user.Username == "" || user.Password == "" || user.Name == "" || user.Email == "" || user.Role > 3 {
		responsMessage(w, r, "Error - invalid input json", http.StatusBadRequest, nil)
		return
	}
    if user.Role == types.STAFF {
        if secret.Secret != utils.GetEnv("STAFF_SECRET", "staff_secret") {
            responsMessage(w, r, "Error - unauthorized", http.StatusUnauthorized, nil)
            return
        }
    }
    if user.Role == types.ADIMINISTATOR {
        if secret.Secret != utils.GetEnv("ADMIN_SECRET", "admin_secret") {
            responsMessage(w, r, "Error - unauthorized", http.StatusUnauthorized, nil)
            return
        }
    }
    user.Password = utils.HashPass(user.Password)
	err := db.RegisterUser(user)
	if err != nil {
		responsMessage(w, r, "Error - registring user", http.StatusInternalServerError, err)
		return
	}
    tokenString, err := utils.GetJWTtoken(user.Username, user.Role)
    if err != nil {
		responsMessage(w, r, "Error - creating jwt", http.StatusInternalServerError, err)
		return
    }
    responsMessage(w, r, "Registered user", http.StatusOK, tokenString)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user types.Users
	var role uint
	utils.ParseBody(r, &user)
	if user.Username == "" || user.Password == "" {
		responsMessage(w, r, "Error - invalid input json", http.StatusBadRequest, nil)
		return
	}
	ok, role, err := db.LoginUser(user.Username, user.Password)
    if err != nil {
        responsMessage(w, r, "Error - logging in", http.StatusInternalServerError, err)
        return
    }
    if !ok {
        responsMessage(w, r, "Error - Invalid Credentials", http.StatusUnauthorized, err)
        return
    }
    tokenString, err := utils.GetJWTtoken(user.Username, role)
    if err != nil {
		responsMessage(w, r, "Error - registring user", http.StatusInternalServerError, err)
		return
    }
    responsMessage(w, r, "User logged in", http.StatusOK, tokenString)
}
