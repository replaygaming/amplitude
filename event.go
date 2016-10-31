package amplitude

import "encoding/json"

// Event implements the Payload interface.
type Event struct {
	UserID          string     `json:"user_id,omitempty"`
	DeviceID        string     `json:"device_id,omitempty"`
	EventType       string     `json:"event_type"`
	Time            int64      `json:"time,omitempty"`
	EventProperties Properties `json:"event_properties,omitempty"`
	UserProperties  Properties `json:"user_properties,omitempty"`
	AppVersion      string     `json:"app_version,omitempty"`
	Platform        string     `json:"platform,omitempty"`
	Language        string     `json:"language,omitempty"`
	Revenue         float64    `json:"revenue,omitempty"`
	LocationLat     float64    `json:"location_lat,omitempty"`
	LocationLng     float64    `json:"location_lng,omitempty"`
	IP              string     `json:"ip,omitempty"`
	IDFA            string     `json:"idfa,omitempty"`
	ADID            string     `json:"adid,omitempty"`
	Device
	Location
	EventAugment
}

const (
	// IdentifyPath is the path for the Identify endpoint.
	EventsPath = "httpapi"
)

// Encode returns the event serialized into JSON.
func (e Event) Encode() ([]byte, error) {
	return json.Marshal(e)
}

// Path returns the url for the client to call
func (e Event) Path() string {
	return EventsPath
}

// Type sets the type of event key for the api
func (e Event) Type() string {
	return "event"
}

// Events represents a list of events. Implements the Payload interface.
type Events []Event

// Encode returns all events serialized into JSON.
func (e Events) Encode() ([]byte, error) {
	return json.Marshal(e)
}

// Path returns the url for the client to call
func (e Events) Path() string {
	return EventsPath
}

// Type sets the type of event key for the api
func (e Events) Type() string {
	return "event"
}

// EventAugment are optional keys interpreted in a special way by Amplitude.
type EventAugment struct {
	EventID   int    `json:"event_id,omitempty"`
	SessionID int64  `json:"session_id,omitempty"`
	InsertID  string `json:"insert_id,omitempty"`
}
