package main

import (
	"flag"
	"log"

	"github.com/j4ng5y/onelogin/api/client"
)

func main() {
	var clientID, clientSecret string
	flag.StringVar(&clientID, "id", "", "your onelogin ClientID")
	flag.StringVar(&clientSecret, "secret", "", "your onelogin ClientSecret")
	flag.Parse()

	C, err := client.New(clientID, clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	GAT, err := C.GetAccessToken(&client.GetAccessTokenRequest{
		GrantType: "client_credentials",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("GetAccessToken: %v\n", GAT)

	RT, err := C.RegenerateToken(&client.RegenerateTokenRequest{
		GrantType:    "refresh_token",
		AccessToken:  GAT.AccessCredentials,
		RefreshToken: GAT.RefreshToken,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("RegenerateAccessToken: %v\n", RT)

	GRL, err := C.GetRateLimits(RT.AccessToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("GetRateLimits: %v\n", GRL)

	RT2, err := C.RevokeToken(&client.RevokeTokenRequest{
		AccessToken:  RT.AccessToken,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("RevokeToken: %v\n", RT2)
}