package amplitude

import "encoding/json"

// Payload is the event data sent to the server.
type Payload interface {
	Encode() ([]byte, error)
	Path() string
	Type() string
}

// Properties is a dictionary type for both user and event properties.
type Properties map[string]*json.RawMessage

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
