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
		t.Errorf("Expected response error to be: %s\n got: %s", expected, result)
	}
}

func TestDefaultClient_NewClient(t *testing.T) {
	s := NewClient(testKey)
	if s.APIKey != testKey {
		t.Errorf("Expected APIKey to be: %s\n got: %s", testKey, s.APIKey)
	}
}

func TestDefaultClient_SendEvent(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	}))
	defer ts.Close()

	s := NewClient(testKey)
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

func TestDefaultClient_SendEvent_UnavailableServer(t *testing.T) {
	ts := httptest.NewUnstartedServer(nil)

	s := NewClient(testKey)
	s.URL = ts.URL

	e := Event{EventType: "test", UserID: "1", Revenue: 9.99}
	_, err := s.SendEvent(e)

	if err == nil {
		t.Error("Expected request to fail")
	}
}

func TestDefaultClient_SendEvent_ResponseError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid Key")
	}))
	defer ts.Close()

	s := NewClient(testKey)
	s.URL = ts.URL

	e := Event{EventType: "test", UserID: "1", Revenue: 9.99}
	expected := "Expected status code 200, got 400. Body: Invalid Key"
	_, err := s.SendEvent(e)

	if err == nil || err.Error() != expected {
		t.Errorf("Expected error to be: %s\ngot: %s", expected, err)
	}
}

func TestNoopClient_SendEvent(t *testing.T) {
	c := &NoopClient{}
	e := Event{EventType: "test", UserID: "1"}
	res, err := c.SendEvent(e)
	if err != nil {
		t.Errorf("Expected send event to return no error\ngot: %s", err)
	}
	if !bytes.Equal(res, []byte("")) {
		t.Errorf("Expected response to be: empty\ngot: %s", res)
	}
}

func TestEncode(t *testing.T) {
	e1 := Event{EventType: "signup", UserID: "1"}
	e2 := Event{EventType: "purchase", UserID: "2"}

	single := `api_key=abc&event={"user_id":"1","event_type":"signup"}`
	multi := `api_key=abc&events=[{"user_id":"1","event_type":"signup"},{"user_id":"2","event_type":"purchase"}]`

	result, _ := encode(testKey, []Event{e1})
	enc, _ := url.QueryUnescape(result.Encode())

	if enc != single {
		t.Errorf("Expected encoding a single event to equal: %q\ngot:%q",
			single, enc)
	}

	result, _ = encode(testKey, []Event{e1, e2})
	enc, _ = url.QueryUnescape(result.Encode())

	if enc != multi {
		t.Errorf("Expected encoding multiple events to equal: %q\ngot:%q",
			multi, enc)
	}
}
