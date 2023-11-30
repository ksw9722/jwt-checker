package main

import (
	"flag"
	"fmt"

	attack "github.com/ksw9722/jwt-checker/pkg/JwtAttacker/attacktype"
)

func main() {
	var (
		tokenString = flag.String("t", "", "JWT 토큰 입력")
		method      = flag.String("m", "none", "공격 타입")
	)
	flag.Parse()

	switch *method {
	case "none": // JWT NONE 알고리즘 공격
		attackObject := attack.NonAlgorithm(*tokenString)
		attackObject.Attack("dummy")
	default:
		fmt.Println("[-] 현재 지원하지 않는 공격 타입을 입력하였습니다.")
		return
	}

}
