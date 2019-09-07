package bundle

const (
	bundlesURL    = "/system/console/bundles"
	bundlePageURL = "/system/console/bundles/%s"
)

// BundleRawState Bundles states from integer to string map
var BundleRawState = map[int]string{
	0:  "Unknown",
	1:  "Uninstalled",
	2:  "Installed",
	4:  "Resolved",
	8:  "Starting",
	16: "Stopping",
	32: "Active",
}

const (
	bundleFormActionField = "action"
	bundleInstall         = "install"
	bundleRefresh         = "refreshPackages"
)

type bundlesFeed struct {
	Status string   `json:"status"`
	S      []int    `json:"s"`
	Data   []Bundle `json:"data"`
}

//Bundle OSGI bundle struct
type Bundle struct {
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
