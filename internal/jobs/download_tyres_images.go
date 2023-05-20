package jobs

import (
	"context"
	"fmt"
	"github.com/MatteoMiotello/prodapi/internal/clients"
	"github.com/MatteoMiotello/prodapi/internal/fs_handlers"
	"github.com/MatteoMiotello/prodapi/internal/images"
	"github.com/MatteoMiotello/prodapi/internal/nosql"
	"github.com/MatteoMiotello/prodapi/schemas"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

func DownloadNextTyreImage() {
	ctx := context.Background()

	filter := bson.D{{"image_path", bson.D{{"$exists", false}}}}

	var tyre schemas.Tyre
	err := nosql.TyreCollection().FindOne(ctx, filter).Decode(&tyre)
	if err != nil {
		panic(err)
	}

	bClient := clients.NewBingClient()

	fmt.Println(tyre.Reference)

	res, err := bClient.SearchTyreImage(tyre.Reference)

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(res.Value) == 0 {
		fmt.Println("no image found")
		return
	}

	imageUrl := res.Value[0].ThumbnailUrl
	iService := images.NewImageService(fs_handlers.NewImagesHandler(viper.GetString("APPLICATION_URL")))

	image, err := iService.SaveImageFromUrl(imageUrl, tyre.Code)
	if err != nil {
		return
	}

	tyre.ImagePath = image.Path
	tyre.ImageExtension = image.Extension

	fmt.Println(tyre)
	update := bson.M{
		"$set": tyre,
	}

	_, err = nosql.TyreCollection().UpdateByID(ctx, tyre.ID, update)

	if err != nil {
		panic(err)
	}
}
