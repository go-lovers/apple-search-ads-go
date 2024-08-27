package asa

import (
	"context"
	"testing"
)

func TestSearchApps(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AppInfoListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.App.SearchApps(ctx, &SearchAppsQuery{})
	})
}
