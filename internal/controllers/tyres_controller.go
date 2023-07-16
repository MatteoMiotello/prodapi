package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/MatteoMiotello/prodapi/internal/fs_handlers"
	"github.com/MatteoMiotello/prodapi/internal/nosql"
	"github.com/MatteoMiotello/prodapi/schemas"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type tyreController struct {
}

func NewTyreController() *tyreController {
	return &tyreController{}
}

func (receiver *tyreController) findTyreByProductCode(ctx context.Context, productCode string) (error, *schemas.Tyre) {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"code", bson.D{{"$eq", "CONTINENTAL"}}}},
				bson.D{{"incomplete", bson.D{{"$ne", true}}}},
			},
		},
	}

	var tyre schemas.Tyre

	found := nosql.TyreCollection().FindOne(ctx, filter)

	err := found.Decode(&tyre)
	if err != nil {
		return errors.New("could not decode record: " + err.Error()), nil
	}

	return nil, &tyre
}

func (receiver *tyreController) FindByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productCode, ok := vars["tyre_code"]

	if !ok {
		http.NotFound(w, r)
		return
	}

	err, tyre := receiver.findTyreByProductCode(r.Context(), productCode)

	if err != nil {
		http.Error(w, "could not decode record: "+err.Error(), http.StatusBadRequest)
		return
	}

	imageUrl := fs_handlers.
		NewImagesHandler(viper.GetString("APPLICATION_URL")).
		GetPublicUrl(tyre.Code + tyre.ImageExtension)

	tyre.ImagePath = imageUrl

	err = json.NewEncoder(w).Encode(tyre)
	if err != nil {
		http.Error(w, "Failed to encode tyre: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (receiver *tyreController) FindImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productCode, ok := vars["tyre_code"]

	if !ok {
		http.NotFound(w, r)
		return
	}

	err, tyre := receiver.findTyreByProductCode(r.Context(), productCode)

	if err != nil {
		http.Error(w, "could not decode record: "+err.Error(), http.StatusNotFound)
		return
	}

	filePath := tyre.ImagePath
	filePath = "." + filePath

	http.ServeFile(w, r, filePath)
}
