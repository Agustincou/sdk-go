package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(
		accessToken,
		config.WithCorporationID("1yuy811998tt11199"),
		config.WithIntegratorID("6888999999000001"),
		config.WithPlatformID("prd_02647ea11edb6888682a831752a"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := paymentmethod.NewClient(cfg)
	paymentMethods, err := client.List(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range paymentMethods {
		fmt.Println(v)
	}
}