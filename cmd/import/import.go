package main

import (
	"context"
	"github.com/MatteoMiotello/prodapi/graph"
	"github.com/MatteoMiotello/prodapi/internal/bootstrap"
	"github.com/MatteoMiotello/prodapi/internal/constants"
	"github.com/MatteoMiotello/prodapi/internal/nosql"
	"github.com/MatteoMiotello/prodapi/schemas"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitMongoDb()
}

func main() {

	defer nosql.Disconnect()
	var offset int = 0
	var totals int = 1
	ctx := context.Background()

	for offset < totals {
		productRes, err := graph.GetProducts(ctx, 2, offset)

		if err != nil {
			panic(err)
		}

		totals = productRes.GetProducts().Pagination.Totals

		products := productRes.GetProducts().Products

		var pSlice []interface{}

		for _, product := range products {
			pSchema := new(schemas.Tyre)
			pSchema.Brand = product.Brand.Name
			pSchema.Code = product.Code

			for _, spec := range product.ProductSpecificationValues {
				switch constants.ProductSpecification(spec.Specification.Code) {
				case constants.TYRE_SPEC_NAME:
					pSchema.Name = spec.Value
					break
				case constants.TYRE_SPEC_REFERENCE:
					pSchema.Reference = spec.Value
					break
				case constants.TYRE_SPEC_WIDTH:
					pSchema.Width = spec.Value
					break
				case constants.TYRE_SPEC_ASPECT_RATIO:
					pSchema.AspectRatio = spec.Value
					break
				case constants.TYRE_SPEC_CONSTRUCTION:
					pSchema.Construction = spec.Value
					break
				case constants.TYRE_SPEC_RIM:
					pSchema.Rim = spec.Value
					break
				case constants.TYRE_SPEC_LOAD:
					pSchema.Load = spec.Value
					break
				case constants.TYRE_SPEC_SPEED:
					pSchema.Speed = spec.Value
					break
				case constants.TYRE_SPEC_SEASON:
					pSchema.Season = spec.Value
					break
				case constants.TYRE_SPEC_EPREL_ID:
					pSchema.EprelId = spec.Value
					break
				case constants.TYRE_SPEC_FUEL_EFFICIENCY:
					pSchema.FuelEfficiency = spec.Value
					break
				case constants.TYRE_SPEC_WET_GRIP_CLASS:
					pSchema.WetGripClass = spec.Value
					break
				case constants.TYRE_SPEC_EXTERNAL_ROLLING_NOISE_CLASS:
					pSchema.ExternalRollingNoiseClass = spec.Value
					break
				case constants.TYRE_SPEC_EXTERNAL_ROLLING_NOISE_LEVEL:
					pSchema.ExternalRollingNoiseLevel = spec.Value
					break
				case constants.TYRE_SPEC_LOAD_VERSION:
					pSchema.LoadVersion = spec.Value
					break
				}
			}

			pSlice = append(pSlice, *pSchema)
		}

		_, err = nosql.TyreCollection().InsertMany(ctx, pSlice)

		if err != nil {
			panic(err)
		}

		offset = offset + 2
	}
}
