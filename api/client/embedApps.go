package client

import "C"
import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/j4ng5y/onelogin-go/api"
)

type GetEmbedAppsRequest struct {
	// EmbeddingToken is obtained via OneLogin.com > Settings > Embedding
	EmbeddingToken string
	Email string
}

type GetEmbedAppsResponse struct {
	Apps []struct {
		App struct {
			ID int `xml:"id"`
			Icon string `xml:"icon"`
			Name string `xml:"name"`
			Provisioned int `xml:"provisioned"`
			ExtensionRequired bool `xml:"extension_required"`
			Personal bool `xml:"personal"`
			LoginID int `xml:"login_id"`
		} `xml:"app"`
	} `xml:"apps"`
}

func (G *GetEmbedAppsResponse) Unmarshal(htmlBody io.ReadCloser) error {
	body, err := ioutil.ReadAll(htmlBody)
	if err != nil {
		return err
	}

	return xml.Unmarshal(body, G)
}

func (C *Client) GetEmbedApps(req *GetEmbedAppsRequest) (*GetEmbedAppsResponse, error) {
	var Resp = &GetEmbedAppsResponse{}
	builderOpts := &api.URLBuilderOptions{
		SubDomain: C.Options.SubDomain,
		BaseURL: api.URLS["EMBED_APPS_URL"]["EMBED_APP_URL"],
		QueryParameters: map[string]string{
			"token": req.EmbeddingToken,
			"email": req.Email,
		},
	}
	URL, err := api.URLBuilder(builderOpts)
	if err != nil {
		return nil, fmt.Errorf("error building URL: %v", err)
	}

	httpReq, err := C.RequestBuilder(&RequestOptions{
		Method: http.MethodGet,
		URL: URL,
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