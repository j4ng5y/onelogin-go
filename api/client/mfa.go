package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/j4ng5y/onelogin-go/api"
)

type GetMFAFactorsRequest struct {
	BearerToken string
	UserID string
}
type GetMFAFactorsResponse struct {
	Status ErrorResponse `json:"status"`
	Data struct {
		AuthFactors []struct {
			Name string `json:"name"`
			FactorID int `json:"factor_id"`
		} `json:"auth_factors"`
	} `json:"data"`
}

func (G *GetMFAFactorsResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

type EnrollMFAFactorRequest struct {
	BearerToken string
	UserID string
	FactorID int `json:"factor_id"`
	DisplayName string `json:"display_name"`
	PhoneNumber string `json:"number"`
	Verified bool `json:"verified"`
}

func (E *EnrollMFAFactorRequest) Marshal() ([]byte, error) {
	return json.Marshal(E)
}

type EnrollMFAFactorResponse struct {
	Status ErrorResponse `json:"status"`
	Data []struct {
		Active bool `json:"active"`
		Default bool `json:"default"`
		StateToken string `json:"state_token"`
		AuthFactorName string `json:"auth_factor_name"`
		PhoneNumber string `json:"phone_number"`
		TypeDisplayName string `json:"type_display_name"`
		NeedsTrigger bool `json:"needs_trigger"`
		UserDisplayName string `json:"user_display_name"`
		ID int `json:"id"`
	} `json:"data"`
}

func (E *EnrollMFAFactorResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, E)
}

type GetEnrolledMFAFactorsRequest struct {}
type GetEnrolledMFAFactorsResponse struct {}

type ActivateMFAFactorRequest struct {}
type ActivateMFAFactorResponse struct {}

type VerifyMFAFactorRequest struct {}
type VerifyMFAFactorResponse struct {}

type RemoveMFAFactorRequest struct {}
type RemoveMFAFactorResponse struct {}

func (C *Client) GetMFAFactors(req *GetMFAFactorsRequest) (*GetMFAFactorsResponse, error) {
	var Resp = &GetMFAFactorsResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		BaseURL: api.URLS["MFA_URLS"]["GET_FACTORS_URL"],
		ObjectID: req.UserID,
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

func (C *Client) EnrollMFAFactor(req *EnrollMFAFactorRequest) (*EnrollMFAFactorResponse, error) {
	var Resp = &EnrollMFAFactorResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		BaseURL: api.URLS["MFA_URLS"]["ENROLL_FACTOR_URL"],
		ObjectID: req.UserID,
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

func (C *Client) GetEnrolledMFAFactors(req *GetEnrolledMFAFactorsRequest) (*GetEnrolledMFAFactorsResponse, error) {
	var Resp = &GetEnrolledMFAFactorsResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		BaseURL: api.URLS["MFA_URLS"]["GET_FACTORS_URL"],
		ObjectID: req.UserID,
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