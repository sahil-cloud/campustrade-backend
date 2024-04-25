package helper

import (
	"crypto/rand"
	"io"
	"log"
	"os"
	"reflect"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

type SignedDetails struct {
	Name      string
	ID        string
	CryptoKey string
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateToken(name string, id string, cryptoKey string) (signedToken string, err error) {
	claims := &SignedDetails{
		Name:      name,
		ID:        id,
		CryptoKey: cryptoKey,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		msg = err.Error()
		return
	}
	return claims, msg
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "Username  or Password is incorrect"
		check = false
	}
	return check, msg
}
func UnpackStruct(a interface{}) []string {
	s := reflect.ValueOf(a)
	ret := make([]string, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		ret[i] = s.Field(i).String()
	}
	return ret
}

var numbers = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateOTP() (string, error) {
	max := 6
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return "", err
	}
	for i := 0; i < len(b); i++ {
		b[i] = numbers[int(b[i])%len(numbers)]
	}
	return string(b), nil
}

func TodayDateTime() string {
	current := time.Now()
	formattedDate := current.Format("2006-01-02T15:04:05Z")
	return formattedDate
}
