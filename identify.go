package amplitude

import "encoding/json"

// Identification implements the Payload interface.
type Identification struct {
	UserID         string     `json:"user_id,omitempty"`
	DeviceID       string     `json:"device_id,omitempty"`
	UserProperties Properties `json:"user_properties,omitempty"`
	AppVersion     string     `json:"app_version,omitempty"`
	Platform       string     `json:"platform,omitempty"`
	Paying         string     `json:"paying,omitempty"`
	StartVersion   string     `json:"start_version,omitempty"`
	Device
	Location
}

const (
	// IdentifyPath is the path for the Identify endpoint.
	IdentifyPath = "identify"
)

// Encode returns the identification serialized into JSON.
func (i Identification) Encode() ([]byte, error) {
	return json.Marshal(i)
}

// Path returns the url for the client to call
func (i Identification) Path() string {
	return IdentifyPath
}

// Type sets the type of identification key for the api
func (i Identification) Type() string {
	return "identification"
}

// Identifications represents a list of identifications. Implements the Payload interface.
type Identifications []Identification

// Encode returns all identifications serialized into JSON.
func (i Identifications) Encode() ([]byte, error) {
	return json.Marshal(i)
}

// Path returns the url for the client to call
func (i Identifications) Path() string {
	return IdentifyPath
}

// Type sets the type of identification key for the api
func (i Identifications) Type() string {
	return "identification"
}
