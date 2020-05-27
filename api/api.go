package api

import (
	"fmt"
	"net/url"

	"github.com/j4ng5y/onelogin-go/api/session"
)

type URLBuilderOptions struct {
	Region session.Region
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
