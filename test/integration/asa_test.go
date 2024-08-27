//go:build integration
// +build integration

package integration

import (
	"fmt"
	"github.com/go-lovers/apple-search-ads-go/asa"
	"os"
	"time"
)

const (
	envOrgID          = "ASA_INTEGRATION_OID"
	envKeyID          = "ASA_INTEGRATION_KID"
	envTeamID         = "ASA_INTEGRATION_TID"
	envClientID       = "ASA_INTEGRATION_CID"
	envPrivateKey     = "ASA_INTEGRATION_PRIVATE_KEY"
	envPrivateKeyPath = "ASA_INTEGRATION_PRIVATE_KEY_PATH"
)

var (
	client *asa.Client
)

func init() {
	token := tokenConfig()
	if token == nil {
		return
	}
	client = asa.NewClient(token.Client())
}

// TokenConfig creates the auth transport using the required information
func tokenConfig() *asa.AuthTransport {
	var privateKey []byte
	var err error
	if key := os.Getenv(envPrivateKey); key != "" {
		privateKey = []byte(key)
	} else if keyPath := os.Getenv(envPrivateKeyPath); keyPath != "" {
		// Read private key file as []byte
		privateKey, err = os.ReadFile(keyPath)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	} else {
		fmt.Println("no private key provided to either the ASA_INTEGRATION_PRIVATE_KEY or ASA_INTEGRATION_PRIVATE_KEY_PATH environment variables")
		return nil
	}
	orgID := os.Getenv(envOrgID)
	keyID := os.Getenv(envKeyID)
	teamID := os.Getenv(envTeamID)
	clientID := os.Getenv(envClientID)
	expiryDuration := 20 * time.Minute
	// Create the token using the required information
	auth, err := asa.NewTokenConfig(orgID, keyID, teamID, clientID, expiryDuration, privateKey)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return auth
}
