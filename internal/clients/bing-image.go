package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"time"
)

type BingImageClient struct {
	endpoint string
}

func NewBingClient() BingImageClient {
	return BingImageClient{
		endpoint: viper.GetString("BING_ENDPOINT"),
	}
}

type ImageResponse struct {
	Type            string `json:"_type"`
	Instrumentation struct {
		Type string `json:"_type"`
	} `json:"instrumentation"`
	ReadLink     string `json:"readLink"`
	WebSearchUrl string `json:"webSearchUrl"`
	QueryContext struct {
		OriginalQuery           string `json:"originalQuery"`
		AlterationDisplayQuery  string `json:"alterationDisplayQuery"`
		AlterationOverrideQuery string `json:"alterationOverrideQuery"`
		AlterationMethod        string `json:"alterationMethod"`
		AlterationType          string `json:"alterationType"`
	} `json:"queryContext"`
	TotalEstimatedMatches int64 `json:"totalEstimatedMatches"`
	NextOffset            int64 `json:"nextOffset"`
	CurrentOffset         int64 `json:"currentOffset"`
	Value                 []struct {
		WebSearchUrl           string    `json:"webSearchUrl"`
		Name                   string    `json:"name"`
		ThumbnailUrl           string    `json:"thumbnailUrl"`
		DatePublished          time.Time `json:"datePublished"`
		IsFamilyFriendly       bool      `json:"isFamilyFriendly"`
		ContentUrl             string    `json:"contentUrl"`
		HostPageUrl            string    `json:"hostPageUrl"`
		ContentSize            string    `json:"contentSize"`
		EncodingFormat         string    `json:"encodingFormat"`
		HostPageDisplayUrl     string    `json:"hostPageDisplayUrl"`
		Width                  int64     `json:"width"`
		Height                 int64     `json:"height"`
		HostPageDiscoveredDate time.Time `json:"hostPageDiscoveredDate"`
		Thumbnail              struct {
			Width  int64 `json:"width"`
			Height int64 `json:"height"`
		} `json:"thumbnail"`
		ImageInsightsToken string      `json:"imageInsightsToken,omitempty"`
		InsightsMetadata   interface{} `json:"insightsMetadata,omitempty"`
		ImageId            string      `json:"imageId,omitempty"`
	} `json:"value"`
}

func makeRequest[T interface{}](endpoint string, queryParams map[string]string, response *T) error {
	client := &http.Client{}

	params := url.Values{}

	for key, value := range queryParams {
		params.Set(key, value)
	}

	req, err := http.NewRequest(http.MethodGet, endpoint+"?"+params.Encode(), nil)

	if err != nil {
		return err
	}

	req.Header.Set("Ocp-Apim-Subscription-Key", viper.GetString("BING_KEY"))
	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(response)

	if err != nil {
		return err
	}

	return nil
}

func (b BingImageClient) SearchTyreImage(searchString string) (*ImageResponse, error) {
	params := map[string]string{
		"count":      "1",
		"q":          searchString,
		"safeSearch": "Strict",
		"aspect":     "Tall",
		"size":       "Large",
		"minWidth":   "400",
	}

	tyreRes := new(ImageResponse)

	err := makeRequest[ImageResponse](b.endpoint+"v7.0/img/search", params, tyreRes)
	if err != nil {
		return nil, err
	}

	return tyreRes, nil
}

func (b BingImageClient) SearchBrandImage(brandName string) (*ImageResponse, error) {
	searchString := brandName + " tyre logo"

	fmt.Println(searchString)

	params := map[string]string{
		"q":        searchString,
		"aspect":   "Wide",
		"size":     "Large",
		"minWidth": "400",
	}

	brandRes := new(ImageResponse)

	err := makeRequest[ImageResponse](b.endpoint+"v7.0/img/search", params, brandRes)
	fmt.Println(brandRes)
	if err != nil {
		return nil, err
	}

	return brandRes, nil
}
