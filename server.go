package amplitude

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
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

// Server wraps the API endpoint and allows events to be sent.
type Server struct {
	// APIKey provided by Amplitude for the account.
	APIKey string

	// URL endpoint for Amplitude HTTP-API.
	URL string

	// Timeout is the max duration before timing out sending the event. NewServer
	// defaults to 10 seconds.
	Timeout time.Duration
}

// NewServer returns a server with default values for the Amplitude HTTP-API
// implementation.
func NewServer(apiKey string) *Server {
	return &Server{
		APIKey:  apiKey,
		URL:     APIURL,
		Timeout: 10 * time.Second,
	}
}

func (s Server) encodePayload(e []Event) (url.Values, error) {
	var err error
	var payload []byte

	v := url.Values{}
	v.Set("api_key", s.APIKey)

	if len(e) == 1 {
		payload, err = json.Marshal(e[0])
		v.Set("event", string(payload))
	} else {
		payload, err = json.Marshal(e)
		v.Set("events", string(payload))
	}
	return v, err
}

// SendEvent sends one or more events to Amplitude using the server config. If
// more than one event is passed, the payload is sent as an array of events.
func (s Server) SendEvent(e ...Event) ([]byte, error) {
	endpoint := strings.Join([]string{s.URL, EventsPath}, "/")

	client := &http.Client{Timeout: s.Timeout}
	payload, err := s.encodePayload(e)
	if err != nil {
		return nil, fmt.Errorf("Encode payload failed (%v)", err)
	}

	res, err := client.PostForm(endpoint, payload)
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
