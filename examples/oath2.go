package main

import (
	"log"

	"github.com/j4ng5y/onelogin-go/api/client"
)

func OAUTH2Requests(C *client.Client) {
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