/*
Package asa is a Go client library for accessing Apple's Apple Search Ads API.
Usage
Import the package as you normally would:

	import "github.com/go-lovers/apple-search-ads-go/asa"

Construct a new Apple Search Ads client, then use the various services on the client to
access different parts of the Apple Search Ads API. For example:

	client := asa.NewClient(nil)

	params := &asa.SearchAppsQuery{
		Query:           "face",
		ReturnOwnedApps: false,
	}
	// list all apps with the query "face"
	apps, _, err := client.App.SearchApps(context.Background(), params)

The client is divided into logical chunks closely corresponding to the layout and structure
of Apple's own documentation at https://developer.apple.com/documentation/apple_search_ads.
For more sample code snippets, head over to the https://github.com/go-lovers/apple-search-ads-go/tree/master/examples directory.
Authentication
You may find that the code snippet above will always fail due to a lack of authorization.
The Apple Search Ads API has no methods that allow for unauthorized requests. To make it
easy to authenticate with Apple Search Ads, the apple-search-ads-go library offers a solution for signing
and rotating JSON Web Tokens automatically. For example, the above snippet could be made
to look a little more like this:

	import (
		"os"
		"time"

		"github.com/go-lovers/apple-search-ads-go/asa"
	)

	func main() {
		// Organization ID in Apple Search Ads
		orgID := "...."
		// Key ID for the given private key, described in Apple Search Ads
		keyID := "...."
		// Team ID for the given private key for the Apple Search Ads
		teamID := "...."
		// ClientID ID for the given private key for the Apple Search Ads
		clientID := "...."
		// A duration value for the lifetime of a token. Apple Search Ads does not accept a token with a lifetime of longer than 20 minutes
		expiryDuration = 20*time.Minute
		// The bytes of the private key created you have uploaded to it Apple Search Ads.
		privateKey = os.ReadFile("path/to/key")

		auth, err := asa.NewTokenConfig(orgID, keyID, teamID, clientID, expiryDuration, privateKey)
		if err != nil {
			return nil, err
		}
		client := asa.NewClient(auth.Client())

		// list all apps with the "face" in the authenticated user's team
		params := &asa.SearchAppsQuery{
			Offset:          0,
			Limit:           100,
			Query:           "face",
			ReturnOwnedApps: true,
		}

		apps, _, err := client.App.SearchApps(context.Background(), params)
		if err != nil {
			return nil, err
		}
	}

The authenticated client created here will automatically regenerate the token if it expires.
Also note that all Apple Search Ads APIs are scoped to the credentials of the pre-configured key,
so you can't use this API to make queries against the entire App Store. For more information on
creating the necessary credentials for the Apple Search Ads API, see the documentation at
https://developer.apple.com/documentation/appstoreconnectapi/creating_api_keys_for_app_store_connect_api.
Pagination
All requests for resource collections (apps, acls, ad groups, campaigns, etc.) support pagination.
Responses for paginated resources will contain a Pagination property of type PageDetail,
with TotalResults, StartIndex and ItemsPerPage.

	auth, _ := asa.NewTokenConfig(orgID, keyID, teamID, clientID, expiryDuration, privateKey)
	client := asa.NewClient(auth.Client())

	var allApps []asa.AppInfo
	params := &asa.SearchAppsQuery{
		Offset:          0,
		Limit:           100,
		Query:           "face",
		ReturnOwnedApps: false,
	}
	for {
		apps, _, err := client.App.SearchApps(context.Background(), params)
		if err != nil {
			return nil, err
		}

		allApps = append(allApps, apps.AppInfos...)

		pageDetail := apps.Pagination
		lastOffset := pageDetail.StartIndex + len(apps.AppInfos)
		if lastOffset < pageDetail.TotalResults {
			params.Offset += int32(len(apps.AppInfos))
		} else {
			break
		}
	}
*/
package asa
