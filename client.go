package amplitude

import (
	"encoding/json"
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

// ResponseError is a description of the error returned from Amplitude API.
type ResponseError struct {
	StatusCode int
	Body       []byte
}

func (e ResponseError) Error() string {
	return fmt.Sprintf("Expected status code 200, got %d. Body: %s",
		e.StatusCode, e.Body)
}

// Client wraps the API endpoint and allows events to be sent.
type Client struct {
	// APIKey provided by Amplitude for the account.
	APIKey string

	// URL endpoint for Amplitude HTTP-API.
	URL string
}

// NewClient returns a client with default values for the Amplitude HTTP-API
// implementation.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		URL:    APIURL,
	}
}

// SendEvent sends one or more events to Amplitude using the client config. If
// more than one event is passed, the payload is sent as an array of events.
func (c Client) SendEvent(e ...Event) ([]byte, error) {
	url := strings.Join([]string{c.URL, EventsPath}, "/")

	vals, err := encode(c.APIKey, e)
	if err != nil {
		return nil, fmt.Errorf("Encode payload failed (%v)", err)
	}

	client := &http.Client{}
	res, err := client.PostForm(url, vals)
	if err != nil {
		return nil, fmt.Errorf("Server request failed (%v)", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, ResponseError{StatusCode: res.StatusCode, Body: body}
	}
	return body, nil
}

func encode(key string, e []Event) (url.Values, error) {
	var err error
	var payload []byte
	vals := url.Values{}
	vals.Set("api_key", key)

	if len(e) == 1 {
		payload, err = json.Marshal(e[0])
		vals.Set("event", string(payload))
	} else {
		payload, err = json.Marshal(e)
		vals.Set("events", string(payload))
	}
	return vals, err
}
