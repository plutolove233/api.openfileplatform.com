//token生成

package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct{
	Username string
	jwt.StandardClaims
}

var MyKey = []byte("MyNameIsShyHao")

func GetToken(username string)(string,error){
	c := MyClaims{
		Username:username,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: time.Now().Unix()+60*60*2,
			Issuer: "PlutoLove233",
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString(MyKey)
}

func ParseToken(tokenString string)(*MyClaims,error){
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MyKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}