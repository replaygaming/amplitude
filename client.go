package amplitude

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// APIURL is the url for the HTTP-API endpoint.
	APIURL = "https://api.amplitude.com/"

	// EventsPath is the url part to post/get requests for events.
	EventsPath = "httpapi"
)

// Client interface requires a Send method to post a payload to Amplitude.
type Client interface {
	Send(Payload) ([]byte, error)
}

// ResponseError is a description of the error returned from Amplitude API.
type ResponseError struct {
	StatusCode int
	Body       []byte
}

func (e ResponseError) Error() string {
	return fmt.Sprintf("Expected status code 200, got %d. Body: %s",
		e.StatusCode, e.Body)
}

// DefaultClient implements the client interface and wraps the API endpoint and
// allows events to be sent.
type DefaultClient struct {
	// APIKey provided by Amplitude for the account.
	APIKey string

	// URL endpoint for Amplitude HTTP-API.
	URL string
}

// NewClient returns a default client with values for the Amplitude HTTP-API
// implementation.
func NewClient(apiKey string) *DefaultClient {
	return &DefaultClient{
		APIKey: apiKey,
		URL:    APIURL,
	}
}

// Send sends a payload (an event or a list of events) to Amplitude using the
// client config.
func (c *DefaultClient) Send(p Payload) ([]byte, error) {
	path := strings.Join([]string{c.URL, EventsPath}, "/")

	data, err := p.Value()
	if err != nil {
		return nil, fmt.Errorf("Payload encoding failed (%v)", err)
	}

	vals := url.Values{}
	vals.Set("api_key", c.APIKey)
	vals.Set(p.Key(), string(data))

	client := &http.Client{}
	res, err := client.PostForm(path, vals)
	if err != nil {
		return nil, fmt.Errorf("Server request failed (%v)", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil || res.StatusCode != 200 {
		return nil, ResponseError{StatusCode: res.StatusCode, Body: body}
	}
	return body, nil
}

// NoopClient implements the client interface, but doesn't do anything.
type NoopClient struct {
}

// Send return a blank body and no error.
func (c NoopClient) Send(Payload) ([]byte, error) {
	return []byte(""), nil
}
