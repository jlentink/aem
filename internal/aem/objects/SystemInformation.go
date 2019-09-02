package objects

//SystemInformation descriptive struct
type SystemInformation struct {
	HealthChecks interface{} `json:"Health Checks"`
	Instance     struct {
		AdobeExperienceManager string `json:"Adobe Experience Manager"`
		RunModes               string `json:"Run Modes"`
		InstanceUpSince        string `json:"Instance Up Since"`
	} `json:"Instance"`
	Repository struct {
		ApacheJackrabbitOak string `json:"Apache Jackrabbit Oak"`
		NodeStore           string `json:"Node Store"`
		RepositorySize      string `json:"Repository Size"`
		FileDataStore       string `json:"File Data Store"`
	} `json:"Repository"`
	MaintenanceTasks  interface{} `json:"Maintenance Tasks"`
	SystemInformation struct {
		MacOSX            string `json:"Mac OS X"`
		Linux             string `json:"Linux"`
		Windows           string `json:"Windows"`
		CurrentOS         string
		SystemLoadAverage string `json:"System Load Average"`
		UsableDiskSpace   string `json:"Usable Disk Space"`
		MaximumHeap       string `json:"Maximum Heap"`
	} `json:"System Information"`
	EstimatedNodeCounts struct {
		Total         string `json:"Total"`
		Tags          string `json:"Tags"`
		Assets        string `json:"Assets"`
		Authorizables string `json:"Authorizables"`
		Pages         string `json:"Pages"`
	} `json:"Estimated Node Counts"`
	ReplicationAgents  interface{} `json:"Replication Agents"`
	DistributionAgents interface{} `json:"Distribution Agents"`
}
