package asa

import (
	"context"
	"testing"
)

func TestGetAllAdGroups(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AdGroups.GetAllAdGroups(ctx, 1, &GetAllAdGroupsQuery{})
	})
}

func TestGetAdGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AdGroups.GetAdGroup(ctx, 1, 99)
	})
}

func TestDeleteAdGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.AdGroups.DeleteAdGroup(ctx, 1, 99)
	})
}

func TestUpdateAdGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AdGroups.UpdateAdGroup(ctx, 1, 99, &AdGroupUpdateRequest{})
	})
}

func TestCreateAdGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AdGroups.CreateAdGroup(ctx, 1, &AdGroup{})
	})
}

func TestFindAdGroups(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &AdGroupListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AdGroups.FindAdGroups(ctx, 1, &Selector{})
	})
}
