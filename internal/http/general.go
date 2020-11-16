package http

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

const configNoCache = "no-cache"

// URLS aem
const (
	URLSystemInformation = "/libs/granite/operations/content/systemoverview/export.json"
	URLActivateTree      = "/etc/replication/treeactivation.html"
	URLBundles           = "/system/console/bundles"
	URLRebuildPackage    = "/crx/packmgr/service/.json%s?cmd=build"
	URLBundlePage        = "/system/console/bundles/%s"
	URLReplication       = "/bin/replicate.json"
	URLPackageList       = "/crx/packmgr/list.jsp"
	URLPackageEndpoint   = "/crx/packmgr/service.jsp"

	ServiceName = "aem-cli"

	JarContentType = "application/java-archive"
	ApplicationFormEncode = "application/x-www-form-urlencoded"
)

//DisableSSLValidation Disables SSL validation
func DisableSSLValidation() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

// URLToURLString  convert the url to string
func URLToURLString(u *url.URL) string {
	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.RequestURI())
}

func setupPlainToURL(uri string, username string, password string) (*url.URL, error){
	URL, err := url.Parse(uri)

	if username != "" || password != "" {
		URL.User = url.UserPassword(username, password)

	}
	if err != nil {
		return nil, err
	}

	if username != "" || password != "" {
		URL.User = url.UserPassword(username, password)

	}
	if err != nil {
		return nil, err
	}
	return URL, nil
}