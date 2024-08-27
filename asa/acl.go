package asa

import (
	"context"
)

// AccessControlListService handles communication with build-related methods of the Apple Search Ads API
//
// https://developer.apple.com/documentation/apple_search_ads/calling_the_apple_search_ads_api
type AccessControlListService service

// UserACLListResponse is a container for ACL call responses
//
// https://developer.apple.com/documentation/apple_search_ads/useracllistresponse
type UserACLListResponse struct {
	UserAcls   []*UserACL         `json:"data,omitempty"`
	Error      *ErrorResponseBody `json:"error,omitempty"`
	Pagination *PageDetail        `json:"pagination,omitempty"`
}

// UserACLRoleName governs what a user can see and do within the account.
type UserACLRoleName string

const (
	// UserACLRoleNameAPIAccountManager is for Manage all campaigns within an account with read-and-write capabilities.
	UserACLRoleNameAPIAccountManager UserACLRoleName = "API Account Manager"
	// UserACLRoleNameAPIAccountReadOnly is for View reporting across the account with read-only permission.
	UserACLRoleNameAPIAccountReadOnly UserACLRoleName = "API Account Read Only"
	// UserACLRoleNameLimitedAccessAPIReadWrite is for View reporting.
	UserACLRoleNameLimitedAccessAPIReadWrite UserACLRoleName = "Limited Access: API Read & Write"
	// UserACLRoleNameLimitedAccessAPIReadOnly is View reporting across the organization.
	UserACLRoleNameLimitedAccessAPIReadOnly UserACLRoleName = "Limited Access: API Read Only"
)

// UserACL is the response to ACL requests
//
// https://developer.apple.com/documentation/apple_search_ads/useracl
type UserACL struct {
	Currency     string                   `json:"currency"`
	OrgID        int64                    `json:"orgId"`
	OrgName      string                   `json:"orgName"`
	ParentOrgID  int64                    `json:"parentOrgId"`
	PaymentModel PaymentModel             `json:"paymentModel,omitempty"`
	RoleNames    []UserACLRoleName        `json:"roleNames"`
	TimeZone     ReportingRequestTimeZone `json:"timeZone"`
}

// GetUserACL Fetches roles and organizations that the API has access to
//
// https://developer.apple.com/documentation/apple_search_ads/get_user_acl
func (s *AccessControlListService) GetUserACL(ctx context.Context) (*UserACLListResponse, *Response, error) {
	url := "acls"
	res := new(UserACLListResponse)
	resp, err := s.client.get(ctx, url, nil, res)

	return res, resp, err
}
