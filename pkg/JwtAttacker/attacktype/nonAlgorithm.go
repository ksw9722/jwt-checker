package attack

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type nonAlgorithm struct {
	Jwt         string
	PayloadList []string
}

func (jwtObject *nonAlgorithm) Attack() {
	token, _, err := new(jwt.Parser).ParseUnverified(jwtObject.Jwt, jwt.MapClaims{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token.Claims)

}

func NonAlgorithm(tokenString string) *nonAlgorithm {
	return &nonAlgorithm{tokenString, make([]string, 0)}
}
