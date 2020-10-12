package jwt

import (
	"essential/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_secret")

type Claims struct {
	UserId             uint
	jwt.StandardClaims //负载
}

func ReleaseToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //token 过期时间
	Claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),     //签发时间
			Issuer:    "oceanlearn.tech",     // 签发人
			Subject:   "user token",          // 主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenString, err := token.SignedString(jwtKey) //制造签名

	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ParseToken(tokenSting string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenSting, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
