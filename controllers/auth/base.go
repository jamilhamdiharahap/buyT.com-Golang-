package auth

import (
	"net/http"

	"gorm.io/gorm"
)

type Tokens struct {
	Token string `json:"token"`
}

type BD struct {
	DB *gorm.DB
}

type Authorization interface {
	ExtractToken(r *http.Request) string
	GetTokenJwt(token http.Handler) http.Handler
	ApiKey(w http.Handler) http.Handler
	Register(w http.Handler) http.Handler
}
