package auth

import (
	"encoding/json"
	"net/http"

	"github.com/RobsonFeitosa/go-driver/internal/users"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = "698dc19d489c4e4db73e28a713eab07b"

func createToken(user users.User) (string, error) {

}

type Claims struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(rw http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
