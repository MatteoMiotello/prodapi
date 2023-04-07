package fs_handlers

import "strings"

type ImagesHandler struct {
	baseUrl  string
	basePath string
}

func NewImagesHandler(baseUrl string) *ImagesHandler {
	return &ImagesHandler{
		baseUrl:  baseUrl,
		basePath: "images/tyres",
	}
}

func clean(path string) string {
	path = strings.TrimSuffix(path, "/")
	return strings.TrimPrefix(path, "/")
}

func concat(paths ...string) string {
	var cleaned []string
	for _, path := range paths {
		cleaned = append(cleaned, clean(path))
	}

	return strings.Join(cleaned, "/")
}

func (i ImagesHandler) GetRelativePath() string {
	return concat("./", i.basePath)
}

func (i ImagesHandler) GetFileRelativePath(fileName string) string {
	return concat(i.GetRelativePath(), fileName)
}

func (i ImagesHandler) GetBaseUrl() string {
	return concat(i.baseUrl, i.basePath)
}

func (i ImagesHandler) GetPublicUrl(filename string) string {
	return concat(i.GetBaseUrl(), filename)
}
