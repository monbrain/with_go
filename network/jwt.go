package network

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

func UpbitJwt(access_key, secret_key string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["access_key"] = access_key
	u := uuid.NewV4()
	claims["nonce"] = u.String()
	token.Claims = claims

	signedToken, err := token.SignedString([]byte(secret_key))
	if err != nil {
		log.Println("SignedString Error")
		log.Fatal(err)
		return ""
	}
	return signedToken
}

// func PayloadJwt(access_key, secret_key string, payload url.Values) string {
// 	claims := make(jwt.MapClaims)
// 	claims["access_key"] = access_key
// 	claims["nonce"] = uuid.NewV4().String()
// 	qh := sha512.New()
// 	qh.Reset()
// 	qh.Write([]byte(payload.Encode()))
// 	claims["query_hash"] = hex.EncodeToString(qh.Sum(nil))
// 	claims["query_hash_alg"] = "SHA512"

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	signedToken, err := token.SignedString([]byte(secret_key))
// 	//log.Println(signedToken)
// 	if err != nil {
// 		log.Println("SignedString Error")
// 		log.Fatal(err)
// 		return ""
// 	}
// 	return signedToken
// }
