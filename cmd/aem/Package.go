package main

import "encoding/xml"

type PackagesFeed struct {
	Package []PackageDescription `json:"results"`
	Total   int                  `json:"total"`
}

type PackageDescription struct {
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
	Filter          []struct {
		Root  string `json:"root"`
		Rules []struct {
			Modifier string `json:"modifier"`
			Pattern  string `json:"pattern"`
		} `json:"rules"`
	} `json:"filter"`
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

type CrxResponse struct {
	XMLName   xml.Name `xml:"crx"`
	Text      string   `xml:",chardata"`
	Version   string   `xml:"version,attr"`
	User      string   `xml:"user,attr"`
	Workspace string   `xml:"workspace,attr"`
	Request   struct {
		Text  string `xml:",chardata"`
		Param []struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"param"`
	} `xml:"request"`
	Response struct {
		Text string `xml:",chardata"`
		Data struct {
			Text    string `xml:",chardata"`
			Package struct {
				Text  string `xml:",chardata"`
				Group struct {
					Text string `xml:",chardata"`
				} `xml:"group"`
				Name struct {
					Text string `xml:",chardata"`
				} `xml:"name"`
				Version struct {
					Text string `xml:",chardata"`
				} `xml:"version"`
				DownloadName struct {
					Text string `xml:",chardata"`
				} `xml:"downloadName"`
				Size struct {
					Text string `xml:",chardata"`
				} `xml:"size"`
				Created struct {
					Text string `xml:",chardata"`
				} `xml:"created"`
				CreatedBy struct {
					Text string `xml:",chardata"`
				} `xml:"createdBy"`
				LastModified struct {
					Text string `xml:",chardata"`
				} `xml:"lastModified"`
				LastModifiedBy struct {
					Text string `xml:",chardata"`
				} `xml:"lastModifiedBy"`
				LastUnpacked struct {
					Text string `xml:",chardata"`
				} `xml:"lastUnpacked"`
				LastUnpackedBy struct {
					Text string `xml:",chardata"`
				} `xml:"lastUnpackedBy"`
			} `xml:"package"`
			Log struct {
				Text string `xml:",chardata"`
			} `xml:"log"`
		} `xml:"data"`
		Status struct {
			Text string `xml:",chardata"`
			Code string `xml:"code,attr"`
		} `xml:"status"`
	} `xml:"response"`
}
