package store

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
	"github.com/mercadopago/sdk-go/pkg/mperror"
)

const templateURL = "https://api.mercadopago.com/users/%d/stores"
const getURL = "https://api.mercadopago.com/stores/%s"

func buildCreateURL(userID int) string {
	return fmt.Sprintf(templateURL, userID)
}

func buildSearchURL(userID int, externalStoreID string) string {
	if externalStoreID != "" {
		return fmt.Sprintf(templateURL, userID) + fmt.Sprintf("/search?external_id=%s", url.QueryEscape(externalStoreID))
	}

	return fmt.Sprintf(templateURL, userID)
}

func buildGetURL(storeID string) string {
	return fmt.Sprintf(getURL, storeID)
}

type Client interface {
	Create(ctx context.Context, userID int, request CreateRequest) (*CreateResponse, error)
	Search(ctx context.Context, userID int, externalStoreID string) (*SearchResponse, error)
	Get(ctx context.Context, storeID string) (*SearchResponse, error)
}

type client struct {
	cfg *config.Config
}

func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Create(ctx context.Context, userID int, request CreateRequest) (*CreateResponse, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    buildCreateURL(userID),
	}
	resource, err := httpclient.DoRequest[*CreateResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Search(ctx context.Context, userID int, externalStoreID string) (*SearchResponse, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    buildSearchURL(userID, externalStoreID),
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		if mpError, isMPError := err.(*mperror.ResponseError); isMPError {
			searchError := SearchErrorResponse{}
			if unmErr := json.Unmarshal([]byte(mpError.Message), &searchError); unmErr == nil {
				if searchError.Error == "store_not_found" {
					return &SearchResponse{Paging: Paging{}, Results: []SearchResponseResult{}}, nil // Recover for 404 Store Not Found
				}
			}
		}

		return nil, err
	}

	return resource, nil
}

func (c *client) Get(ctx context.Context, storeID string) (*SearchResponse, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    buildGetURL(storeID),
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
