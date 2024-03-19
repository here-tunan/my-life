package service

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"go-my-life/internal/domain/repository/moneydb"
	"go-my-life/internal/infrastructure"
	"log"
)

func PutTransactionCategory(category *moneydb.TransactionCategory) error {
	var err error
	if category.Id == 0 {
		err = category.Insert()
	} else {
		err = category.Update()
	}
	return err
}

func AllCategory() ([]moneydb.TransactionCategory, error) {
	return moneydb.AllCategory()
}

// AnalysisCategory 根据描述进行分析
func AnalysisCategory(desc string) int {
	size := 1
	res, err := infrastructure.EsClient.Search().
		Index(moneydb.EsIndex).
		Request(&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"description": {Query: desc},
				},
			},
			Size: &size,
		}).Do(context.Background())
	if err != nil {
		return 0
	}

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	if res.Hits.Total.Value > 0 {
		// 取第一个
		jsonRaw := res.Hits.Hits[0].Source_
		doc := &moneydb.TransactionIndex{}
		err := json.Unmarshal(jsonRaw, doc)
		if err != nil {
			return 1
		}
		return doc.Category
	}
	return 1
}
