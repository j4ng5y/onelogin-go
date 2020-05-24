package client

import "C"
import (
	"bytes"
	"encoding/base64"
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
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", C.Session.ClientID, C.Session.ClientSecret)))))
	}

	if options.CustomAllowedOriginHeader != "" {
		req.Header.Set("Custom-Allowed-Origin-Header-1", options.CustomAllowedOriginHeader)
	}

	return req, nil
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
