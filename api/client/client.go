package client

import "C"
import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/j4ng5y/onelogin-go/api/session"
)

type ErrorResponse struct {
	Error bool `json:"error"`
	Code int `json:"code"`
	Type string `json:"type"`
	Message string `json:"message"`
}

type Client struct {
	Session *session.Session
	Options *Options
	HTTPClient *http.Client
}

type GETRequestOptions struct {
	Bearer bool
	AccessToken string
	URL string
	CustomAllowedOriginHeader string
}

type POSTRequestOptions struct {
	AccessToken string
	URL string
	Body []byte
	CustomAllowedOriginHeader string
}

func (C *Client) AuthenticatedGETRequestBuilder(options *GETRequestOptions) (*http.Request, error) {
	if !options.Bearer {
		return nil, fmt.Errorf("GETRequestOptions.Bearer is set to 'false', consider using Client.UnauthenticatedGETRequestBuilder instead")
	} else {
		if options.AccessToken == "" {
			return nil, fmt.Errorf("GETRequestOptions.AccessToken is required when GETRequestOptions.Bearer is set to 'true'")
		}
	}

	req, err := http.NewRequest(http.MethodGet, options.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request due to error: %+v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("bearer:%s", options.AccessToken))

	if options.CustomAllowedOriginHeader != "" {
		req.Header.Set("Custom-Allowed-Origin-Header-1", options.CustomAllowedOriginHeader)
	}

	return req, err
}

func (C *Client) UnauthenticatedGETRequestBuilder(options *GETRequestOptions) (*http.Request, error) {
	if options.Bearer {
		return nil, fmt.Errorf("GETRequestOptions.Bearer is set to 'true', consider using Client.AuthenticatedGETRequestBuilder instead")
	} else {
		if options.AccessToken != "" {
			return nil, fmt.Errorf("GETRequestOptions.AccessToken is not empty, consider using Client.AuthenticatedGETRequestBuilder instead")
		}
	}

	req, err := http.NewRequest(http.MethodGet, options.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request due to error: %+v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	if options.CustomAllowedOriginHeader != "" {
		req.Header.Set("Custom-Allowed-Origin-Header-1", options.CustomAllowedOriginHeader)
	}

	return req, err
}

func (C *Client) AuthenticatedPOSTRequestBuilder(options *POSTRequestOptions) (*http.Request, error) {
	if options.AccessToken == "" {
		return nil, fmt.Errorf("POSTRequestOptions.AccessToken is required")
	}

	req, err := http.NewRequest(http.MethodPost, options.URL, bytes.NewBuffer(options.Body))
	if err != nil {
		return nil, fmt.Errorf("error creating request due to error: %+v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("bearer:%s", options.AccessToken))

	if options.CustomAllowedOriginHeader != "" {
		req.Header.Set("Custom-Allowed-Origin-Header-1", options.CustomAllowedOriginHeader)
	}

	return req, err
}

type Options struct {
	SubDomain string
	MaxResults     int
	DefaultTimeout time.Duration
}

func NewWithOptions(sess *session.Session, opts *Options) (*Client, error) {
	switch {
	case opts.SubDomain == "":
		return nil, fmt.Errorf("Options.SubDomain must not be blank")
	case opts.MaxResults == 0:
		C.Options.MaxResults = 1000
	case opts.DefaultTimeout == time.Duration(0):
		C.Options.DefaultTimeout = time.Duration(60) * time.Second
	}

	return &Client{
		Session: sess,
		Options: opts,
		HTTPClient: http.DefaultClient,
	}, nil
}

func New(sess *session.Session, subDomain string) (*Client, error) {
	switch {
	case subDomain == "":
		return nil, fmt.Errorf("subDomain must not be blank")
	}
	return &Client{
		Session: sess,
		HTTPClient: http.DefaultClient,
		Options: &Options{
			SubDomain: subDomain,
			MaxResults:     1000,
			DefaultTimeout: time.Duration(60) * time.Second,
		},
	}, nil
}
