package client

import (
	"github.com/aURORA-JC/wows-go/pkg/constants"
	"github.com/imroc/req/v3"
	"time"
)

type WowsClient struct {
	APIKey     string
	Endpoint   string
	Language   string
	HTTPClient *req.Client
	Debug      bool
}

func NewClient(key string, endpointName string) *WowsClient {
	return &WowsClient{
		APIKey:   key,
		Endpoint: endpointName,
		Language: constants.WowsLanguageEnglish,
		HTTPClient: req.C().SetBaseURL(constants.GetEndpointByName(endpointName)).
			SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0").
			SetTimeout(5 * time.Second),
	}
}

func NewAsiaClient(key string) *WowsClient {
	return NewClient(key, constants.WowsAsiaServerName)
}

func NewEuClient(key string) *WowsClient {
	return NewClient(key, constants.WowsEuropeServerName)
}

func NewNAClient(key string) *WowsClient {
	return NewClient(key, constants.WowsNorthAmericanServerName)
}

func (c *WowsClient) SetAPIKey(key string) {
	c.APIKey = key
}

func (c *WowsClient) SetEndpoint(endpointName string) {
	c.Endpoint = endpointName
	c.HTTPClient.SetBaseURL(constants.GetEndpointByName(endpointName))
}

func (c *WowsClient) SetLanguage(language string) {
	c.Language = language
}

func (c *WowsClient) DebugMode() *WowsClient {
	c.Debug = true
	c.HTTPClient = c.HTTPClient.DevMode()
	return c
}
