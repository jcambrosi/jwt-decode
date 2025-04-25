// jca 2023-2025 jwt-decode does raw decode of jwt tokens without validation, prints claims
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

	// read jwt token from console
	fmt.Println("jwtdecode - raw decoding of a jwt token, no validations, prints claims, for development use only")
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
