package objects

// Description of pkg in aem
type Description struct {
	Pid             string        `json:"pid"`
	Path            string        `json:"path"`
	Name            string        `json:"name"`
	DownloadName    string        `json:"downloadName"`
	Group           string        `json:"group"`
	GroupTitle      string        `json:"groupTitle"`
	Version         string        `json:"version"`
	Description     string        `json:"description,omitempty"`
	Thumbnail       string        `json:"thumbnail"`
	BuildCount      uint          `json:"buildCount"`
	Created         int64         `json:"created,omitempty"`
	CreatedBy       string        `json:"createdBy,omitempty"`
	LastUnpacked    int64         `json:"lastUnpacked,omitempty"`
	LastUnpackedBy  string        `json:"lastUnpackedBy,omitempty"`
	LastUnwrapped   int64         `json:"lastUnwrapped,omitempty"`
	Size            uint64        `json:"size"`
	HasSnapshot     bool          `json:"hasSnapshot"`
	NeedsRewrap     bool          `json:"needsRewrap"`
	RequiresRoot    bool          `json:"requiresRoot"`
	RequiresRestart bool          `json:"requiresRestart"`
	AcHandling      string        `json:"acHandling"`
	Dependencies    []interface{} `json:"dependencies"`
	Resolved        bool          `json:"resolved"`
	Filter          []Filter      `json:"filter"`
	Screenshots     []interface{} `json:"screenshots"`
	LastModified    int64         `json:"lastModified,omitempty"`
	LastModifiedBy  string        `json:"lastModifiedBy,omitempty"`
	LastWrapped     int64         `json:"lastWrapped,omitempty"`
	LastWrappedBy   string        `json:"lastWrappedBy,omitempty"`
	LastUnwrappedBy string        `json:"lastUnwrappedBy,omitempty"`
	BuiltWith       string        `json:"builtWith,omitempty"`
	TestedWith      string        `json:"testedWith,omitempty"`
	FixedBugs       string        `json:"fixedBugs,omitempty"`
	ProviderName    string        `json:"providerName,omitempty"`
	ProviderURL     string        `json:"providerUrl,omitempty"`
	ProviderLink    string        `json:"providerLink,omitempty"`
}

// Filter filter of pkg
type Filter struct {
	Root  string  `json:"root"`
	Rules []Rules `json:"rules"`
}

// Rules of filter in pkg
type Rules struct {
	Modifier string `json:"modifier"`
	Pattern  string `json:"pattern"`
}
