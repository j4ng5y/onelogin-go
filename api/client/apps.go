package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/j4ng5y/onelogin-go/api"
)

type GetAppsRequest struct {
	BearerToken string
	//ID is the ID of the app that you want to return
	ID int `json:"id"`
}

func (G *GetAppsRequest) Marshal() ([]byte, error) {
	return json.Marshal(G)
}

type GetAppsResponse struct {
	ID int `json:"id"`
	Name string `json:"string"`
	Visible bool `json:"visible"`
	Description string `json:"description"`
	Notes string `json:"notes"`
	IconURL string `json:"icon_url"`
	AuthMethod int `json:"auth_method"`
	PolicyID int `json:"policy_id"`
	Provisioning struct{
		Enabled bool `json:"enabled"`
	} `json:"provisioning"`
	AllowAssumedSignIn bool `json:"allow_assumed_signin"`
	TabID int `json:"tab_id"`
	ConnectorID int `json:"connector_id"`
	SSO struct{
		MetaDataURL string `json:"metadata_url"`
		ACSURL string `json:"acs_url"`
		SLSURL string `json:"sls_url"`
		Issuer string `json:"issuer"`
		Certificate struct {
			Value string `json:"value"`
			ID int `json:"id"`
			Name string `json:"name"`
		} `json:"certificate"`
	} `json:"sso"`
	Configuration struct{
		ProviderARN string `json:"provider_arn"`
		SignatureAlgorithm string `json:"signature_algorithm"`
	} `json:"configuration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RoleIDs []int `json:"role_ids"`
	Parameters map[string]map[string]interface{} `json:"parameters"`
}

func (G *GetAppsResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

func (C *Client) GetApps(req *GetAppsRequest) (*GetAppsResponse, error) {
	var Resp = &GetAppsResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Options.Region,
		BaseURL: api.URLS["APPS_URLS"]["GET_APPS_URLS"],
	}
	URL, err := api.URLBuilder(builderOpts)
	if err != nil {
		return nil, fmt.Errorf("error building URL: %v", err)
	}
	b, err := req.Marshal()
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}
	httpReq, err := C.RequestBuilder(&RequestOptions{
		Method: http.MethodPost,
		URL: URL,
		Body: b,
	})
	if err != nil {
		return nil, fmt.Errorf("error bulding request: %v", err)
	}

	resp, err := C.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	if err := Resp.Unmarshal(resp.Body); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return Resp, nil
}