package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// collections names
const (
	TykAnalyticsDB = "tyk_analytics"

	PortalCatalogue     = "portal_catalogue"
	PortalConfiguration = "portal_configuration"
	PortalCSS           = "portal_css"
	PortalMenus         = "portal_menus"

	TykAnalyticsCollection = "tyk_analytics"
	TykAnalyticsAggregates = "tyk_analytics_aggregates"
	TykAnalyticsLicense    = "tyk_analytics_license"
	TykAnalyticsUsers      = "tyk_analytics_users"
	TykApi                 = "tyk_apis"
	TykOrganisation        = "tyk_organisation"
	TykPolicies            = "tyk_policies"
	TykUptimeAnalytics     = "tyk_uptime_analytics"
)

const MongoConfigPath = "/opt/tyk-gateway/mongodb.yml"

type MongoDB struct {
	Config struct {
		Schema   string `yaml:"schema"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
	}
}

func UnmarshalMongoConfig(path string) (*MongoDB, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Can't unmarshal config file. Err: %s", err)
	}

	mongoDBConfig := &MongoDB{}

	err = yaml.Unmarshal(b, mongoDBConfig)
	if err != nil {
		return nil, fmt.Errorf("Can't unmarshal yaml, please make sure that the file structure is spelled correctly. Err: %s", err)
	}

	return mongoDBConfig, nil
}

func CreateMongoDBURI(path string) (string, error) {
	mongoDB, err := UnmarshalMongoConfig(path)
	if err != nil {
		return "", err
	}

	uri := fmt.Sprintf("%s://%s:%s/%s", mongoDB.Config.Schema, mongoDB.Config.Host, mongoDB.Config.Port, mongoDB.Config.Database)
	return uri, nil
}
