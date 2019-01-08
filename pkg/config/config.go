package config

import (
	"strings"

	"github.com/spf13/viper"
)

const FileName = "podsync"

type AppConfig struct {
	YouTubeAPIKey          string `yaml:"youtubeApiKey"`
	VimeoAPIKey            string `yaml:"vimeoApiKey"`
	PatreonClientID        string `yaml:"patreonClientId"`
	PatreonSecret          string `yaml:"patreonSecret"`
	PatreonRedirectURL     string `yaml:"patreonRedirectUrl"`
	PatreonWebhooksSecret  string `json:"patreonWebhooksSecret"`
	PostgresConnectionURL  string `yaml:"postgresConnectionUrl"`
	RedisURL               string `yaml:"redisUrl"`
	CookieSecret           string `yaml:"cookieSecret"`
	AWSAccessKey           string `yaml:"awsAccessKey"`
	AWSAccessSecret        string `yaml:"awsAccessSecret"`
	DynamoFeedsTableName   string `yaml:"dynamoFeedsTableName"`
	DynamoPledgesTableName string `yaml:"dynamoPledgesTableName"`
}

func ReadConfiguration() (*AppConfig, error) {
	viper.SetConfigName(FileName)

	// Configuration file
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app/config/")

	// Env variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	envmap := map[string]string{
		"youtubeApiKey":          "YOUTUBE_API_KEY",
		"vimeoApiKey":            "VIMEO_API_KEY",
		"patreonClientId":        "PATREON_CLIENT_ID",
		"patreonSecret":          "PATREON_SECRET",
		"patreonRedirectUrl":     "PATREON_REDIRECT_URL",
		"patreonWebhooksSecret":  "PATREON_WEBHOOKS_SECRET",
		"postgresConnectionUrl":  "POSTGRES_CONNECTION_URL",
		"redisUrl":               "REDIS_CONNECTION_URL",
		"cookieSecret":           "COOKIE_SECRET",
		"awsAccessKey":           "AWS_ACCESS_KEY",
		"awsAccessSecret":        "AWS_ACCESS_SECRET",
		"dynamoFeedsTableName":   "DYNAMO_FEEDS_TABLE_NAME",
		"dynamoPledgesTableName": "DYNAMO_PLEDGES_TABLE_NAME",
	}

	for k, v := range envmap {
		viper.BindEnv(k, v)
	}

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	cfg := &AppConfig{}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
