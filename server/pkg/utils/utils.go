package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/alanphil2k01/SSMC/pkg/types"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var signingKey string = GetEnv("JWT_SECRET_KEY", "SuperSectretPassword")

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func ParseBody(r *http.Request, x interface{}) {
	if err := json.NewDecoder(r.Body).Decode(x); err != nil {
		return
	}
}

func HashPass(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash)
}

func CompareHashPass(hash, password string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func ValidateAddress(name string) bool {
	nameRegex := regexp.MustCompile(`^[_,\/\.-0-9a-zA-Z\s]+$`)
	return nameRegex.MatchString(name)
}

func ValidateNameWithNumbers(name string) bool {
	nameRegex := regexp.MustCompile(`^[a-zA-Z0-9][_0-9a-zA-Z\s]*$`)
	return nameRegex.MatchString(name)
}

func ValidateStrID(id string) bool {
	idRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return idRegex.MatchString(id)
}

func ValidateDate(date string) bool {
	dateRegex := regexp.MustCompile(`^20\d{2}(-|\/)((0[1-9])|(1[0-2]))(-|\/)((0[1-9])|([1-2][0-9])|(3[0-1]))$`)
	return dateRegex.MatchString(date)
}

func ValidateName(name string) bool {
	nameRegex := regexp.MustCompile(`^[a-zA-Z][a-zA-Z\s]*$`)
	return nameRegex.MatchString(name)
}

func ValidatePhoneNo(phoneNo string) bool {
	emailRegex := regexp.MustCompile(`^\+?[0-9]{0,3}?[ ]?[0-9]{5,15}$`)
	return emailRegex.MatchString(phoneNo)
}

func GetJWTtoken(username string, role uint) (string, error) {
    claims := types.UserClaims {
        Username: username,
        Role: role,
        StandardClaims: jwt.StandardClaims {
            ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    ss, err := token.SignedString([]byte(signingKey))
    if err != nil {
        return "", err
    }
    return ss, err
}

func ValidateToken(jwtToken string, role uint) error {
    token, err := jwt.ParseWithClaims(
        jwtToken,
        &types.UserClaims{},
        func(token *jwt.Token) (interface{}, error) {
            return []byte(signingKey), nil
        },
    )
    if err  != nil {
        return err
    }
    claims, ok := token.Claims.(*types.UserClaims)
    if !ok {
        return errors.New("error parsing claims")
    }
    if claims.ExpiresAt < time.Now().UTC().Unix() {
        return errors.New("token expired")
    }
    if role > claims.Role {
        return errors.New("role is invalid in token")
    }
    return nil
}
