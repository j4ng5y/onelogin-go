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

type GetAccessTokenRequest struct {
	GrantType string `json:"grant_type"`
}

func (G *GetAccessTokenRequest) Marshal() ([]byte, error) {
	return json.Marshal(G)
}

type GetAccessTokenResponse struct {
	AccessCredentials string `json:"access_token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType string `json:"token_type"`
	AccountID int `json:"account_id"`
	Status ErrorResponse `json:"status"`
}

func (G *GetAccessTokenResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

func (C *Client) GetAccessToken(req *GetAccessTokenRequest) (*GetAccessTokenResponse, error) {
	var Resp = &GetAccessTokenResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Options.Region,
		BaseURL: api.URLS["OAUTH2_TOKEN_URLS"]["TOKEN_REQUEST_URL"],
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

type RegenerateTokenRequest struct {
	GrantType string `json:"grant_type"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (R *RegenerateTokenRequest) Marshal() ([]byte, error) {
	return json.Marshal(R)
}

type RegenerateTokenResponse struct {
	AccessToken string `json:"access_token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType string `json:"token_type"`
	Status ErrorResponse `json:"status"`
}

func (R *RegenerateTokenResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, R)
}

func (C *Client) RegenerateToken(req *RegenerateTokenRequest) (*RegenerateTokenResponse, error) {
	var Resp = &RegenerateTokenResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Options.Region,
		BaseURL: api.URLS["OAUTH2_TOKEN_URLS"]["TOKEN_REQUEST_URL"],
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

type RevokeTokenRequest struct {
	AccessToken string `json:"access_token"`
}

func (R *RevokeTokenRequest) Marshal() ([]byte, error) {
	return json.Marshal(R)
}

type RevokeTokenResponse struct {
	Status ErrorResponse `json:"status"`
}

func (R *RevokeTokenResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, R)
}

func (C *Client) RevokeToken(req *RevokeTokenRequest) (*RevokeTokenResponse, error) {
	var Resp = &RevokeTokenResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Options.Region,
		BaseURL: api.URLS["OAUTH2_TOKEN_URLS"]["TOKEN_REVOKE_URL"],
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

type GetRateLimitsRequest struct {
	BearerToken string
}

type GetRateLimitsResponse struct {
	Status ErrorResponse `json:"status"`
	Data struct {
		Limit int `json:"X-RateLimit-Limit"`
		Remaining int `json:"X-RateLimit-Remaining"`
		Reset int `json:"X-RateLimit-Reset"`
	} `json:"data"`
}

func (G *GetRateLimitsResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

func (C *Client) GetRateLimits(req *GetRateLimitsRequest) (*GetRateLimitsResponse, error) {
	var Resp = &GetRateLimitsResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Options.Region,
		BaseURL: api.URLS["OAUTH2_TOKEN_URLS"]["GET_RATE_URL"],
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