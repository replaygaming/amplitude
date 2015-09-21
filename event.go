package amplitude

import "encoding/json"

// Properties is a dictionary type for both user and event properties
type Properties map[string]*json.RawMessage

// Event is the payload sent to the server
type Event struct {
	UserID             string     `json:"user_id,omitempty"`
	DeviceID           string     `json:"device_id,omitempty"`
	EventType          string     `json:"event_type"`
	Time               int64      `json:"time,omitempty"`
	EventProperties    Properties `json:"event_properties,omitempty"`
	UserProperties     Properties `json:"user_properties,omitempty"`
	AppVersion         string     `json:"app_version,omitempty"`
	Platform           string     `json:"platform,omitempty"`
	OSName             string     `json:"os_name,omitempty"`
	OSVersion          string     `json:"os_version,omitempty"`
	DeviceBrand        string     `json:"device_brand,omitempty"`
	DeviceManufacturer string     `json:"device_manufacturer,omitempty"`
	DeviceModel        string     `json:"device_model,omitempty"`
	DeviceType         string     `json:"device_type,omitempty"`
	Carrier            string     `json:"carrier,omitempty"`
	Country            string     `json:"country,omitempty"`
	Region             string     `json:"region,omitempty"`
	City               string     `json:"city,omitempty"`
	DMA                string     `json:"dma,omitempty"`
	Language           string     `json:"language,omitempty"`
	Revenue            float64    `json:"revenue,omitempty"`
	LocationLat        float64    `json:"location_lat,omitempty"`
	LocationLng        float64    `json:"location_lng,omitempty"`
	IP                 string     `json:"ip,omitempty"`
	IDFA               string     `json:"idfa,omitempty"`
	ADID               string     `json:"adid,omitempty"`
}
