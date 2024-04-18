package salepoint

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const baseURL = "https://api.mercadopago.com/pos"

func buildSearchURL(request SearchRequest) string {
	return baseURL + fmt.Sprintf("?%s", generateQueryParams(request))
}

func buildGetURL(salePointID string) string {
	return baseURL + fmt.Sprintf("/%s", url.QueryEscape(salePointID))
}

func generateQueryParams(request SearchRequest) string {
	queryParams := url.Values{}
	v := reflect.ValueOf(request)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Tag.Get("json")
		fieldValue := v.Field(i).Interface()

		// Si el valor del campo no es cero o cadena vacía, lo agregamos como parámetro de consulta
		if !reflect.DeepEqual(fieldValue, reflect.Zero(v.Field(i).Type()).Interface()) && fieldValue != "" {
			switch fieldValue := fieldValue.(type) {
			case bool:
				queryParams.Add(fieldName, url.QueryEscape(strconv.FormatBool(fieldValue)))
			case int:
				queryParams.Add(fieldName, url.QueryEscape(strconv.Itoa(fieldValue)))
			case string:
				queryParams.Add(fieldName, url.QueryEscape(fieldValue))
			}
		}
	}

	return queryParams.Encode()
}

// Client contains the method to interact with sale points API.
type Client interface {
	Create(ctx context.Context, request CreateRequest) (*CreateResponse, error)
	Search(ctx context.Context, request SearchRequest) (*SearchResponse, error)
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

func (c *client) Search(ctx context.Context, request SearchRequest) (*SearchResponse, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    buildSearchURL(request),
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Get(ctx context.Context, salePointID string) (*SearchResponse, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    buildGetURL(salePointID),
	}
	resource, err := httpclient.DoRequest[*SearchResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
