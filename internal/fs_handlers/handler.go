package fs_handlers

import "strings"

type FsHandler struct {
	baseUrl   string
	basePath  string
	publicUrl string
}

func NewImagesHandler(baseUrl string) *FsHandler {
	return &FsHandler{
		baseUrl:   baseUrl,
		basePath:  "img/tyres",
		publicUrl: "resources/tyres",
	}
}

func NewBrandsHandler(baseUrl string) *FsHandler {
	return &FsHandler{
		baseUrl:   baseUrl,
		basePath:  "img/brands",
		publicUrl: "resources/brands",
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

func (i FsHandler) GetRelativePath() string {
	return concat("./", i.basePath)
}

func (i FsHandler) GetFileRelativePath(fileName string) string {
	return concat(i.GetRelativePath(), fileName)
}

func (i FsHandler) GetFileBaseRelativePath(fileName string) string {
	return concat("/", i.basePath, fileName)
}

func (i FsHandler) GetBaseUrl() string {
	return concat(i.baseUrl, i.publicUrl)
}

func (i FsHandler) GetPublicUrl(filename string) string {
	return concat(i.GetBaseUrl(), filename)
}
