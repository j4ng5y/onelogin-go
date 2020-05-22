package client

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Error bool `json:"error"`
	Code int `json:"code"`
	Type string `json:"type"`
	Message string `json:"message"`
}

type Client struct {
	Options *Options
	HTTPClient *http.Client
}

type RequestOptions struct {
	Bearer bool
	AccessToken string
	Method string
	URL string
	Body []byte
	CustomAllowedOriginHeader string
}

func (C *Client) RequestBuilder(options *RequestOptions) (req *http.Request, err error) {
	switch options.Method {
	case http.MethodGet:
		req, err = http.NewRequest(options.Method, options.URL, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %v", err)
		}
	case http.MethodPost:
		req, err = http.NewRequest(options.Method, options.URL, bytes.NewBuffer(options.Body))
		if err != nil {
			return nil, fmt.Errorf("error creating request: %v", err)
		}
	default:
		return nil, fmt.Errorf("unsuppored method %s", options.Method)
	}

	req.Header.Set("Content-Type", "application/json")

	switch options.Bearer {
	case true:
		if options.AccessToken == "" {
			return nil, fmt.Errorf("RequestOptions.AccessToken can not be blank")
		}
		req.Header.Set("Authorization", fmt.Sprintf("bearer:%s", options.AccessToken))
	case false:
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", C.Options.ClientID, C.Options.ClientSecret)))))
	}

	if options.CustomAllowedOriginHeader != "" {
		req.Header.Set("Custom-Allowed-Origin-Header-1", options.CustomAllowedOriginHeader)
	}

	return req, nil
}

type Options struct {
	ClientID       string
	ClientSecret   string
	Region         string
	SubDomain string
	MaxResults     int
	DefaultTimeout time.Duration
}

func NewWithOptions(opts *Options) (*Client, error) {
	C := &Client{}

	switch {
	case opts.ClientID == "":
		return nil, fmt.Errorf("Options.ClientID must not be blank")
	case opts.ClientSecret == "":
		return nil, fmt.Errorf("Options.ClientSecret must not be blank")
	case opts.Region == "":
		return nil, fmt.Errorf("Options.Region must not be blank")
	case opts.SubDomain == "":
		return nil, fmt.Errorf("Options.SubDomain must not be blank")
	case opts.MaxResults == 0:
		C.Options.MaxResults = 1000
	case opts.DefaultTimeout == time.Duration(0):
		C.Options.DefaultTimeout = time.Duration(60) * time.Second
	}

	C.Options = opts
	C.HTTPClient = http.DefaultClient
	return C, nil
}

func New(clientID, clientSecret, subDomain string) (*Client, error) {
	switch {
	case clientID == "":
		return nil, fmt.Errorf("clientID must not be blank")
	case clientSecret == "":
		return nil, fmt.Errorf("clientSecret must not be blank")
	case subDomain == "":
		return nil, fmt.Errorf("subDomain must not be blank")
	}
	return &Client{
		HTTPClient: http.DefaultClient,
		Options: &Options{
			ClientID:       clientID,
			ClientSecret:   clientSecret,
			Region:         "us",
			SubDomain: subDomain,
			MaxResults:     1000,
			DefaultTimeout: time.Duration(60) * time.Second,
		},
	}, nil
}
