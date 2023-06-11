package services

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"yhkim/gowebserver/types"
	"yhkim/gowebserver/webconfig"

	"github.com/elastic/go-elasticsearch/v8"
)

// ElasticSearch Clinet 객체를 소유하고 있는 구조체
type ElasticSearchService struct {
	Client *elasticsearch.Client
}

// ElasticSearch 서버와 통신하고 데이터를 주고 받을 수 있는 Client 객체 초기화 작업
func (es *ElasticSearchService) InitElasicsearch(config *webconfig.EsConfig) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			config.EsAddress,
		},
		Username:               config.EsUserName,
		Password:               config.EsPassword,
		CertificateFingerprint: config.EsFingerprint,
	}

	es.Client, _ = elasticsearch.NewClient(cfg)
}

// 검색어를 통해서 카드 조회
func (es *ElasticSearchService) FindCardBenefits(keyword string) []types.CardInfo {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"benefitList": keyword,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Client.Search(
		es.Client.Search.WithContext(context.Background()),
		es.Client.Search.WithIndex("credit_card"),
		es.Client.Search.WithBody(&buf),
		es.Client.Search.WithTrackTotalHits(true),
		es.Client.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(result["took"].(float64)),
	)

	cardInfos := []types.CardInfo{}

	// Print the ID and document source for each hit.
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		carInfo := types.CardInfo{
			Code:    source["code"].(string),
			Name:    source["name"].(string),
			Benefit: source["benefitList"].([]interface{}),
		}

		cardInfos = append(cardInfos, carInfo)
	}

	return cardInfos
}
