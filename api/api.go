package api

import (
	"fmt"
	"net/url"
)

type URLBuilderOptions struct {
	Region string
	// BaseURL can be retrieved by getting a value from the "URLS" map
	BaseURL  string
	ObjectID string
	ExtraID  string
	QueryParameters map[string]string
	SubDomain string
}

func URLBuilder(opts *URLBuilderOptions) (URL string, err error) {
	if opts.Region == "" {
		if opts.SubDomain == "" {
			return "", fmt.Errorf("one ofURLBuildOptions.Region or URLBuildOptions.SubDomain must be set")
		} else {
			if opts.ObjectID != "" {
				if opts.ExtraID != "" {
					if len(opts.QueryParameters) != 0 {
						base, err := url.Parse(opts.BaseURL)
						if err != nil {
							return "", nil
						}
						for k, v := range opts.QueryParameters {
							base.Query().Set(k, v)
						}
						return fmt.Sprintf(base.String(), opts.SubDomain, opts.ObjectID, opts.ExtraID), err
					} else {
						return fmt.Sprintf(opts.BaseURL, opts.SubDomain, opts.ObjectID, opts.ExtraID), err
					}
				} else {
					if len(opts.QueryParameters) != 0 {
						base, err := url.Parse(opts.BaseURL)
						if err != nil {
							return "", nil
						}
						for k, v := range opts.QueryParameters {
							base.Query().Set(k, v)
						}
						return fmt.Sprintf(base.String(), opts.SubDomain, opts.ObjectID), err
					} else {
						return fmt.Sprintf(opts.BaseURL, opts.SubDomain, opts.ObjectID), err
					}
				}
			}
			if opts.ExtraID != "" {
				if opts.ObjectID != "" {
					if len(opts.QueryParameters) != 0 {
						base, err := url.Parse(opts.BaseURL)
						if err != nil {
							return "", nil
						}
						for k, v := range opts.QueryParameters {
							base.Query().Set(k, v)
						}
						return fmt.Sprintf(base.String(), opts.SubDomain, opts.ObjectID, opts.ExtraID), err
					} else {
						return fmt.Sprintf(opts.BaseURL, opts.SubDomain, opts.ObjectID, opts.ExtraID), err
					}
				} else {
					if len(opts.QueryParameters) != 0 {
						base, err := url.Parse(opts.BaseURL)
						if err != nil {
							return "", nil
						}
						for k, v := range opts.QueryParameters {
							base.Query().Set(k, v)
						}
						return fmt.Sprintf(base.String(), opts.SubDomain, opts.ExtraID), err
					} else {
						return fmt.Sprintf(opts.BaseURL, opts.SubDomain, opts.ExtraID), err
					}
				}
			}
			if len(opts.QueryParameters) != 0 {
				base, err := url.Parse(opts.BaseURL)
				if err != nil {
					return "", nil
				}
				for k, v := range opts.QueryParameters {
					base.Query().Set(k, v)
				}
				return fmt.Sprintf(base.String(), opts.Region), err
			} else {
				return fmt.Sprintf(opts.BaseURL, opts.Region), err
			}
		}
	} else {
		if opts.ObjectID != "" {
			if opts.ExtraID != "" {
				if len(opts.QueryParameters) != 0 {
					base, err := url.Parse(opts.BaseURL)
					if err != nil {
						return "", nil
					}
					for k, v := range opts.QueryParameters {
						base.Query().Set(k, v)
					}
					return fmt.Sprintf(base.String(), opts.Region, opts.ObjectID, opts.ExtraID), err
				} else {
					return fmt.Sprintf(opts.BaseURL, opts.Region, opts.ObjectID, opts.ExtraID), err
				}
			} else {
				if len(opts.QueryParameters) != 0 {
					base, err := url.Parse(opts.BaseURL)
					if err != nil {
						return "", nil
					}
					for k, v := range opts.QueryParameters {
						base.Query().Set(k, v)
					}
					return fmt.Sprintf(base.String(), opts.Region, opts.ObjectID), err
				} else {
					return fmt.Sprintf(opts.BaseURL, opts.Region, opts.ObjectID), err
				}
			}
		}
		if opts.ExtraID != "" {
			if opts.ObjectID != "" {
				if len(opts.QueryParameters) != 0 {
					base, err := url.Parse(opts.BaseURL)
					if err != nil {
						return "", nil
					}
					for k, v := range opts.QueryParameters {
						base.Query().Set(k, v)
					}
					return fmt.Sprintf(base.String(), opts.Region, opts.ObjectID, opts.ExtraID), err
				} else {
					return fmt.Sprintf(opts.BaseURL, opts.Region, opts.ObjectID, opts.ExtraID), err
				}
			} else {
				if len(opts.QueryParameters) != 0 {
					base, err := url.Parse(opts.BaseURL)
					if err != nil {
						return "", nil
					}
					for k, v := range opts.QueryParameters {
						base.Query().Set(k, v)
					}
					return fmt.Sprintf(base.String(), opts.Region, opts.ExtraID), err
				} else {
					return fmt.Sprintf(opts.BaseURL, opts.Region, opts.ExtraID), err
				}
			}
		}
		if len(opts.QueryParameters) != 0 {
			base, err := url.Parse(opts.BaseURL)
			if err != nil {
				return "", nil
			}
			for k, v := range opts.QueryParameters {
				base.Query().Set(k, v)
			}
			return fmt.Sprintf(base.String(), opts.Region), err
		} else {
			return fmt.Sprintf(opts.BaseURL, opts.Region), err
		}
	}
}

