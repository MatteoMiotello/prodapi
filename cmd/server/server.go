package main

import (
	"fmt"
	"github.com/MatteoMiotello/prodapi/internal/bootstrap"
	"github.com/MatteoMiotello/prodapi/internal/clients"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitMongoDb()
}

func main() {
	bClient := clients.NewBingClient()

	res, err := bClient.SearchTyreImage("265/45 R20 TL 108Y AZ-850 XL")

	fmt.Println(res.Value[0].ThumbnailUrl)

	if err != nil {
		panic(err)
	}
}
