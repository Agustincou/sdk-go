package qr

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const baseURLTemplate = "https://api.mercadopago.com/instore/orders/qr/seller/collectors/%d/pos/%s/qrs"

func buildURL(userID int, externalSalePointID string) string {
	return fmt.Sprintf(baseURLTemplate, userID, url.QueryEscape(externalSalePointID))
}

// Client contains the method to interact with sale points API.
type Client interface {
	Create(ctx context.Context, userID int, externalSalePointID string, request CreateRequest) (*CreateResponse, error)
}

type client struct {
	cfg *config.Config
}

func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Create(ctx context.Context, userID int, externalSalePointID string, request CreateRequest) (*CreateResponse, error) {
	requestData := httpclient.RequestData{
		Body:   request,
		Method: http.MethodPut,
		URL:    buildURL(userID, externalSalePointID),
	}
	resource, err := httpclient.DoRequest[*CreateResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
