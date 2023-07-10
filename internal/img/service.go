package img

import (
	"crypto/tls"
	"errors"
	"github.com/MatteoMiotello/prodapi/internal/fs_handlers"
	"io"
	"mime"
	"net/http"
	"os"
)

type ImageService struct {
	FsHandler *fs_handlers.FsHandler
}

func NewImageService(fsHandler *fs_handlers.FsHandler) *ImageService {
	return &ImageService{
		FsHandler: fsHandler,
	}
}

func (s ImageService) SaveImageFromUrl(imageUrl string, fileName string) (*Image, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Get(imageUrl)
	if err != nil {
		return nil, errors.New("image_not_found")
	}

	_, err = os.ReadDir(s.FsHandler.GetRelativePath())

	if err != nil {
		err := os.Mkdir(s.FsHandler.GetRelativePath(), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	contentType := response.Header.Get("Content-Type")
	extensionsByType, err := mime.ExtensionsByType(contentType)

	extension := new(string)

	if extensionsByType == nil {
		extension = nil
	}

	if len(extensionsByType) > 0 {
		extension = &extensionsByType[len(extensionsByType)-1]
	}

	if extension == nil || (*extension != ".png" && *extension != ".jpg" && *extension != ".jpeg") {
		return nil, errors.New("incompatible_image")
	}

	filename := fileName + *extension
	filePath := s.FsHandler.GetFileRelativePath(filename)

	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	image := new(Image)

	image.Path = s.FsHandler.GetFileBaseRelativePath(filename)
	image.Extension = *extension
	image.OriginalUrl = &imageUrl

	return image, nil
}
