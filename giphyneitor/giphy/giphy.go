package giphy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiURL = "https://api.giphy.com/v1/"

type Client struct {
	apiKey string
}

type Gif struct {
	Name string
	URL  string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c Client) RandomGif() (*Gif, error) {
	endpoint := apiURL + "gifs/random"
	r, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request for: %s endpoint | %v", endpoint, err)
	}

	q := r.URL.Query()
	q.Add("api_key", c.apiKey)
	q.Add("rating", "G")
	r.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("error calling to %s | %v", err)
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, fmt.Errorf("error reading the response | %v", err)
	}

	var respJson struct {
		Data struct {
			Title  string `json:"title""`
			Images struct {
				Image struct {
					URL string `json:"url"`
				} `json:"original"`
			} `json:"images"`
		} `json:data`
	}
	if err := json.Unmarshal(body, &respJson); err != nil {
		return nil, fmt.Errorf("error unmarshalling response | %v", err)
	}

	return &Gif{Name: respJson.Data.Title, URL: respJson.Data.Images.Image.URL}, nil
}
