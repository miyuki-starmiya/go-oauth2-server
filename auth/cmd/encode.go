package main

import (
	"fmt"
	"os"

	"github.com/miyuki-starmiya/go-oauth2-server/auth/util"
)

func main() {
	// encode client id and secret
	argsWithoutProg := os.Args[1:]

	fmt.Println(util.EncodeClientBase64(argsWithoutProg[0], argsWithoutProg[1]))
}
