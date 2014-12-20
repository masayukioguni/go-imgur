package imgur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// const
const (
	LibraryVersion = "0.1"
	defaultBaseURL = "https://api.imgur.com/3/"
	userAgent      = "go-imgr/" + LibraryVersion
)

// Option Optional parameters
type Option struct {
	ClientID string
}

// A Client manages communication with the Digital Ocean API.
type Client struct {
	Option    *Option
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string

	GalleryService *GalleryService
}

// An Response by API request.
type ResponseStatus struct {
	Success bool
	Status  int
}

// ErrorData https://api.imgur.com/errorhandling
type ErrorData struct {
	Error   string `json:"error,omitempty"`
	Request string `json:"request,omitempty"`
	Method  string `json:"method,omitempty"`
}

// An ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response  *http.Response
	Success   bool `json:"success,omitempty"`
	Status    int  `json:"status,omitempty"`
	ErrorData ErrorData
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %+v ",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.ErrorData.Error)
}

// NewClient returns a new Imgur API client.
func NewClient(option *Option) (*Client, error) {
	httpClient := http.DefaultClient
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, Option: option}
	c.GalleryService = &GalleryService{client: c}

	return c, nil
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	authHeader := "Client-ID " + c.Option.ClientID

	req.Header.Set("Authorization", authHeader)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return resp, err

}

// CheckResponse checks the API response for errors, and returns them if present.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, &errorResponse.ErrorData)
	}

	return errorResponse
}
