package client

import (
	"net/url"
)

func (c *WowsClient) Do(target string, values *url.Values, result any) error {
	values.Set("application_id", c.APIKey)
	values.Set("language", c.Language)

	_, err := c.HTTPClient.R().
		SetFormDataFromValues(*values).
		SetSuccessResult(result).Post(target)

	return err
}
