package main

import (
	"flag"
	"log"

	"github.com/j4ng5y/onelogin-go/api/client"
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

	OAUTH2Requests(C)
}