package exmo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
	Result   bool   `json:"result"`
	Message  string `json:"error"`
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
	u := c.BaseURL.ResolveReference(rel)

	var req *http.Request

	req, err = http.NewRequest(method, u.String(), nil)

	if err != nil {
		return nil, err
	}

	return req, nil
}

// newAuthenticatedRequest creates new http request for authenticated routes.
func (c *Client) newAuthenticatedRequest(refURL string, params url.Values) (*http.Request, error) {
	params.Add("nonce", nonce())

	content := params.Encode()

	req, _ := http.NewRequest("POST", BaseURL+refURL, bytes.NewBuffer([]byte(content)))

	sign := signPayload(content, c.APISecret)

	req.Header.Set("Key", c.APIKey)
	req.Header.Set("Sign", sign)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(content)))

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

// checkResponse checks response status code and response
// for errors.
func checkResponse(r *Response) error {
	errorResponse := &ErrorResponse{}
	err := json.Unmarshal(r.Body, errorResponse)

	if err != nil {
		errorResponse.Message = "Error decoding response error message. " +
			"Please see response body for more information."
	} else if !(errorResponse.Message == "") {
		return errorResponse
	}

	return nil
}

func signPayload(message string, secret string) string {
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(message))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func nonce() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Response.Request.Method,
		r.Response.Response.Request.URL,
		r.Response.Response.StatusCode,
		r.Message,
	)
}
