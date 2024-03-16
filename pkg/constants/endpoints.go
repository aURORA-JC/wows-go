package constants

import "log"

const (
	WowsAsiaServerName          = "asia"
	WowsEuropeServerName        = "eu"
	WowsNorthAmericanServerName = "na"

	WowsAsiaApiAddress          = "https://api.worldofwarships.asia/wows"
	WowsEuropeApiAddress        = "https://api.worldofwarships.eu/wows"
	WowsNorthAmericanApiAddress = "https://api.worldofwarships.com/wows"
)

var WowsApiEndpoints = map[string]string{
	WowsAsiaServerName:          WowsAsiaApiAddress,
	WowsEuropeServerName:        WowsEuropeApiAddress,
	WowsNorthAmericanServerName: WowsNorthAmericanApiAddress,
}

func GetEndpointByName(name string) string {
	endpoint := WowsApiEndpoints[name]

	if endpoint == "" {
		log.Println("[wows-go] Warning: can not find language, using default language: en")
		endpoint = WowsAsiaApiAddress
	}

	return endpoint
}
