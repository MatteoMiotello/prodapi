package controllers

import (
	"context"
	"errors"
	"github.com/MatteoMiotello/prodapi/internal/nosql"
	"github.com/MatteoMiotello/prodapi/schemas"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type BrandController struct {
}

func NewBrandController() *BrandController {
	return &BrandController{}
}

func (b BrandController) findBrandByProduct(ctx context.Context, brandCode string) (*schemas.Brand, error) {
	filter := bson.D{{
		"code", bson.D{{
			"$eq", brandCode,
		}},
	}}

	brand := new(schemas.Brand)

	found := nosql.BrandCollection().FindOne(ctx, filter)

	err := found.Decode(brand)
	if err != nil {
		return nil, errors.New("could not decode record: " + err.Error())
	}

	return brand, nil
}

func (b BrandController) FindImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productCode, ok := vars["brand_code"]

	if !ok {
		http.NotFound(w, r)
		return
	}

	brand, err := b.findBrandByProduct(r.Context(), productCode)

	if err != nil {
		http.Error(w, "could not decode record: "+err.Error(), http.StatusNotFound)
		return
	}

	filePath := brand.ImagePath
	filePath = "." + filePath

	http.ServeFile(w, r, filePath)
}
