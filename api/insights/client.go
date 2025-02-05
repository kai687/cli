package insights

import (
	"encoding/json"
	"time"

	algoliaInsights "github.com/algolia/algoliasearch-client-go/v4/algolia/insights"
)

// Client wraps the default Insights API client so that we can declare methods on it
type Client struct {
	*algoliaInsights.APIClient
}

// TODO: Need to add the CLIs user agent to this
// NewClient instantiates a new client for interacting with the Insights API
func NewClient(appID, apiKey string, region algoliaInsights.Region) (*Client, error) {
	client, err := algoliaInsights.NewClient(appID, apiKey, region)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

// GetEvents retrieves a number of events from the Algolia Insights API.
func (c *Client) GetEvents(startDate, endDate time.Time, limit int) (*EventsRes, error) {
	layout := "2006-01-02T15:04:05.000Z"
	params := map[string]any{
		"startDate": startDate.Format(layout),
		"endDate":   endDate.Format(layout),
		"limit":     limit,
	}
	res, err := c.CustomGet(c.NewApiCustomGetRequest("1/events").WithParameters(params))
	if err != nil {
		return nil, err
	}
	tmp, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	var eventsRes EventsRes
	err = json.Unmarshal(tmp, &eventsRes)
	if err != nil {
		return nil, err
	}

	return &eventsRes, err
}
