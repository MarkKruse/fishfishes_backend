package configuration

import (
	"fishfishes_backend/common/mongo"
	"flag"
	"github.com/pkg/errors"
)

type ServiceConfiguration struct {
	DB            mongo.Config
	BackendAPIKey string
}

func NewServiceConfiguration() *ServiceConfiguration {
	return &ServiceConfiguration{}
}

func (cfg *ServiceConfiguration) ServiceFlags() {

	mongo.BindConfig(&cfg.DB)
	flag.StringVar(&cfg.BackendAPIKey, "backendAPIKey", "", "The API Key that is used for the Backend Access")

	flag.Parse()
}

func (c ServiceConfiguration) Check() error {

	if !c.DB.IsValid() {
		return errors.Errorf("Please provide a valid DB config.")
	}

	return nil
}
