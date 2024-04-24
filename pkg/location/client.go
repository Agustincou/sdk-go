package location

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const stateDetailsTemplateURL = "https://api.mercadolibre.com/classified_locations/states/%s"
const zipCodeTemplateURL = "https://api.mercadolibre.com/countries/%s/zip_codes/%s"
const cityTemplateURL = "https://api.mercadolibre.com/classified_locations/cities/%s"

func buildGetStateDetailsURL(stateID string) string {
	return fmt.Sprintf(stateDetailsTemplateURL, url.QueryEscape(stateID))
}

func buildGetStateByZipCodeURL(countryID, zipCode string) string {
	return fmt.Sprintf(zipCodeTemplateURL, url.QueryEscape(countryID), url.QueryEscape(zipCode))
}

func buildGetCitiesDetailsURL(cityID string) string {
	return fmt.Sprintf(cityTemplateURL, url.QueryEscape(cityID))
}

// Client contains the method to interact with sale points API.
type Client interface {
	GetStateByZipCode(ctx context.Context, countryID, zipCode string) (*StateByZipCodeResponse, error)
	GetStateDetails(ctx context.Context, stateID string) (*StateDetailsResponse, error)
	GetCityDetails(ctx context.Context, cityID string) (*CityDetailsResponse, error)
}

type client struct {
	cfg *config.Config
}

func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) GetStateByZipCode(ctx context.Context, countryID, zipCode string) (*StateByZipCodeResponse, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    buildGetStateByZipCodeURL(countryID, zipCode),
	}
	resource, err := httpclient.DoRequest[*StateByZipCodeResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) GetStateDetails(ctx context.Context, stateID string) (*StateDetailsResponse, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    buildGetStateDetailsURL(stateID),
	}
	resource, err := httpclient.DoRequest[*StateDetailsResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) GetCityDetails(ctx context.Context, cityID string) (*CityDetailsResponse, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    buildGetCitiesDetailsURL(cityID),
	}
	resource, err := httpclient.DoRequest[*CityDetailsResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
