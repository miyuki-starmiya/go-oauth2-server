package main

import (
	"fmt"
	"os"

	"go-oauth2-server/auth/util"
)

func main() {
	// argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	fmt.Println(util.EncodeClientBase64(argsWithoutProg[0], argsWithoutProg[1]))
}
