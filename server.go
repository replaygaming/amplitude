package amplitude

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// PkgVersion is the current version of this package. Follows SemVer
	// conventions.
	PkgVersion = "0.0.1"

	// APIURL is the url for the HTTP-API endpoint
	APIURL = "https://api.amplitude.com/"

	// EventsPath is the url part to post/get requests for events
	EventsPath = "httpapi"
)

// Server wraps the API endpoint and allows events to be sent
type Server struct {
	// APIKey provided by Amplitude for the account
	APIKey string

	// URL endpoint for Amplitude HTTP-API
	URL string
}

// NewServer returns a server with default values for the Amplitude HTTP-API
// implementation.
func NewServer(apiKey string) *Server {
	return &Server{
		APIKey: apiKey,
		URL:    APIURL,
	}
}

// SendEvent posts a single event to Amplitude using the server config
func (s *Server) SendEvent(e Event) error {
	payload, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("Marshal payload failed (%v)", err)
	}
	url := strings.Join([]string{s.URL, EventsPath}, "/")
	result, err := s.post(url, payload)
	if err != nil {
		return err
	}

	log.Printf("[INFO AMPLITUDE] Event sent (%s), response: %s\n", payload, result)

	return nil
}

// Post sends a payload using the server config
func (s *Server) post(endpoint string, payload []byte) ([]byte, error) {
	v := url.Values{}
	v.Set("api_key", s.APIKey)
	v.Set("event", string(payload))

	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.PostForm(endpoint, v)
	if err != nil {
		return nil, fmt.Errorf("Server request failed (%v)", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Expected status code 200, got %d. Body: %s",
			res.StatusCode, body)
	}
	return body, nil
}
