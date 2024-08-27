//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCampaigns(t *testing.T) {
	campaignListResponse, _, err := client.Campaigns.GetAllCampaigns(context.Background(), nil)
	assert.NoError(t, err, "ListCampaigns responded with an error")
	assert.NotEmpty(t, campaignListResponse.Campaigns, "ListCampaigns returned no campaignListResponse")

	firstCampaign := campaignListResponse.Campaigns[0]
	campaignResponse, _, err := client.Campaigns.GetCampaign(context.Background(), firstCampaign.ID)
	assert.NoError(t, err, "GetCampaign responded with an error")
	assert.NotNil(t, campaignResponse.Campaign.Name)

}
