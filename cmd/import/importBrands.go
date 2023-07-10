package main

import (
	"github.com/MatteoMiotello/prodapi/internal/bootstrap"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitMongoDb()
}

//
//func main() {
//	defer nosql.Disconnect()
//
//	ctx := context.Background()
//
//	brandRes, err := graph.GetBrands(ctx)
//
//	if err != nil {
//		panic(err)
//	}
//
//	for _, brand := range brandRes.Brands {
//		b := new(schemas.Brand)
//
//		b.Name = brand.Name
//		b.Code = brand.Code
//		b.Incomplete = false
//		b.ImageIndex = 0
//
//		upsert := true
//		_, err := nosql.BrandCollection().ReplaceOne(
//			ctx,
//			bson.D{{"code", bson.D{{"$eq", brand.Code}}}},
//			b,
//			&options.ReplaceOptions{
//				Upsert: &upsert,
//			},
//		)
//
//		if err != nil {
//			panic(err)
//		}
//	}
//}
