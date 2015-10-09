package amplitude

import (
	"bytes"
	"testing"
)

func TestEvent_Encode(t *testing.T) {
	e := Event{EventType: "signup", UserID: "1"}
	expected := []byte(`{"user_id":"1","event_type":"signup"}`)
	result, _ := e.Encode()

	if !bytes.Equal(expected, result) {
		t.Errorf("Expected encoding a single event to equal: %q\ngot:%q",
			expected, result)
	}
}

func TestEvents_Encode(t *testing.T) {
	e := Events{
		{EventType: "signup", UserID: "1"},
		{EventType: "purchase", UserID: "2"},
	}
	expected := []byte(`[{"user_id":"1","event_type":"signup"},{"user_id":"2","event_type":"purchase"}]`)
	result, _ := e.Encode()

	if !bytes.Equal(expected, result) {
		t.Errorf("Expected encoding multiple events to equal: %q\ngot:%q",
			expected, result)
	}
}
