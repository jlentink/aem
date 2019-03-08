package main

type bundlesFeed struct {
	Status string   `json:"status"`
	S      []int    `json:"s"`
	Data   []bundle `json:"data"`
}

type bundle struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Fragment     bool   `json:"fragment"`
	StateRaw     int    `json:"stateRaw"`
	State        string `json:"state"`
	Version      string `json:"version"`
	SymbolicName string `json:"symbolicName"`
	Category     string `json:"category"`
}

type bundleResponse struct {
	Fragment bool `json:"fragment"`
	StateRaw int  `json:"stateRaw"`
}

var bundleRawState = map[int]string{
	1:  "Uninstalled",
	2:  "Installed",
	4:  "Resolved",
	8:  "Starting",
	16: "Stopping",
	32: "Active",
}
