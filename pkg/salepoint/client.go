package salepoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const baseURL = "https://api.mercadopago.com/pos"

func buildSearchURL(externalSalePointID string) string {
	if externalSalePointID != "" {
		return baseURL + fmt.Sprintf("?external_id=%s", externalSalePointID)
	}

	return baseURL
}

// Client contains the method to interact with sale points API.
type Client interface {
	Create(ctx context.Context, request CreateRequest) (*CreateResponse, error)
	Get(ctx context.Context, externalSalePointID string) (*SearchResponse, error)
}

type client struct {
	cfg *config.Config
}

func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Create(ctx context.Context, request CreateRequest) (*CreateResponse, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPost,
		URL:    baseURL,
	}
	resource, err := httpclient.DoRequest[*CreateResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Get(ctx context.Context, externalSalePointID string) (*SearchResponse, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    buildSearchURL(externalSalePointID),
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
