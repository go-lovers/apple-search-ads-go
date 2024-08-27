package asa

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserAcl(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &UserACLListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AccessControlList.GetUserACL(ctx)
	})
}

func TestDeserialization(t *testing.T) {
	t.Parallel()

	sampleJSONResponse := "../test/response_body_json_files/user_acl_list_response.json"

	content, err := ioutil.ReadFile(sampleJSONResponse)
	assert.NoError(t, err)

	model := &UserACLListResponse{}

	err = json.Unmarshal(content, model)
	assert.NoError(t, err)

	assert.Equal(t, len(model.UserAcls), 1)
	assert.Nil(t, model.Pagination)

	userACL := model.UserAcls[0]

	assert.Equal(t, "USD", userACL.Currency)
	assert.Equal(t, 1, int(userACL.OrgID))
	assert.Equal(t, "orgName", userACL.OrgName)
	assert.Equal(t, PaymentModelLoc, userACL.PaymentModel)

	assert.Equal(t, []UserACLRoleName{UserACLRoleNameAPIAccountManager}, userACL.RoleNames)
	assert.Equal(t, "Asia/Hong_Kong", string(userACL.TimeZone))
}
