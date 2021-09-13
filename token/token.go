package token

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// Structured version of Claims Section, as referenced at
// https://tools.ietf.org/html/rfc7519#section-4.1
// See examples for how to use this with your own claim types
type TokenClaims struct {
	jwt.StandardClaims
	TokenKind    int      `json:"kin,omitempty"`
	AccountID    int      `json:"aid,omitempty"`
	Port         int      `json:"port,omitempty"`
	SessionLimit int      `json:"slt,omitempty"`
	Uid          int      `json:"uid,omitempty"`
	CompanyID    int      `json:"cid,omitempty"`
	CompanyKind  string   `json:"ckind,omitempty"`
	PID          int      `json:"pid,omitempty"`
	Roles        []string `json:"roles,omitempty"`
}

func CreateToken(key string, m map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for index, val := range m {
		claims[index] = val
	}
	// fmt.Println(_map)
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

func CreateTokenDefault(m map[string]interface{}) string {
	return CreateToken(gDefaultTokenKey, m)
}

func CreateTokenWithClaims(key string, claims *TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func CreateTokenWithClaimsDefault(claims *TokenClaims) (string, error) {
	return CreateTokenWithClaims(gDefaultTokenKey, claims)
}

func ParseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}

func ParseTokenDefault(tokenString string) (interface{}, bool) {
	return ParseToken(tokenString, gDefaultTokenKey)
}

func ParseTokenWithClaims(tokenString string, key string) (*TokenClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		fmt.Println("ParseTokenWithClaims, error:", err.Error())
		return nil, false
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}

func ParseTokenWithClaimsDefault(tokenString string) (*TokenClaims, bool) {
	return ParseTokenWithClaims(tokenString, gDefaultTokenKey)
}

func ParseTokenWithClaimsUnverified(tokenString string) (*TokenClaims, bool) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &TokenClaims{})
	if err != nil {
		fmt.Println("ParseTokenWithClaimsUnverified, error:", err.Error())
		return nil, false
	}

	if claims, ok := token.Claims.(*TokenClaims); ok {
		return claims, true
	} else {
		return nil, false
	}
}
