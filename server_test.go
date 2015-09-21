package amplitude

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var testKey = "abc"

func TestResponseError_Error(t *testing.T) {
	err := ResponseError{StatusCode: 400, Body: []byte("Invalid key")}
	expected := "Expected status code 200, got 400. Body: Invalid key"
	result := err.Error()
	if expected != result {
		t.Errorf("Expecting response error to be: %s\n got: %s", expected, result)
	}
}

func TestServer_NewServer(t *testing.T) {
	s := NewServer(testKey)
	if s.APIKey != testKey {
		t.Errorf("Expecting APIKey to be: %s\n got: %s", testKey, s.APIKey)
	}
}

func TestServer_EncodeEvent(t *testing.T) {
	e1 := Event{EventType: "signup", UserID: "1"}
	e2 := Event{EventType: "purchase", UserID: "2"}
	s := NewServer(testKey)

	single := `api_key=abc&event={"user_id":"1","event_type":"signup"}`
	multi := `api_key=abc&events=[{"user_id":"1","event_type":"signup"},{"user_id":"2","event_type":"purchase"}]`

	result, _ := s.encodePayload([]Event{e1})
	enc, _ := url.QueryUnescape(result.Encode())

	if enc != single {
		t.Errorf("Expecting encoding a single event to equal: %q\ngot:%q",
			single, enc)
	}

	result, _ = s.encodePayload([]Event{e1, e2})
	enc, _ = url.QueryUnescape(result.Encode())

	if enc != multi {
		t.Errorf("Expecting encoding multiple events to equal: %q\ngot:%q",
			multi, enc)
	}
}

func TestServer_SendEvent(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	}))
	defer ts.Close()

	s := NewServer(testKey)
	s.URL = ts.URL

	e := Event{EventType: "test", UserID: "1", Revenue: 9.99}
	res, err := s.SendEvent(e)

	if err != nil {
		t.Errorf("Expected response to succeed, got error: %s", err)
	}

	expected := []byte("OK")
	if !bytes.Equal(res, expected) {
		t.Errorf("Expected response to be: %s, was: %s", expected, res)
	}
}

func TestServer_SendEvent_ResponseError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid Key")
	}))
	defer ts.Close()

	s := NewServer(testKey)
	s.URL = ts.URL

	e := Event{EventType: "test", UserID: "1", Revenue: 9.99}
	expected := "Expected status code 200, got 400. Body: Invalid Key"
	_, err := s.SendEvent(e)

	if err == nil || err.Error() != expected {
		t.Errorf("Expecting error to be: %s\ngot: %s", expected, err)
	}
}
