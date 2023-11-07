package configuration

import (
	"fishfishes_backend/common/mongo"
)

type ServiceConfiguration struct {
	DB            mongo.Config
	BackendAPIKey string
	PathServerPem string
	PathServerKey string
}

func NewServiceConfiguration(uri, database, apiKey string) *ServiceConfiguration {
	return &ServiceConfiguration{
		DB: mongo.Config{
			URI:      uri,
			Database: database,
		},
		BackendAPIKey: apiKey,
	}
}
