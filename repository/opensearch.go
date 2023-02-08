package repository

import (
	"context"
	"fmt"
	"sentraponsel-product-generator/db"
	"strings"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func GetItem(email string, stage string, suffix string) (*opensearchapi.Response, error) {
	queryString := fmt.Sprintf(`{
		"query": {
			"bool": {
				"must": [
					{
						"match": {
							"Email.keyword": "%s"
						}
					}
				]
			}
		}	
	}`, email)

	content := strings.NewReader(queryString)
	search := opensearchapi.SearchRequest{
		Index: []string{fmt.Sprintf("parastar-%s-%s", stage, suffix)},
		Body:  content,
	}

	searchResponse, err := search.Do(context.Background(), db.OpensearchClient)
	if err != nil {
		return nil, err
	}

	return searchResponse, nil
}

func GetExistItem(ownerMobile string, noKtp string, noChip string, stage string, suffix string) (*opensearchapi.Response, error) {
	queryString := fmt.Sprintf(`{
		"query": {
			"bool": {
				"must": [
					{
						"match": {
							"OwnerMobile.keyword": "%s"
						}
					},
					{
						"match": {
							"NoKtp.keyword": "%s"
						}
					},
					{
						"match": {
							"NoChip.keyword": "%s"
						}
					}
				]
			}
		}	
	}`, ownerMobile, noKtp, noChip)

	fmt.Println(queryString)

	content := strings.NewReader(queryString)
	search := opensearchapi.SearchRequest{
		Index: []string{fmt.Sprintf("parastar-%s-%s", stage, suffix)},
		Body:  content,
	}

	searchResponse, err := search.Do(context.Background(), db.OpensearchClient)
	if err != nil {
		return nil, err
	}

	return searchResponse, nil
}
