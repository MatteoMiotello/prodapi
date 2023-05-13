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
	"io"
	"mime"
	"net/http"
	"os"
)

func DownloadNextImage() {
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

	response, err := http.Get(imageUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println(imageUrl)

	fsHandler := fs_handlers.NewImagesHandler(viper.GetString("APPLICATION_URL"))
	_, err = os.ReadDir(fsHandler.GetRelativePath())

	if err != nil {
		err := os.Mkdir(fsHandler.GetRelativePath(), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	contentType := response.Header.Get("Content-Type")
	extensionsByType, err := mime.ExtensionsByType(contentType)

	filename := tyre.Code + extensionsByType[1]
	filePath := fsHandler.GetFileRelativePath(filename)

	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return
	}

	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		return
	}

	tyre.ImagePath = fsHandler.GetFileBaseRelativePath(filename)
	tyre.ImageExtension = extensionsByType[1]

	fmt.Println(tyre)
	update := bson.M{
		"$set": tyre,
	}

	_, err = nosql.TyreCollection().UpdateByID(ctx, tyre.ID, update)

	if err != nil {
		panic(err)
	}
}
