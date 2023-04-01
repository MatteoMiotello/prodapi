package main

import (
	"context"
	"fmt"
	"github.com/MatteoMiotello/prodapi/graph"
	"github.com/MatteoMiotello/prodapi/internal/bootstrap"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitMongoDb()
}

func main() {
	ctx := context.Background()
	products, err := graph.GetProducts(ctx, 10, 0)

	if err != nil {
		panic(err)
	}

	fmt.Println(products)
}
