package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-lovers/apple-search-ads-go/asa"
	"github.com/go-lovers/apple-search-ads-go/examples/util"
	"log"
)

var (
	campaignName = flag.String("campaign_name", "", "Campaign Name for an Campaign")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the Apple Search Ads client
	client := asa.NewClient(auth.Client())

	campaign, err := util.GetCampaign(ctx, client, &asa.Selector{
		Conditions: []*asa.Condition{
			{
				Field:    "Name",
				Operator: asa.ConditionOperatorEquals,
				Values:   []string{*campaignName},
			},
		},
	})
	if err != nil {
		log.Fatalf("%s", err)
	}

	params := &asa.GetAllAdGroupsQuery{}
	for {
		adGroupsResponse, _, err := client.AdGroups.GetAllAdGroups(ctx, campaign.ID, params)
		if err != nil {
			log.Fatal(err)
		}

		for _, adGroup := range adGroupsResponse.AdGroups {
			fmt.Println(adGroup.Name)
		}
		pageDetail := adGroupsResponse.Pagination
		lastOffset := pageDetail.StartIndex + len(adGroupsResponse.AdGroups)
		if lastOffset < pageDetail.TotalResults {
			params.Offset += int32(len(adGroupsResponse.AdGroups))
		} else {
			break
		}
	}
}
