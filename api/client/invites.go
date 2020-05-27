package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/j4ng5y/onelogin-go/api"
)

type GenerateInviteLinkRequest struct {
	BearerToken string
	Email string `json:"email"`
}

func (G *GenerateInviteLinkRequest) Marshal() ([]byte, error) {
	return json.Marshal(G)
}

type GenerateInviteLinkResponse struct {
	Status ErrorResponse `json:"status"`
	Data []string `json:"data"`
}

func (G *GenerateInviteLinkResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

type SendInviteLinkRequest struct {
	BearerToken string
	Email string `json:"email"`
	PersonalEmail string `json:"personal_email"`
}

func (S *SendInviteLinkRequest) Marshal() ([]byte, error) {
	return json.Marshal(S)
}

type SendInviteLinkResponse struct {
	Status ErrorResponse `json:"status"`
}

func (S *SendInviteLinkResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, S)
}

func (C *Client) GenerateInviteLink(req *GenerateInviteLinkRequest) (*GenerateInviteLinkResponse, error) {
	var Resp = &GenerateInviteLinkResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		BaseURL: api.URLS["INVITE_LINK_URLS"]["GENERATE_INVITE_LINK_URL"],
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

func (C *Client) SendInviteLink(req *SendInviteLinkRequest) (*SendInviteLinkResponse, error) {
	var Resp = &SendInviteLinkResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		BaseURL: api.URLS["INVITE_LINK_URLS"]["SEND_INVITE_LINK_URL"],
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