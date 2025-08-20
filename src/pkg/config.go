package pkg

import (
	"encoding/json"
	"github.com/spf13/viper"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type ApplicationConfig struct {
	Service           string `mapstructure:"service"`
	LogLevel          string `mapstructure:"log_level"`
	GrpcServerAddress string `mapstructure:"grpc_server_address"`
	DbConnection      string `mapstructure:"db_connection"`
	Postgres          struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Db       string `mapstructure:"db"`
		SslMode  string `mapstructure:"ssl_mode"`
		Timezone string `mapstructure:"timezone"`
	} `mapstructure:"postgres"`
}

var Config *ApplicationConfig

func LoadConfig() *ApplicationConfig {
	if Config != nil {
		return Config
	}

	configContents, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var rawConfig map[string]any
	err = json.Unmarshal(configContents, &rawConfig)
	if err != nil {
		panic(err)
	}

	configMap := transformKeys(rawConfig)
	err = viper.MergeConfigMap(configMap)
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv()

	Config = &ApplicationConfig{}
	err = viper.Unmarshal(Config)
	if err != nil {
		panic(err)
	}

	return Config
}

// transformKeys recursively converts camelCase keys to SCREAMING_SNAKE_CASE in a map[string]any.
func transformKeys(input map[string]any) map[string]any {
	output := make(map[string]any)
	for k, v := range input {
		newKey := camelToSnake(k)

		switch value := v.(type) {
		case map[string]any:
			output[newKey] = transformKeys(value)
		case []any:
			var newSlice []any
			for _, item := range value {
				if itemMap, ok := item.(map[string]any); ok {
					newSlice = append(newSlice, transformKeys(itemMap))
				} else {
					newSlice = append(newSlice, item)
				}
			}
			output[newKey] = newSlice
		default:
			output[newKey] = value
		}
	}
	return output
}

// camelToSnake converts a camelCase string to SCREAMING_SNAKE_CASE.
func camelToSnake(s string) string {
	var builder strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				nextRune, _ := utf8.DecodeRuneInString(s[i+1:])
				if nextRune != '\uFFFD' && !unicode.IsUpper(nextRune) && !unicode.IsDigit(nextRune) {
					builder.WriteRune('_')
				}
			}
		}
		builder.WriteRune(unicode.ToUpper(r))
	}
	result := builder.String()
	result = strings.ReplaceAll(result, "__", "_")

	return result
}
