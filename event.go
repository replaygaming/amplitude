package amplitude

import "encoding/json"

// Payload is the event data sent to the server.
type Payload interface {
	Encode() ([]byte, error)
}

// Properties is a dictionary type for both user and event properties.
type Properties map[string]*json.RawMessage

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

// Encode returns the event serialized into JSON.
func (e Event) Encode() ([]byte, error) {
	return json.Marshal(e)
}

// Events represents a list of events. Implements the Payload interface.
type Events []Event

// Encode returns all events serialized into JSON.
func (e Events) Encode() ([]byte, error) {
	return json.Marshal(e)
}

// Device fields must all be updated together. Setting any of these fields
// will automatically reset all of the others if they are not also set on the
// same event.
type Device struct {
	OSName             string `json:"os_name,omitempty"`
	OSVersion          string `json:"os_version,omitempty"`
	DeviceBrand        string `json:"device_brand,omitempty"`
	DeviceManufacturer string `json:"device_manufacturer,omitempty"`
	DeviceModel        string `json:"device_model,omitempty"`
	DeviceType         string `json:"device_type,omitempty"`
	Carrier            string `json:"carrier,omitempty"`
}

// Location fields must all be updated together. Setting any of these
// fields will automatically reset all of the others if they are not also set
// on the same event.
type Location struct {
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
	City    string `json:"city,omitempty"`
	DMA     string `json:"dma,omitempty"`
}

// EventAugment are optional keys interpreted in a special way by Amplitude.
type EventAugment struct {
	EventID   int    `json:"event_id,omitempty"`
	SessionID int64  `json:"session_id,omitempty"`
	InsertID  string `json:"insert_id,omitempty"`
}
