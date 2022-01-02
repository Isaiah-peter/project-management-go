package utils

import (
	"time"

	"project-management/database"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var db = database.GetDB()

type Token struct {
	UserId  string
	IsAdmin bool
	jwt.StandardClaims
}

func GenerateToken(password, email string) map[string]interface{} {
	user := models.User

	if err := db.Where("email =?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}

	expireAt := time.Now().Add(time.Minute * 100000).Unix()

	msg, errf := Checkpassword(user.Password, password)
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		var resp = map[string]interface{}{"status": false, "message": msg}
		return resp
	}

	tk := &models.Token{
		UserID:  int(user.ID),
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		panic(err)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user
	return resp

}
