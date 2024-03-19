package infrastructure

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"go-my-life/env"
	"log"
)

var EsClient *elasticsearch.TypedClient

func init() {
	address := fmt.Sprintf("%s:%d", env.Prop.ElasticSearch.Host, env.Prop.ElasticSearch.Port)
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}
	if env.Prop.ElasticSearch.Password != "" {
		cfg.Username = env.Prop.ElasticSearch.Username
		cfg.Password = env.Prop.ElasticSearch.Password
		cfg.CertificateFingerprint = env.Prop.ElasticSearch.Fingerprint
	}

	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	info := es.Info()
	log.Print(info)
	EsClient = es
}
