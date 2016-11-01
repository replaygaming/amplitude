package amplitude

import (
	"bytes"
	"testing"
)

func TestIdentification_Encode(t *testing.T) {
	e := Identification{Paying: "true", UserID: "1"}
	expected := []byte(`{"user_id":"1","paying":"true"}`)
	result, _ := e.Encode()

	if !bytes.Equal(expected, result) {
		t.Errorf("Expected encoding a single event to equal: %q\ngot:%q",
			expected, result)
	}
}

func TestIdentifications_Encode(t *testing.T) {
	e := Identifications{
		{Paying: "false", UserID: "1"},
		{Paying: "true", UserID: "2"},
	}
	expected := []byte(`[{"user_id":"1","paying":"false"},{"user_id":"2","paying":"true"}]`)
	result, _ := e.Encode()

	if !bytes.Equal(expected, result) {
		t.Errorf("Expected encoding multiple events to equal: %q\ngot:%q",
			expected, result)
	}
}
