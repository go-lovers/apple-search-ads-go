package asa

import (
	"context"
	"testing"
)

func TestGetBudgetOrder(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BudgetOrderInfoResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Budget.GetBudgetOrder(ctx, 1)
	})
}

func TestGetAllBudgetOrders(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &BudgetOrderInfoListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Budget.GetAllBudgetOrders(ctx, &GetAllBudgetOrdersQuery{})
	})
}
