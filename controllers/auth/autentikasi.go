package auth

import (
	"encoding/json"
	"fmt"
	"kereta/config"
	"kereta/controllers"
	"kereta/models"
	"kereta/tools"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB
var email = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func init() {
	DB = config.Connect()
}

func ApiKey(key http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api := r.Header.Get("x-api-key")
		if api != "Golang" {
			controllers.Respon(w, 401, nil, "Invalid Api Key")
			return
		}
		key.ServeHTTP(w, r)
	})
}

func GetTokenJwt(token http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if reflect.ValueOf(auth).IsZero() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokens := extractToken(auth)

		_, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
			autcheck := "BuyT"
			checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(autcheck, false)
			if !checkAudience {
				w.WriteHeader(http.StatusUnauthorized)
				return nil, fmt.Errorf("Invalid Token")
			}
			Iss := "BuyT.com"
			checkIssuer := token.Claims.(jwt.MapClaims).VerifyIssuer(Iss, false)
			if !checkIssuer {
				w.WriteHeader(http.StatusUnauthorized)
				return nil, fmt.Errorf("Invalid Issuer")
			}
			return []byte("LOGIN_SECRET"), nil

		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token.ServeHTTP(w, r)
	})

}

func extractToken(auth string) string {

	split := strings.Split(auth, " ")
	if len(split) == 2 {
		return split[1]
	}
	return ""

}

func Login(w http.ResponseWriter, r *http.Request) {
	var user map[string]string
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		controllers.Respon(w, 400, nil, "Bad Request")
		return
	}
	var data models.User
	DB.Where("email = ?", user["email"]).First(&data)
	if data.Email == "" {
		controllers.Respon(w, 400, nil, "Email Not Found")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user["password"]))
	if err != nil {
		controllers.Respon(w, 400, nil, "Password Wrong")
		return
	}

	timesExpired := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    "BuyT.com",
		"Audience":  "BuyT",
		"ExpiresAt": timesExpired.Unix(),
	})

	tokenString, err := token.SignedString([]byte("LOGIN_SECRET"))

	fmt.Println(tokenString, err)
	log.Println(tokenString)
	if err != nil {
		log.Println("Error", err)
		controllers.Respon(w, 500, nil, "Error Generate JWT !!")
		return
	}

	var tokensMaps Tokens
	tokensMaps.Token = tokenString
	controllers.Respon(w, 200, tokensMaps, "Login Success")
}

func EmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return email.MatchString(e)
}

func PasswordValid(e string) bool {
	if len(e) < 6 || len(e) > 254 {
		return false
	}
	return true
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user map[string]string
	method := "POST"
	baseUrl := "http://localhost:8002/api/u1/user"

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		controllers.Respon(w, 400, nil, "Bad Request")
		return
	}

	if PasswordValid(user["password"]) == false || PasswordValid(user["name"]) == false {
		controllers.Respon(w, 400, nil, "Invalid Password Or Name")
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user["password"]), 14)
	data := models.User{
		Name:     user["name"],
		Email:    user["email"],
		Password: password,
	}
	// var dataUser models.User

	if EmailValid(data.Email) == false {
		controllers.Respon(w, 400, nil, "Invalid Email")
		return
	}

	if data.Name == "" || data.Email == "" {
		controllers.Respon(w, 400, nil, "Bad Request")
		return
	}

	dataString, _ := json.Marshal(data)
	code, result, err := tools.HTTPResponse(method, baseUrl, string(dataString), nil)
	if err != nil {
		controllers.Respon(w, code, nil, err.Error())
		return
	}

	log.Println(result)

	DB.Create(&data)
	controllers.Respon(w, 200, data, "Register Success")
}
