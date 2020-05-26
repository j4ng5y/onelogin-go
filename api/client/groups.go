package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/j4ng5y/onelogin-go/api"
)

type GetGroupsRequest struct {
	BearerToken string
	QueryParameters struct {
		ID string
	}
}

type GetGroupsResponse struct {
	Status ErrorResponse `json:"status"`
	Pagination struct {
		BeforeCursor string `json:"before_cursor"`
		AfterCursor string `json:"after_cursor"`
		PreviousLink string `json:"previous_link"`
		NextLink string `json:"next_link"`
	} `json:"pagination"`
	Data []struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Reference string `json:"reference"`
	} `json:"data"`
}

func (G *GetGroupsResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

type GetGroupRequest struct {
	BearerToken string
	GroupID string
}
type GetGroupResponse struct {
	Status ErrorResponse `json:"status"`
	Data []struct {
		ID int `json:"id"`
		Name string `json:"name"`
		Reference string `json:"reference"`
	} `json:"data"`
}

func (G *GetGroupResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

func (C *Client) GetGroups(req *GetGroupsRequest) (*GetGroupsResponse, error) {
	var Resp = &GetGroupsResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		BaseURL: api.URLS["GROUP_URLS"]["GET_GROUPS_URL"],
		QueryParameters: map[string]string{
			"id": req.QueryParameters.ID,
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

func (C *Client) GetGroup(req *GetGroupRequest) (*GetGroupResponse, error) {
	var Resp = &GetGroupResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		ObjectID: req.GroupID,
		BaseURL: api.URLS["GROUP_URLS"]["GET_GROUP_URL"],
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