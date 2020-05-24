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

type CreateSessionLoginTokenRequest struct {
	CustomAllowedOriginHeader string
	BearerToken string
	UsernameOrEmail string `json:"username-or-email"`
	Password string `json:"password"`
	Subdomain string `json:"subdomain"`

	//Fields is a comma-separated list of user fields to return
	Fields string `json:"fields"`
}

func (C *CreateSessionLoginTokenRequest) Marshal() ([]byte, error) {
	return json.Marshal(C)
}

type CreateSessionLoginTokenResponse struct {
	Status ErrorResponse `json:"status"`
	Data []struct {
		Status string `json:"status"`
		User struct {
			Username string `json:"username"`
			Email string `json:"email"`
			FirstName string `json:"firstname"`
			ID int `json:"id"`
			LastName string `json:"lastname"`
		} `json:"user"`
		ReturnToURL string `json:"return_to_url"`
		ExpiresAt time.Time `json:"expires_at"`
		SessionToken string `json:"session_token"`
	} `json:"data"`
}

func (C *CreateSessionLoginTokenResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, C)
}

type GetSessionTokenWithMFARequest struct {
	CustomAllowedOriginHeader string
	BearerToken string
	DeviceID string `json:"device_id"`
	StateToken string `json:"state_token"`
	OTPToken string `json:"otp_token"`
	DoNotNotify bool `json:"do_not_notify"`
}

func (G *GetSessionTokenWithMFARequest) Marshal() ([]byte, error) {
	return json.Marshal(G)
}

type GetSessionTokenWithMFAResponse struct {
	Status ErrorResponse `json:"status"`
	Data []struct {
		ReturnToURL string `json:"return_to_url"`
		User struct{
			Username string `json:"username"`
			Email string `json:"email"`
			FirstName string `json:"firstname"`
			LastName string `json:"lastname"`
			ID int `json:"id"`
		} `json:"user"`
		Status string `json:"status"`
		SessionToken string `json:"session_token"`
		ExpiresAt time.Time `json:"expires_at"`
	} `json:"data"`
}

func (G *GetSessionTokenWithMFAResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

func (C *Client) CreateSessionLoginToken(req *CreateSessionLoginTokenRequest) (*CreateSessionLoginTokenResponse, error) {
	var Resp = &CreateSessionLoginTokenResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		BaseURL: api.URLS["CUSTOM_LOGIN_URLS"]["SESSION_LOGIN_TOKEN_URL"],
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
		CustomAllowedOriginHeader: req.CustomAllowedOriginHeader,
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

func (C *Client) GetSessionTokenWithMFA(req *GetSessionTokenWithMFARequest) (*GetSessionTokenWithMFAResponse, error) {
	var Resp = &GetSessionTokenWithMFAResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		BaseURL: api.URLS["CUSTOM_LOGIN_URLS"]["GET_TOKEN_VERIFY_FACTOR"],
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
		CustomAllowedOriginHeader: req.CustomAllowedOriginHeader,
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