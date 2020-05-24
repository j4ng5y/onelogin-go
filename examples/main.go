package main

import (
	"flag"
	"log"

	"github.com/j4ng5y/onelogin-go/api/client"
	"github.com/j4ng5y/onelogin-go/api/session"
)

func main() {
	var clientID, clientSecret, region, subdomain string
	flag.StringVar(&clientID, "id", "", "your onelogin ClientID")
	flag.StringVar(&clientSecret, "secret", "", "your onelogin ClientSecret")
	flag.StringVar(&region, "region", "", "your onelogin Region")
	flag.StringVar(&subdomain, "subdomain", "", "your onelogin SubDomain")
	flag.Parse()

	sess, err := session.New(session.ClientID(clientID), session.ClientSecret(clientSecret), session.Region(region))
	if err != nil {
		log.Fatal(err)
	}

	C, err := client.New(sess, subdomain)
	if err != nil {
		log.Fatal(err)
	}

	OAUTH2Requests(C)
}