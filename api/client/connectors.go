package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/j4ng5y/onelogin-go/api"
)

var ConnectorsAuthMethod = map[string]string{
	"Password": "0",
	"OpenID": "1",
	"SAML": "2",
	"API": "3",
	"Google": "4",
	"FormsBasedApp": "6",
	"WSFED": "7",
	"OpenIDConnect": "8",
}

type ListConnectorsRequest struct {
	// BearerToken is your bearer token
	BearerToken string
	// Name is the full or partial name of the app to search for
	Name string
	// AuthMethod will return all connectors of a given type
	//
	// Use client.ConnectorsAuthMethod["friendly name of method"] to use the appropriate ID
	AuthMethod string
}

func (L *ListConnectorsRequest) Marshal() ([]byte, error) {
	return json.Marshal(L)
}

type ListConnectorsResponse []struct {
	ID int `json:"id"`
	Name string `json:"name"`
	AuthMethod int `json:"auth_method"`
	AllowsNewParameters bool `json:"allows_new_parameters"`
	IconURL string `json:"icon_url"`
}

func (L *ListConnectorsResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, L)
}

func (C *Client) ListConnectors(req *ListConnectorsRequest) (*ListConnectorsResponse, error) {
	var Resp = &ListConnectorsResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		BaseURL: api.URLS["OAUTH2_TOKEN_URLS"]["GET_RATE_URL"],
		QueryParameters: map[string]string{
			"name": req.Name,
			"auth_method": req.AuthMethod,
		},
	}
	URL, err := api.URLBuilder(builderOpts)
	if err != nil {
		return nil, fmt.Errorf("error building URL: %v", err)
	}

	httpReq, err := C.RequestBuilder(&RequestOptions{
		Method: http.MethodGet,
		URL: URL,
		Bearer: true,
		AccessToken: req.BearerToken,

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