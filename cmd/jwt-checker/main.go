package main

import (
	"fmt"

	attack "github.com/ksw9722/jwt-checker/pkg/JwtAttacker/attacktype"
)

func main() {
	fmt.Println("Hello World")
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"
	test := attack.NonAlgorithm(tokenString)
	fmt.Println(test)
	test.Attack()

}
