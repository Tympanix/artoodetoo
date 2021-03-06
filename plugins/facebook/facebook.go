package facebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	// VERSION is the version of the facebok api
	VERSION = "v2.9"
	// URI is the base url for the facebook api
	URI = "https://graph.facebook.com"
	// TIME is a time interpolation example used to parse time
	TIME = "2006-01-02T15:04:05-0700"
)

type facebookStyle struct{}

func (facebookStyle) Color() uint {
	return 0x3b5998
}

func (facebookStyle) Icon() string {
	return "fa-facebook"
}

// Token is a authentication token used to communicate with the facebook api
type Token string

// Options defines the options with which the api is called
type Options struct {
	Fields string
}

// Data is a json structure returned by the facebook api
type Data struct {
	Data   *json.RawMessage `json:"data"`
	Paging Pagination       `json:"paging"`
}

// Pagination is a struct which handles pagination of api nodes
type Pagination struct {
	Cursor Cursor `json:"cursor"`
	Next   string `json:"next"`
}

// Cursor handles cursers from/to content
type Cursor struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

func getURL(path string, token Token, options Options) (url *url.URL, err error) {
	url, err = url.Parse(URI)

	q := url.Query()
	q.Set("access_token", string(token))

	if len(options.Fields) > 0 {
		q.Add("fields", options.Fields)
	}

	url.RawQuery = q.Encode()

	url.Path = fmt.Sprintf("/%s%s", VERSION, path)

	return
}

func callAPI(path string, token Token, options Options) (resp *http.Response, err error) {
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}

	url, err := getURL(path, token, options)
	if err != nil {
		return
	}

	if resp, err = http.Get(url.String()); err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New("Could not call facebook API")
		return
	}
	return
}

func getData(path string, token Token, options Options, v interface{}) (paging Pagination, err error) {
	r, err := callAPI(path, token, options)
	if err != nil {
		return
	}
	var data Data
	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		return
	}
	if data.Data == nil {
		err = errors.New("No data returned from facebook API")
		return
	}
	paging = data.Paging
	if err = json.Unmarshal(*data.Data, v); err != nil {
		return
	}
	return
}

func getNode(path string, token Token, options Options, v interface{}) (err error) {
	r, err := callAPI(path, token, options)

	if err != nil {
		return
	}

	if err = json.NewDecoder(r.Body).Decode(v); err != nil {
		return
	}
	return nil
}
