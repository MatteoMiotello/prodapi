package main

import (
	"context"
	"fmt"
	"github.com/MatteoMiotello/prodapi/internal/bootstrap"
	"github.com/MatteoMiotello/prodapi/internal/controllers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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

	router.HandleFunc("/products/tyres/{tyre_code}", controllers.NewTyreController().FindByCode).Methods(http.MethodGet)

	router.HandleFunc("/resources/tyres/{tyre_code}", controllers.NewTyreController().FindImage).Methods(http.MethodGet)
	router.HandleFunc("/resources/brands/{brand_code}", controllers.NewBrandController().FindImage).Methods(http.MethodGet)

	fmt.Println("server started at " + viper.GetString(viper.GetString("APPLICATION_DOMAIN")))

	_ = router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, _ := route.GetMethods()
		template, _ := route.GetPathTemplate()
		fmt.Println(methods, template)

		return nil
	})

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
