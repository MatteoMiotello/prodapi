package controllers

import (
	"encoding/json"
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

func (receiver *tyreController) FindByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productCode, ok := vars["tyre_code"]

	if !ok {
		http.NotFound(w, r)
		return
	}

	filter := bson.D{{
		"code", bson.D{{
			"$eq", productCode,
		}},
	}}

	var tyre schemas.Tyre

	found := nosql.TyreCollection().FindOne(r.Context(), filter)

	err := found.Decode(&tyre)
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
