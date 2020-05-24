package session

import (
	"fmt"
	"strings"
)

type ClientID string

func (C ClientID) Validate() (ClientID, error) {
	switch strings.ToLower(string(C)) {
	case "":
		return "", fmt.Errorf("invalid value, ClientID must not be blank")
	default:
		return C, nil
	}
}

type ClientSecret string

func (C ClientSecret) Validate() (ClientSecret, error) {
	switch strings.ToLower(string(C)) {
	case "":
		return "", fmt.Errorf("invalid value, ClientSecret must not be blank")
	default:
		return C, nil
	}
}

type Region string

func (R Region) Validate() (Region, error) {
	switch strings.ToLower(string(R)) {
	case "us":
		return R, nil
	case "eu":
		return R, nil
	default:
		return "", fmt.Errorf("invalid value: Region must be one of US or EU")
	}
}

type Session struct {
	ClientID ClientID
	ClientSecret ClientSecret
	Region Region
}

func New(id ClientID, secret ClientSecret, region Region) (*Session, error) {
	validatedID, err := id.Validate()
	if err != nil {
		return nil, err
	}

	validatedSecret, err := secret.Validate()
	if err != nil {
		return nil, err
	}

	validatedRegion, err := region.Validate()
	if err != nil {
		return nil, err
	}

	return &Session{
		ClientID: validatedID,
		ClientSecret: validatedSecret,
		Region: validatedRegion,
	}, nil
}
