package init

type configAnswers struct {
	UseKeyRing         bool
	MorePackages       bool
	Lazybones          bool
	LazybonesTemplate  string
	JarURL             string
	JarLocationType    string
	JarLocation        string
	JarUsername        string
	JarPassword        string
	LicenseCustomer    string
	LicenseVersion     string
	LicenseDownloadID  string
	AdditionalPackage  string
	AdditionalPackages []string
}
