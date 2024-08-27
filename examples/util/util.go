package util

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/go-lovers/apple-search-ads-go/asa"
	"io"
	"log"
	"os"
	"time"
)

var (
	orgID          = flag.String("oid", "", "org ID")
	keyID          = flag.String("kid", "", "key ID")
	teamID         = flag.String("tid", "", "team ID")
	clientID       = flag.String("cid", "", "client ID")
	privateKey     = flag.String("privatekey", "", "private key used to sign authorization token")
	privateKeyPath = flag.String("privatekeypath", "", "path to a private key used to sign authorization token")
)

// TokenConfig creates the auth transport using the required information
func TokenConfig() (auth *asa.AuthTransport, err error) {
	var secret []byte
	if *privateKey != "" {
		secret = []byte(*privateKey)
	} else if *privateKeyPath != "" {
		// Read private key file as []byte
		secret, err = os.ReadFile(*privateKeyPath)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("no private key provided to either the -privatekey or -privatekeypath flags")
	}

	// Create the token using the required information
	auth, err = asa.NewTokenConfig(*orgID, *keyID, *teamID, *clientID, 20*time.Minute, secret)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// GetCampaign returns a single asa.Campaign by filtering by its selector on Apple Search Ads
func GetCampaign(ctx context.Context, client *asa.Client, params *asa.Selector) (*asa.Campaign, error) {
	apps, _, err := client.Campaigns.FindCampaigns(ctx, params)
	if err != nil {
		return nil, err
	} else if len(apps.Campaigns) == 0 {
		return nil, fmt.Errorf("query for campaign returned no campaign")
	}
	return apps.Campaigns[0], nil
}

// Close closes an open descriptor.
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
