package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/j4ng5y/onelogin-go/api"
)

type GetSAMLAssertionRequest struct {
	BearerToken string
	UsernameOrEmail string `json:"username_or_email"`
	Password string `json:"password"`
	AppID string `json:"app_id"`
	Subdomain string `json:"subdomain"`
	IPAddress string `json:"ip_address"`
}

func (G *GetSAMLAssertionRequest) Marshal() ([]byte, error) {
	return json.Marshal(G)
}

type GetSAMLAssertionResponse struct {
	Data string `json:"data"`
	Status ErrorResponse `json:"status"`
}

func (G *GetSAMLAssertionResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

type GetSAMLAssertionWithMFARequest struct {
	BearerToken string
	AppID string `json:"app_id"`
	DeviceID string `json:"device_id"`
	StateToken string `json:"state_token"`
	OTPToken string `json:"otp_token"`
	DoNotNotify bool `json:"do_not_notify"`
}

func (G *GetSAMLAssertionWithMFARequest) Marshal() ([]byte, error) {
	return json.Marshal(G)
}

type GetSAMLAssertionWithMFAResponse struct {
	Data string `json:"data"`
	Status ErrorResponse `json:"status"`
}

func (G *GetSAMLAssertionWithMFAResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

func (C *Client) GetSAMLAssertion(req *GetSAMLAssertionRequest) (*GetSAMLAssertionResponse, error) {
	var Resp = &GetSAMLAssertionResponse{}
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

func (C *Client) GetSAMLAssertionWithMFA(req *GetSAMLAssertionWithMFARequest) (*GetSAMLAssertionWithMFAResponse, error) {
	var Resp = &GetSAMLAssertionWithMFAResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Options.Region,
		BaseURL: api.URLS["OAUTH2_TOKEN_URLS"]["GET_SAML_VERIFY_FACTOR"],
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