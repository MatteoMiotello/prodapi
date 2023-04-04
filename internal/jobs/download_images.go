package jobs

import (
	"context"
	"fmt"
	"github.com/MatteoMiotello/prodapi/internal/clients"
	"github.com/MatteoMiotello/prodapi/internal/fs_handlers"
	"github.com/MatteoMiotello/prodapi/internal/nosql"
	"github.com/MatteoMiotello/prodapi/schemas"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"mime"
	"net/http"
	"os"
)

func DownloadNextImage() {
	ctx := context.Background()

	filter := bson.D{{"image_url", bson.D{{"$exists", false}}}}

	var tyre schemas.Tyre
	nosql.TyreCollection().FindOne(ctx, filter).Decode(&tyre)

	bClient := clients.NewBingClient()

	res, err := bClient.SearchTyreImage(tyre.Reference)

	if err != nil {
		panic(err)
	}

	imageUrl := res.Value[0].ThumbnailUrl

	tyre.ImageUrl = imageUrl

	update := bson.M{
		"$set": tyre,
	}

	_, err = nosql.TyreCollection().UpdateByID(ctx, tyre.ID, update)

	if err != nil {
		panic(err)
	}

	fmt.Println(tyre.ID)

	return

	response, err := http.Get(imageUrl)
	if err != nil {
		panic(err)
	}

	_, err = os.ReadDir("./images")

	if err != nil {
		err := os.Mkdir("./images", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	contentType := response.Header.Get("Content-Type")

	byType, err := mime.ExtensionsByType(contentType)

	filename := tyre.Code + byType[1]
	filePath := "./images/" + filename

	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)

	err = response.Write(file)

	if err != nil {
		panic(err)
	}

	file.Close()

	fsHandler := fs_handlers.NewImagesHandler(viper.GetString("APPLICATION_URL"))
	url := fsHandler.GetPublicUrl(filename)

	tyre.ImageUrl = url

	fmt.Println(tyre)

	_, err = nosql.TyreCollection().UpdateByID(ctx, tyre.ID, tyre)
	if err != nil {
		panic(err)
	}

	fmt.Println(tyre.Code)
}
