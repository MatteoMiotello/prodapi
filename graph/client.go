package graph

import (
	"context"
	"github.com/Khan/genqlient/graphql"
	"net/http"
)

func getClient() graphql.Client {
	httpClient := http.Client{}

	return graphql.NewClient("http://localhost:8080/query", &httpClient)
}

func GetProducts(ctx context.Context, limit int, offset int) (*getProductsResponse, error) {
	return getProducts(ctx, getClient(), limit, offset)
}
