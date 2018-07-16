package wpt

// LocationsResponse is the response wraper for a location list request.
type LocationsResponse struct {
	StatusCode int                 `json:"statusCode"`
	StatusText string              `json:"statusText"`
	Data       map[string]Location `json:"data"`
}

// Location is a single location.
type Location struct {
	Label         string `json:"Label"`
	Location      string `json:"location"`
	Browser       string `json:"Browser"`
	RelayServer   string `json:"relayServer"`
	RelayLocation string `json:"relayLocation"`
	LabelShort    string `json:"labelShort"`
	Default       bool   `json:"default"`
	PendingTests  struct {
		Total        int `json:"Total"`
		HighPriority int `json:"HighPriority"`
		LowPriority  int `json:"LowPriority"`
		Testing      int `json:"Testing"`
		Idle         int `json:"Idle"`
	} `json:"PendingTests"`
}

// Locations is a list of Location
type Locations []Location

// GetDefault returns the default location if your API Key runs a test.
func (locations Locations) GetDefault() (defaultLocation Location) {
	for _, location := range locations {
		if location.Default {
			defaultLocation = location
		}
	}
	return defaultLocation
}

// Valid checks to see if name exists in this list of locations.
func (locations Locations) Valid(name string) bool {
	for _, location := range locations {
		if location.Location == name {
			return true
		}
	}

	return false
}
