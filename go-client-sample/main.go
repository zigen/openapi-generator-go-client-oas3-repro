package main

import (
	"context"
	"github.com/zigen/go-sample-api"
)

func main() {
	println("hello world")

	cfg := go_sample_api.NewConfiguration()
	client := go_sample_api.NewAPIClient(cfg)
	ctx := context.Background()

	client.DefaultApi.GetState(ctx)
}