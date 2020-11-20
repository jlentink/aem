package objects

const (
	serviceName = `aemCLI`
)

// Cnf Active config instance
var Cnf *Config

// Config toml object to read config
type Config struct {
	ProjectName        string     `toml:"project-name"`
	Verbose            bool       `toml:"verbose"`
	Version            string     `toml:"version"`
	VersionSuffix      string     `toml:"version-suffix"`
	Packages           []string   `toml:"commandPullContent"`
	PackagesExcluded   []string   `toml:"packageExclude"`
	Command            string     `toml:"command,omitempty"`
	CommandArgs        []string   `toml:"command,omitempty"`
	DefaultInstance    string     `toml:"defaultInstance"`
	DefaultVersion     string     `toml:"default-version"`
	Instances          []Instance `toml:"instance"`
	JVMOptions         []string   `toml:"jvm-options"`
	JVMDebugOptions    []string   `toml:"jvm-debug-options"`
	AemJar             []AemJar   `toml:"aemJar"`
	LicenseCustomer    string     `toml:"licenseCustomer"`
	LicenseVersion     string     `toml:"licenseVersion"`
	LicenseDownloadID  string     `toml:"licenseDownloadID"`
	WatchPath          []string   `toml:"watchPath"`
	Port               int        `toml:"port"`
	Role               string     `toml:"role"`
	KeyRing            bool       `toml:"use-keyring"`
	JcrRoot            string     `toml:"jcrRoot"`
	JVMOpts            []string   `toml:"jvmOptions"`
	AdditionalPackages []string   `toml:"additionalPackages"`
	ContentPackages    []string   `toml:"contentPackages"`
	OakOptions         []string   `toml:"oakOptions"`
	OakVersion         string     `toml:"oakDefaultVersion"`
	VltPaths           []string   `toml:"vltSyncPaths"`
	BuildCommands      string     `toml:"buildCommand"`
	ValidateSSL        bool       `toml:"ssl-validate"`
	InvalidatePaths    []string   `toml:"invalidatePaths"`
	CloudManagerGit    string     `toml:"cloudManagerGit"`
	ContentBackupPaths []string   `toml:"contentBackupPaths"`
	ContentBackupName  string     `toml:"contentBackupName"`
	ContentBackupGroup string     `toml:"contentBackupGroup"`
}

// AemJar Descriptive jar
type AemJar struct {
	Location string `toml:"location"`
	Version  string `toml:"version"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

//func joinStrings(s []string) string {
//	returnString := ""
//	for _, path := range s {
//		returnString = returnString + "\t\"" + path + "\",\n"
//	}
//
//	return returnString
//}

//func Exists() bool {
//	path, _ := project.GetConfigFileLocation()
//	if project.Exists(path) {
//		return true
//	}
//	return false
//}

//func Render() string {
//	return fmt.Sprintf(config.template,
//		config.KeyChain,
//		config.JarLocation,
//		config.JarUsername,
//		config.JarPassword,
//		config.LicenseCustomer,
//		config.LicenseVersion,
//		config.LicenseDownloadID,
//		config.DefaultInstanceStr,
//		joinStrings(config.AdditionalPackages),
//	)
//}

//func WriteConfigFile() (int, error) {
//	p, err := project.GetConfigFileLocation()
//
//	if err != nil {
//		return 0, err
//	}
//	return project.WriteTextFile(p, Render())
//}

//func GetConfig() (*Config, error) {
//	p, err := project.GetConfigFileLocation()
//	if err != nil {
//		return nil, fmt.Errorf("could not find config file")
//	}
//
//	config := Config{}
//	_, err = toml.DecodeFile(p, &config)
//	if err != nil {
//		return nil, fmt.Errorf("could not decode config file: %s", err.Error())
//	}
//	return &config, nil
//}

// FindDefault Instance based on resolution order
//func DefaultInstance() string {
//	envName := os.Getenv(instanceEnv)
//	if len(envName) > 0 {
//		return envName
//	}
//
//	c, err := GetConfig()
//	if err != nil {
//		output.Printf(output.VERBOSE, "Error in  config returning default author")
//		return instanceMainDefault
//	}
//
//	if len(c.DefaultInstance) > 0 {
//		return c.DefaultInstance
//	}
//
//	return instanceMainDefault
//}
