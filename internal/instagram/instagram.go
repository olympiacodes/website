package instagram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct{}

func (c *Client) MediaForUser(username string) ([]Media, error) {
	res, err := http.Get("https://www.instagram.com/" + username + "/?__a=1")
	if err != nil {
		return []Media{}, err
	}

	if res.StatusCode != http.StatusOK {
		return []Media{}, fmt.Errorf("API returned non-200 status code: %d", res.StatusCode)
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Media{}, err
	}

	var r UserResponse
	err = json.Unmarshal(data, &r)
	if err != nil {
		return []Media{}, err
	}

	fmt.Printf("Media: %+v\n", r.User.Media.Nodes)

	return r.User.Media.Nodes, nil
}

type UserResponse struct {
	User User `json:"user"`
}
type User struct {
	Username  string `json:"username"`
	Biography string `json:"biography"`
	Media     struct {
		Nodes []Media `json:"nodes"`
	} `json:"media"`
}

type Media struct {
	ID         string    `json:"id"`
	Type       MediaType `json:"__typename"`
	Dimensions struct {
		Width  int
		Height int
	} `json:"dimensions"`
	Thumbnail string `json:"thumbnail_src"`
	Caption   string `json:"caption"`
}

type MediaType string

const ImageMediaType = MediaType("GraphImage")
