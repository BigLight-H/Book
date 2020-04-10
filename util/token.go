package util

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const (
	KEY string = "JWT-ARY-STARK"
	DEFAULT_EXPIRE_SECONDS int = 600 // default 10 minutes

)

type User struct {
	Id string `json:"id"`
	Name string `json:"json"`
}


// JWT -- json web token
// HEADER PAYLOAD SIGNATURE
// This struct is the PAYLOAD
type MyCustomClaims struct {
	User
	jwt.StandardClaims
}


// update expireAt and return a new token
func RefreshToken(tokenString string)(string, error) {
	// first get previous token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token)(interface{}, error) {
			return []byte(KEY), nil
		})
	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return "", err
	}
	mySigningKey := []byte(KEY)
	expireAt  := time.Now().Add(time.Second * time.Duration(DEFAULT_EXPIRE_SECONDS)).Unix()
	newClaims := MyCustomClaims{
		claims.User,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    claims.User.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	// generate new token with new claims
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenStr, err := newToken.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("generate new fresh json web token failed !! error :", err)
		return  "" , err
	}
	return tokenStr, err
}


func ValidateToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token)(interface{}, error) {
			return []byte(KEY), nil
		})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		spew.Dump(claims.User, claims.StandardClaims.ExpiresAt)
		FileLogs.Debug("token will be expired at ", time.Unix(claims.StandardClaims.ExpiresAt, 0))
	} else {
		FileLogs.Debug("validate tokenString failed !!!",err)
		return err
	}
	return nil
}

func GetTokenUserId(tokenString string) (id int){
	token, _ := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token)(interface{}, error) {
			return []byte(KEY), nil
		})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return StrToInt(claims.User.Id)
	}
	return 0
}



func GenerateToken(expiredSeconds int, id int, name string) (tokenString string) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}
	// Create the Claims
	mySigningKey := []byte(KEY)
	expireAt  := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	//fmt.Println("token will be expired at ", time.Unix(expireAt, 0) )
	FileLogs.Debug("token will be expired at ", time.Unix(expireAt, 0))
	// pass parameter to this func or not
	user := User{strconv.Itoa(id), name}
	claims := MyCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    user.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		//fmt.Println("generate json web token failed !! error :", err)
		FileLogs.Debug("generate json web token failed !! error :", err)
	}
	return tokenStr

}

// return this result to client then all later request should have header "Authorization: Bearer <token> "
func getHeaderTokenValue(tokenString string) string {
	//Authorization: Bearer <token>
	return fmt.Sprintf("Bearer %s", tokenString)
}

