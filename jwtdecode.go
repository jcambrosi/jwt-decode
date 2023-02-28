// jca 2023 jwt-decode does raw decode of jwt tokens without validation, prints claims
// this is for development/test use only

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
)

func main() {

	var tokenString string
	// var tokenString = "eyJraWQiOiJiaFhhOXZCMjlzanlMOHNzYXBRX0NLdVlIQWt0YkVRd1lTYjljSHJQSXhnIiwiYWxnIjoiUlMyNTYifQ.eyJ2ZXIiOjEsImp0aSI6IkFULjZFb3ZQNjNSTEw0OXdVSUY3aVhvWDRNVzAyME1vQjlyZFBfRkJLTWtYSncub2FyMXpqc2UxYmdqYk5mb2IwaDciLCJpc3MiOiJodHRwczovL2hwZS1ncmVlbmxha2UtY2Zycjd1am9wbmxxdmM4ZGc2bWcub2t0YXByZXZpZXcuY29tL29hdXRoMi9kZWZhdWx0IiwiYXVkIjoiYXBpOi8vZGVmYXVsdCIsImlhdCI6MTY3NzU3NDc5MSwiZXhwIjoxNjc3NTc1NjkxLCJjaWQiOiIwb2Exa2FmOTRmZktmbzV1MTBoOCIsInVpZCI6IjAwdTFrYW9jdHpiWU83WFVQMGg4Iiwic2NwIjpbIm9mZmxpbmVfYWNjZXNzIiwiZW1haWwiLCJvcGVuaWQiLCJwcm9maWxlIl0sImF1dGhfdGltZSI6MTY3NzU3NDc5MCwic3ViIjoiMDB1MWthb2N0emJZTzdYVVAwaDgiLCJ0ZW5hbnRJZCI6ImNmcnI3dWpvcG5scXZjOGRnNm1nIn0.I6m0pPQ6-RYaI2JG2xKD42VV8EBKnwQ7_FP2_JDoo364NAt_0cDWsGl9k4i77O7geewCvwDt4sXlEy2LWlaAzcaF4zMGMn_9cgSgOQP3mWV6MIlj1Uo0HbMct7xT-_8W6pw6bmLciKNVPeHbEXv7YR30NBxYwuW-EZPZrpEyCoG8BBaflehWRn3ktkOLXi5JxXGzehJkj6yl4PMIGWHx6rkQL1xrF6n0X_-9-7NvlXm5cyqUc6bQU7rdtc1eu2xEC9fVNwGDsSCnWvhE5CbazlhtceKY5aE4CG6AF6oBCx5sAkUeNqaDa7x3elxdEubXLNOAp-2Ww_gg-3BPUKlxVw"

	// read jwt token from console
	fmt.Println("jwtdecode - raw decode jwt tokens, no validations, prints claims, for development use only")
	fmt.Println("enter jwt token:")
	reader := bufio.NewReader(os.Stdin)
	tokenString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	// parse and list all claims in token
	token, err := jwt.Parse(tokenString, nil)
	if err != nil {
		fmt.Println(err)
	}
	if token == nil {
		fmt.Println("error: empty parsed token")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// jwt.MapClaims is map[string]interface{}
		fmt.Println("\nclaims map")
		fmt.Println("==========")
		for k, v := range claims {
			fmt.Println(k, ":", v)
		}
		fmt.Println("==========")
	} else {
		fmt.Println("error: token.Claims is not of jwt.MapClaims type")
	}
}