var (
	// URLS is a mapping of API URLs categorized by function
	URLS = map[string]map[string]string{
		"OAUTH2_TOKEN_URLS": {
			"TOKEN_REQUEST_URL": `https://api.%s.onelogin.com/auth/oauth2/v2/token`,
			"TOKEN_REFRESH_URL": `https://api.%s.onelogin.com/auth/oauth2/v2/token`,
			"TOKEN_REVOKE_URL":  `https://api.%s.onelogin.com/auth/oauth2/revoke`,
			"GET_RATE_URL":      `https://api.%s.onelogin.com/auth/rate_limit`,
		},
		"USER_URLS": {
			"GET_USERS_URL":                    `https://api.%s.onelogin.com/api/1/users`,
			"GET_USER_URL":                     `https://api.%s.onelogin.com/api/1/users/%s`,
			"GET_APPS_FOR_USER_URL":            `https://api.%s.onelogin.com/api/1/users/%s/apps`,
			"GET_ROLES_FOR_USER_URL":           `https://api.%s.onelogin.com/api/1/users/%s/roles`,
			"GET_CUSTOM_ATTRIBUTES_URL":        `https://api.%s.onelogin.com/api/1/users/custom_attributes`,
			"CREATE_USER_URL":                  `https://api.%s.onelogin.com/api/1/users`,
			"UPDATE_USER_URL":                  `https://api.%s.onelogin.com/api/1/users/%s`,
			"DELETE_USER_URL":                  `https://api.%s.onelogin.com/api/1/users/%s`,
			"ADD_ROLE_TO_USER_URL":             `https://api.%s.onelogin.com/api/1/users/%s/add_roles`,
			"DELETE_ROLE_TO_USER_URL":          `https://api.%s.onelogin.com/api/1/users/%s/remove_roles`,
			"SET_PW_CLEARTEXT":                 `https://api.%s.onelogin.com/api/1/users/set_password_clear_text/%s`,
			"SET_PW_SALT":                      `https://api.%s.onelogin.com/api/1/users/set_password_using_salt/%s`,
			"SET_STATE_TO_USER_URL":            `https://api.%s.onelogin.com/api/1/users/%s/set_state`,
			"SET_CUSTOM_ATTRIBUTE_TO_USER_URL": `https://api.%s.onelogin.com/api/1/users/%s/set_custom_attributes`,
			"LOG_USER_OUT_URL":                 `https://api.%s.onelogin.com/api/1/users/%s/logout`,
			"LOCK_USER_URL":                    `https://api.%s.onelogin.com/api/1/users/%s/lock_user`,
			"GENERATE_MFA_TOKEN_URL":           `https://api.%s.onelogin.com/api/1/users/%s/mfa_token`,
		},
		"APPS_URLS": {
			"GET_APPS_URL": `https://api.%s.onelogin.com/api/1/apps`,
		},
		"ROLE_URLS": {
			"GET_ROLES_URL":   `https://api.%s.onelogin.com/api/1/roles`,
			"CREATE_ROLE_URL": `https://api.%s.onelogin.com/api/1/roles`,
			"GET_ROLE_URL":    `https://api.%s.onelogin.com/api/1/roles/%s`,
		},
		"EVENT_URLS": {
			"GET_EVENT_TYPES_URL": `https://api.%s.onelogin.com/api/1/events/types`,
			"GET_EVENTS_URL":      `https://api.%s.onelogin.com/api/1/events`,
			"CREATE_EVENT_URL":    `https://api.%s.onelogin.com/api/1/events`,
			"GET_EVENT_URL":       `https://api.%s.onelogin.com/api/1/events/%s`,
		},
		"GROUP_URLS": {
			"GET_GROUPS_URL":   `https://api.%s.onelogin.com/api/1/groups`,
			"CREATE_GROUP_URL": `https://api.%s.onelogin.com/api/1/groups`,
			"GET_GROUP_URL":    `https://api.%s.onelogin.com/api/1/groups/%s`,
		},
		"CUSTOM_LOGIN_URLS": {
			"SESSION_LOGIN_TOKEN_URL": `https://api.%s.onelogin.com/api/1/login/auth`,
			"GET_TOKEN_VERIFY_FACTOR": `https://api.%s.onelogin.com/api/1/login/verify_factor`,
		},
		"SAML_ASSERTION_URLS": {
			"GET_SAML_ASSERTION_URL": `https://api.%s.onelogin.com/api/1/saml_assertion`,
			"GET_SAML_VERIFY_FACTOR": `https://api.%s.onelogin.com/api/1/saml_assertion/verify_factor`,
		},
		"MFA_URLS": {
			"GET_FACTORS_URL":          `https://api.%s.onelogin.com/api/1/users/%s/auth_factors`,
			"ENROLL_FACTOR_URL":        `https://api.%s.onelogin.com/api/1/users/%s/otp_devices`,
			"GET_ENROLLED_FACTORS_URL": `https://api.%s.onelogin.com/api/1/users/%s/otp_devices`,
			"ACTIVATE_FACTOR_URL":      `https://api.%s.onelogin.com/api/1/users/%s/otp_devices/%s/trigger`,
			"VERIFY_FACTOR_URL":        `https://api.%s.onelogin.com/api/1/users/%s/otp_devices/%s/verify`,
			"DELETE_FACTOR_URL":        `https://api.%s.onelogin.com/api/1/users/%s/otp_devices/%s`,
		},
		"INVITE_LINK_URLS": {
			"GENERATE_INVITE_LINK_URL": `https://api.%s.onelogin.com/api/1/invites/get_invite_link`,
			"SEND_INVITE_LINK_URL":     `https://api.%s.onelogin.com/api/1/invites/send_invite_link`,
		},
		"EMBED_APPS_URL": {
			"EMBED_APP_URL": `https://api.onelogin.com/client/apps/embed2`,
		},
		"PRIVILEGES_URLS": {
			"LIST_PRIVILEGES_URL":                 `https://api.%s.onelogin.com/api/1/privileges`,
			"CREATE_PRIVILEGE_URL":                `https://api.%s.onelogin.com/api/1/privileges`,
			"UPDATE_PRIVILEGE_URL":                `https://api.%s.onelogin.com/api/1/privileges/%s`,
			"GET_PRIVILEGE_URL":                   `https://api.%s.onelogin.com/api/1/privileges/%s`,
			"DELETE_PRIVILEGE_URL":                `https://api.%s.onelogin.com/api/1/privileges/%s`,
			"GET_ROLES_ASSIGNED_TO_PRIVILEGE_URL": `https://api.%s.onelogin.com/api/1/privileges/%s/roles`,
			"ASSIGN_ROLES_TO_PRIVILEGE_URL":       `https://api.%s.onelogin.com/api/1/privileges/%s/roles`,
			"REMOVE_ROLE_FROM_PRIVILEGE_URL":      `https://api.%s.onelogin.com/api/1/privileges/%s/roles/%s`,
			"GET_USERS_ASSIGNED_TO_PRIVILEGE_URL": `https://api.%s.onelogin.com/api/1/privileges/%s/users`,
			"ASSIGN_USERS_TO_PRIVILEGE_URL":       `https://api.%s.onelogin.com/api/1/privileges/%s/users`,
			"REMOVE_USER_FROM_PRIVILEGE_URL":      `https://api.%s.onelogin.com/api/1/privileges/%s/users/%s`,
		},
	}

	// ValidActions is a list of valid actions that can be performed against API URLs
	ValidActions = []string{
		"apps:List",
		"apps:Get",
		"apps:Create",
		"apps:Update",
		"apps:Delete",
		"apps:ManageRoles",
		"apps:ManageUsers",
		"directories:List",
		"directories:Get",
		"directories:Create",
		"directories:Update",
		"directories:Delete",
		"directories:SyncUsers",
		"directories:RefreshSchema",
		"events:List",
		"events:Get",
		"mappings:List",
		"mappings:Get",
		"mappings:Create",
		"mappings:Update",
		"mappings:Delete",
		"mappings:ReapplyAll",
		"policies:List",
		"policies:user:Get",
		"policies:user:Create",
		"policies:user:Update",
		"policies:user:Delete",
		"policies:app:Get",
		"policies:app:Create",
		"policies:app:Update",
		"policies:app:Delete",
		"privileges:List",
		"privileges:Get",
		"privileges:Create",
		"privileges:Update",
		"privileges:Delete",
		"privileges:ListUsers",
		"privileges:ListRoles",
		"privileges:ManageUsers",
		"privileges:ManageRoles",
		"reports:List",
		"reports:Get",
		"reports:Create",
		"reports:Update",
		"reports:Delete",
		"reports:Run",
		"roles:List",
		"roles:Get",
		"roles:Create",
		"roles:Update",
		"roles:Delete",
		"roles:ManageUsers",
		"roles:ManageApps",
		"trustedidp:List",
		"trustedidp:Get",
		"trustedidp:Create",
		"trustedidp:Update",
		"trustedidp:Delete",
		"users:List",
		"users:Get",
		"users:Create",
		"users:Update",
		"users:Delete",
		"users:Unlock",
		"users:ResetPassword",
		"users:ForceLogout",
		"users:Invite",
		"users:ReapplyMappings",
		"users:ManageRoles",
		"users:ManageApps",
		"users:GenerateTempMfaToken",
	}
)
