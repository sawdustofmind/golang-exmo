package exmo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// BaseURL is root API url
	BaseURL = "https://api.exmo.com/v1/"
)

// Response is wrapper for standard http.Response and provides
// more methods.
type Response struct {
	Response *http.Response
	Body     []byte
}

// ErrorResponse is the custom error type that is returned if the API returns an
// error.
type ErrorResponse struct {
	Response *Response
	Message  string `json:"message"`
}

// NewRequest create new API request. Relative url can be provided in refURL.
func (c *Client) newRequest(method string, refURL string, params url.Values) (*http.Request, error) {
	rel, err := url.Parse(refURL)
	if err != nil {
		return nil, err
	}
	if params != nil {
		rel.RawQuery = params.Encode()
	}
	var req *http.Request
	u := c.BaseURL.ResolveReference(rel)
	req, err = http.NewRequest(method, u.String(), nil)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) performRequest(req *http.Request, v interface{}) (*Response, error) {
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		body = []byte(`Error reading body:` + err.Error())
	}

	response := &Response{resp, body}

	err = checkResponse(response)
	if err != nil {
		// Return response in case caller need to debug it.
		return response, err
	}

	if v != nil {
		err = json.Unmarshal(response.Body, v)
		if err != nil {
			return response, err
		}
	}

	return response, nil
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Response.Request.Method,
		r.Response.Response.Request.URL,
		r.Response.Response.StatusCode,
		r.Message,
	)
}
