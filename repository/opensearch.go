package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sentraponsel-product-generator/db"
	"sentraponsel-product-generator/model"
	"sentraponsel-product-generator/util"
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

func GetBatchItemActivity(size float64, stage string, suffix string, hasNext *bool) ([]model.Activity, float64, error) {
	var (
		res          map[string]interface{}
		queryString  string
		value        float64
		activityList []model.Activity
	)

	if *hasNext {
		queryString = fmt.Sprintf(`{
			"query": {
				"bool": {
					"must_not": [
						{
							"match": {
								"CompanyId.keyword": "01G2MVQNXQCFK2BPS5JE7AHV8D"
							}
						}
					]
				}
			},
			"size": 5000
		}`)
	}

	content := strings.NewReader(queryString)
	search := opensearchapi.SearchRequest{
		Index: []string{fmt.Sprintf("parastar-%s-%s", stage, suffix)},
		Body:  content,
	}

	searchResponse, err := search.Do(context.Background(), db.OpensearchClient)
	if err != nil {
		return nil, 0, err
	}

	if searchResponse == nil {
		log.Panicf("Data not exist")
	}

	responseBody, err := ioutil.ReadAll(searchResponse.Body)
	if nil != err {
		log.Panicf("Failed read opensearch response")
	}

	err = json.Unmarshal(responseBody, &res)
	if nil != err {
		log.Panicf("Failed unmarshal opensearch response")
	}

	if res["hits"] != nil {
		activityList = []model.Activity{}
		response := res["hits"].(map[string]interface{})

		total := response["total"].(map[string]interface{})
		value = total["value"].(float64)
		fmt.Println(value)
		for _, v := range response["hits"].([]interface{}) {
			var activity model.Activity

			data := make(map[string]interface{})

			err = util.MapToStruct(v, &data)
			if err != nil {
				log.Panicf(err.Error())
			}

			err = util.MapToStruct(data["_source"], &activity)
			if err != nil {
				log.Panicf(err.Error())
			}

			activityList = append(activityList, activity)
		}
	}

	return activityList, value, nil
}
