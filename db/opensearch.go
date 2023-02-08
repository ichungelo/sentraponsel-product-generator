package db

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/opensearch-project/opensearch-go"
)

var OpensearchClient *opensearch.Client

func ConnOpensearch() {
	var(

		address = os.Getenv("OPENSEARCH_ADDRESS")
		username = os.Getenv("OPENSEARCH_USERNAME")
		password = os.Getenv("OPENSEARCH_PASSWORD")
	)
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{address},
		Username:  username,
		Password:  password,
	})

	if err != nil {
		log.Panicln("Openseach connection failed: => ", err.Error())
	}

	fmt.Println("success connect to opensearch")
	OpensearchClient = client
}
