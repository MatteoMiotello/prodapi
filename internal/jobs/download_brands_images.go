package jobs

import (
	"context"
	"errors"
	"fmt"
	"github.com/MatteoMiotello/prodapi/internal/clients"
	"github.com/MatteoMiotello/prodapi/internal/fs_handlers"
	"github.com/MatteoMiotello/prodapi/internal/img"
	"github.com/MatteoMiotello/prodapi/internal/nosql"
	"github.com/MatteoMiotello/prodapi/schemas"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func DownloadNextBrandImage() {
	ctx := context.Background()

	filter := bson.D{
		{"$or",
			bson.A{
				bson.D{{"image_path", bson.D{{"$exists", false}}}},
				bson.D{{"retry_image", bson.D{{"$eq", true}}}},
			},
		},
	}

	brand := new(schemas.Brand)
	err := nosql.BrandCollection().FindOne(ctx, filter).Decode(&brand)

	if strings.Contains(err.Error(), "no documents in result") {
		return
	}

	if err != nil {
		panic(err)
	}

	bClient := clients.NewBingClient()

	fmt.Println(brand.Name)
	res, err := bClient.SearchBrandImage(brand.Name)

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(res.Value) == 0 {
		fmt.Println("no image found")
		return
	}

	index := 0

	if brand.ImageIndex != 0 {
		index = brand.ImageIndex
	}

	if brand.RetryImage {
		index = index + 1
	}

	err = updateImage(ctx, res, brand, index)

	if err != nil {
		if err.Error() == "no_image_found" {
			brand.RetryImage = false
			brand.Incomplete = true

			update := bson.M{
				"$set": brand,
			}

			_, err = nosql.BrandCollection().UpdateByID(ctx, brand.ID, update)

			if err != nil {
				panic(err)
			}

			return
		}

		panic(err)
	}
}

func updateImage(ctx context.Context, res *clients.ImageResponse, brand *schemas.Brand, index int) error {
	if len(res.Value) < index+1 {
		return errors.New("no_image_found")
	}

	imageUrl := res.Value[index].ContentUrl
	iService := img.NewImageService(fs_handlers.NewBrandsHandler(viper.GetString("APPLICATION_URL")))

	image, err := iService.SaveImageFromUrl(imageUrl, brand.Code)

	if err != nil {
		if err.Error() == "incompatible_image" || err.Error() == "image_not_found" {
			index = index + 1
			return updateImage(ctx, res, brand, index)
		}

		return err
	}

	brand.ImagePath = image.Path
	brand.ImageExtension = image.Extension
	brand.ImageIndex = index
	brand.RetryImage = false
	brand.Incomplete = false

	fmt.Println(brand)
	update := bson.M{
		"$set": brand,
	}

	_, err = nosql.BrandCollection().UpdateByID(ctx, brand.ID, update)

	if err != nil {
		return err
	}

	return nil
}
