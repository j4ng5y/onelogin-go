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

type GetEventTypesResponse struct {
	Status ErrorResponse `json:"status"`
	Data []struct {
		Name string `json:"name"`
		Description string `json:"description"`
		ID int `json:"id"`
	} `json:"data"`
}

func (G *GetEventTypesResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

type GetEventsRequest struct {
	BearerToken string
	QueryParameters struct {
		ClientID string
		CreatedAt string
		DirectoryID string
		EventTypeID string
		ID string
		Resolution string
		Since string
		Until string
		UserID string
	}
}

type GetEventsResponse struct {
	Status ErrorResponse `json:"status"`
	Pagination struct {
		BeforeCursor string `json:"before_cursor"`
		AfterCursor string `json:"after_cursor"`
		PreviousLink string `json:"previous_link"`
		NextLink string `json:"next_link"`
	} `json:"pagination"`
	Data []struct{
		ID int `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		AccountID int `json:"account_id"`
		UserID int `json:"user_id"`
		EventTypeID int `json:"event_type_id"`
		Notes string `json:"notes"`
		IPAddr string `json:"ipaddr"`
		ActorUserID int `json:"actor_user_id"`
		AssumingActorUserID int `json:"assuming_actor_user_id"`
		RoleID int `json:"role_id"`
		AppID int `json:"app_id"`
		GroupID int `json:"group_id"`
		OTPDeviceID int `json:"otp_device_id"`
		PolicyID int `json:"policy_id"`
		ActorSystem string `json:"actor_system"`
		CustomMessage string `json:"custom_message"`
		RoleName string `json:"role_name"`
		AppName string `json:"app_name"`
		GroupName string `json:"group_name"`
		ActorUserName string `json:"actor_user_name"`
		UserName string `json:"user_name"`
		PolicyName string `json:"policy_name"`
		OTPDeviceName string `json:"otp_device_name"`
		OperationName string `json:"operation_name"`
		DirectorySyncRunID int `json:"directory_sync_run_id"`
		DirectoryID int `json:"directory_id"`
		Resolution string `json:"resolution"`
		ClientID int `json:"client_id"`
		ResourceTypeID int `json:"resource_type_id"`
		ErrorDescription string `json:"error_description"`
		ProxyIP string `json:"proxy_ip"`
		RiskScore int `json:"risk_score"`
		RiskReasons string `json:"risk_reasons"`
		RiskCookieID int `json:"risk_cookie_id"`
		BrowserFingerprint string `json:"browser_fingerprint"`
	} `json:"data"`
}

func (G *GetEventsResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

type GetEventRequest struct {
	BearerToken string
	EventID string
}
type GetEventResponse struct {
	Status ErrorResponse `json:"status"`
	Data []struct{
		ID int `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		AccountID int `json:"account_id"`
		UserID int `json:"user_id"`
		EventTypeID int `json:"event_type_id"`
		Notes string `json:"notes"`
		IPAddr string `json:"ipaddr"`
		ActorUserID int `json:"actor_user_id"`
		AssumingActorUserID int `json:"assuming_actor_user_id"`
		RoleID int `json:"role_id"`
		AppID int `json:"app_id"`
		GroupID int `json:"group_id"`
		OTPDeviceID int `json:"otp_device_id"`
		PolicyID int `json:"policy_id"`
		ActorSystem string `json:"actor_system"`
		CustomMessage string `json:"custom_message"`
		RoleName string `json:"role_name"`
		AppName string `json:"app_name"`
		GroupName string `json:"group_name"`
		ActorUserName string `json:"actor_user_name"`
		UserName string `json:"user_name"`
		PolicyName string `json:"policy_name"`
		OTPDeviceName string `json:"otp_device_name"`
		OperationName string `json:"operation_name"`
		DirectorySyncRunID int `json:"directory_sync_run_id"`
		DirectoryID int `json:"directory_id"`
		Resolution string `json:"resolution"`
		ClientID int `json:"client_id"`
		ResourceTypeID int `json:"resource_type_id"`
		ErrorDescription string `json:"error_description"`
		ProxyIP string `json:"proxy_ip"`
		RiskScore int `json:"risk_score"`
		RiskReasons string `json:"risk_reasons"`
		RiskCookieID int `json:"risk_cookie_id"`
		BrowserFingerprint string `json:"browser_fingerprint"`
	} `json:"data"`
}

func (G *GetEventResponse) Unmarshal(httpBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, G)
}

type CreateEventRequest struct {}
type CreateEventResponse struct {}

func (C *Client) GetEventTypes() (*GetEventTypesResponse, error) {
	var Resp = &GetEventTypesResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		BaseURL: api.URLS["EVENT_URLS"]["GET_EVENT_TYPES_URL"],
	}
	URL, err := api.URLBuilder(builderOpts)
	if err != nil {
		return nil, fmt.Errorf("error building URL: %v", err)
	}

	resp, err := http.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("error bulding request: %v", err)
	}

	if err := Resp.Unmarshal(resp.Body); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return Resp, nil
}

func (C *Client) GetEvents(req *GetEventsRequest) (*GetEventsResponse, error) {
	var Resp = &GetEventsResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		BaseURL: api.URLS["EVENT_URLS"]["GET_EVENTS_URL"],
		QueryParameters: map[string]string{
			"client_id": req.QueryParameters.ClientID,
			"created_at": req.QueryParameters.CreatedAt,
			"directory_id": req.QueryParameters.DirectoryID,
			"event_type_id": req.QueryParameters.EventTypeID,
			"id": req.QueryParameters.ID,
			"resolution": req.QueryParameters.Resolution,
			"since": req.QueryParameters.Since,
			"until": req.QueryParameters.Until,
			"user_id": req.QueryParameters.UserID,
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

func (C *Client) GetEvent(req *GetEventRequest) (*GetEventResponse, error) {
	var Resp = &GetEventResponse{}
	builderOpts := &api.URLBuilderOptions{
		Region: C.Session.Region,
		ObjectID: req.EventID,
		BaseURL: api.URLS["EVENT_URLS"]["GET_EVENT_URL"],
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

func (C *Client) CreateEvent(req *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, fmt.Errorf("unimplimented")
}