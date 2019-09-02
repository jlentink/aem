package objects

import "encoding/xml"

//CrxResponse descriptive struct
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
