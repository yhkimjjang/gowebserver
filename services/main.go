package services

import (
	"yhkim/gowebserver/webconfig"
)

var SearchService *ElasticSearchService

// 사용할 서비스들을 등록함
func RegisterServices() {
	addSearchService()
}

// 검색 서비스 생성
func addSearchService() {
	esConfig := webconfig.SetupElasticsearchConfig()
	SearchService = new(ElasticSearchService)
	SearchService.InitElasicsearch(&esConfig)
}
