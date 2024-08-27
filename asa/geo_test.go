package asa

import (
	"context"
	"testing"
)

func TestSearchGeos(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &SearchEntityListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Geo.SearchGeos(ctx, &SearchGeoQuery{})
	})
}

func TestGetGeos(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &SearchEntityListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Geo.GetGeos(ctx, &ListGeoQuery{}, []*GeoRequest{})
	})
}
