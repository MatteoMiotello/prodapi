package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MatteoMiotello/prodapi/internal/bootstrap"
	"github.com/MatteoMiotello/prodapi/internal/nosql"
	"github.com/MatteoMiotello/prodapi/schemas"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"net"
	"net/http"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.InitMongoDb()
}

func main() {
	ctx := context.Background()
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./images/tyres"))
	router.Handle("/resources/tyres/", http.StripPrefix("/resources/tyres/", fs))

	router.HandleFunc("/products/tyres/{tyre_code}", func(w http.ResponseWriter, r *http.Request) {
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

		err = json.NewEncoder(w).Encode(tyre)
		if err != nil {
			http.Error(w, "Failed to encode tyre: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("server started at port 8080")

	server := http.Server{
		Addr:    viper.GetString("APPLICATION_DOMAIN"),
		Handler: router,
		BaseContext: func(listener net.Listener) context.Context {
			return context.WithValue(ctx, "address", listener.Addr().String())
		},
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
