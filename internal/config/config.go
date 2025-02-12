// Package config provides configuration management using AWS Systems Manager Parameter Store
package config

import (
	"context"
	"log"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var (
	once     sync.Once
	instance *Config
)

type Config struct {
	Port              string
	Cors              map[string]bool
	AdminPassword     string
	DbDsn             string
	GmailAppPassword  string
	GmailAddress      string
	AlertPhoneNumbers []string
}

type ConfigDefinition struct {
	Path         string
	Type         string
	DefaultValue string
}

var configDefinitions = map[string]ConfigDefinition{
	"CORS_ORIGIN": {
		Path: "/backend/internal/admin-cors-origin",
		Type: "StringList",
	},
	"DB_DSN": {
		Path: "/backend/internal/db_dsn",
		Type: "SecureString",
	},
	"PORT": {
		Path:         "/backend/ports/admin",
		Type:         "String",
		DefaultValue: ":3100",
	},
	"ADMIN_PASSWORD": {
		Path: "/backend/internal/admin-password",
		Type: "SecureString",
	},
	"GMAIL_APP_PASSWORD": {
		Path: "/backend/internal/gmail-app-password",
		Type: "SecureString",
	},
	"GMAIL_ADDRESS": {
		Path: "/backend/internal/gmail-address",
		Type: "String",
	},
	"ALERT_PHONE_NUMBERS": {
		Path: "/backend/internal/alert-phone-numbers",
		Type: "StringList",
	},
}

func getSystemsManagerParameter(paramName string, ssmClient *ssm.Client) string {

	paramInfo, exists := configDefinitions[paramName]
	if !exists {
		log.Fatalf("***ERROR (config): Parameter '%s' not found in configDefinitions", paramName)
	}
	isEncrypted := paramInfo.Type == "SecureString"

	param, err := ssmClient.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           &paramInfo.Path,
		WithDecryption: &isEncrypted,
	})

	if err != nil {
		if paramInfo.DefaultValue != "" { // if parameter not found, return default value from configDefinitions
			return paramInfo.DefaultValue
		}
		log.Fatalf("***ERROR (config): Parameter '%s' not found in Systems Manager", paramName)
	}
	return *param.Parameter.Value
}

func LoadConfig() *Config {

	once.Do(func() {
		config, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal("***ERROR (config): Unable to load AWS SDK config")
		}
		ssmClient := ssm.NewFromConfig(config)

		corsString := getSystemsManagerParameter("CORS_ORIGIN", ssmClient) // comma separated list of URLs
		corsUrls := strings.Split(corsString, ",")                         // convert to list of URLs

		corsOrigin := make(map[string]bool, len(corsUrls)) // create map of URLs
		for _, url := range corsUrls {
			trimmedURL := strings.TrimSpace(url)
			if trimmedURL != "" {
				corsOrigin[trimmedURL] = true
			}
		}

		port := getSystemsManagerParameter("PORT", ssmClient)
		adminPassword := getSystemsManagerParameter("ADMIN_PASSWORD", ssmClient)
		dbDsn := getSystemsManagerParameter("DB_DSN", ssmClient)
		gmailAppPassword := getSystemsManagerParameter("GMAIL_APP_PASSWORD", ssmClient)
		gmailAddress := getSystemsManagerParameter("GMAIL_ADDRESS", ssmClient)
		alertPhoneNumbersString := getSystemsManagerParameter("ALERT_PHONE_NUMBERS", ssmClient)

		instance = &Config{
			Port:              port,
			Cors:              corsOrigin,
			AdminPassword:     adminPassword,
			DbDsn:             dbDsn,
			GmailAppPassword:  gmailAppPassword,
			GmailAddress:      gmailAddress,
			AlertPhoneNumbers: strings.Split(alertPhoneNumbersString, ","),
		}

	})

	return instance
}
