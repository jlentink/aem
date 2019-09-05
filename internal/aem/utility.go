package aem

import (
	"fmt"
	"github.com/jlentink/aem/internal/aem/objects"
	"github.com/jlentink/aem/internal/cli/project"
	"github.com/jlentink/aem/internal/output"
	"os/user"
	"regexp"
)

var (
	u *user.User
)

// WriteLicense Write license file to disk for instance
func WriteLicense(i *objects.Instance, c *objects.Config) (int, error) {
	licensePath, _ := project.GetLicenseLocation(*i)
	if !project.Exists(licensePath) && len(c.LicenseCustomer) > 0 {
		output.Printf(output.VERBOSE, "No license in place found one in config writing\n")
		license := fmt.Sprintf(
			"# Adobe Granite License Properties\n"+
				"license.product.name=Adobe Experience Manager\n"+
				"license.customer.name=%s\n"+
				"license.product.version=%s\n"+
				"license.downloadID=%s\n", c.LicenseCustomer, c.LicenseVersion, c.LicenseDownloadID)
		return project.WriteTextFile(licensePath, license)
	}
	return 0, nil
}

func getCurrentUser() (*user.User, error) {
	if u != nil {
		return u, nil
	}

	return user.Current()
}
func isURL(s string) bool {
	if len(s) < 7 {
		return false
	}

	r, _ := regexp.Compile(`(?i)^http(s?)://`)
	return r.MatchString(s)
}
