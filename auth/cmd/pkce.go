package main

import (
	"fmt"

	"github.com/miyuki-starmiya/go-oauth2-server/auth/util"
)

func main() {
	// generate a code verifier and code challenge
	codeVerifier := util.GenerateCodeVerifier()
	codeChallenge := util.GenerateCodeChallenge(codeVerifier, "S256")

	fmt.Println("Code Verifier:", codeVerifier)
	fmt.Println("Code Challenge:", codeChallenge)
}
