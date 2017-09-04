package instagram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct{}

func (c *Client) MediaForUser(username string) (MediaResponse, error) {
	var media MediaResponse

	res, err := http.Get("https://www.instagram.com/" + username + "/media/")
	if err != nil {
		return media, err
	}

	if res.StatusCode != http.StatusOK {
		return media, fmt.Errorf("API returned non-200 status code: %d", res.StatusCode)
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return media, err
	}

	err = json.Unmarshal(data, &media)
	if err != nil {
		return media, err
	}

	return media, nil

}

type MediaResponse struct {
	Media         []Media `json:"items"`
	MoreAvailable bool    `json:"more_available"`
	Status        Status  `json:"status"`
}

type Status string

const (
	StatosOK = Status("ok")
)

type Media struct {
	ID      string              `json:"id"`
	Code    string              `json:"code"`
	URL     string              `json:"link"`
	Type    string              `json:"type"`
	Caption Caption             `json:"caption"`
	Images  map[ImageSize]Image `json:"images"`
}

type MediaType string

const ImageMediaType = MediaType("image")

type ImageSize string

const (
	ImageSizeThumbnail          = ImageSize("thumbnail")
	ImageSizeLowResultion       = ImageSize("low_resolution")
	ImageSizeStandardResolution = ImageSize("standard_resolution")
)

type Image struct {
	URL    string `json:"url"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
}

type Caption struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
