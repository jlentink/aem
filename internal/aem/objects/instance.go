package objects

import (
	"fmt"
	"github.com/jlentink/aem/internal/output"
	"github.com/jlentink/aem/internal/sliceutil"
	"github.com/zalando/go-keyring"
)

// Instance for aem instance
type Instance struct {
	Name              string   `toml:"name"`
	Aliases           []string `toml:"aliases"`
	Group             string   `toml:"group"`
	Debug             bool     `toml:"debug"`
	Protocol          string   `toml:"proto"`
	IP                string   `toml:"ip"`
	Hostname          string   `toml:"hostname"`
	Port              int      `toml:"port"`
	SPort             int      `toml:"secure-port"`
	DispatcherVersion string   `toml:"dispatcher-version"`
	Author            string   `toml:"author"`
	Publisher         string   `toml:"publisher"`
	Type              string   `toml:"type"`
	RunMode           string   `toml:"runmode"`
	Username          string   `toml:"username"`
	SSHUsername       string   `toml:"ssh-username"`
	Password          string   `toml:"password"`
	JVMOptions        []string `toml:"jvm-options"`
	JVMDebugOptions   []string `toml:"jvm-debug-options"`
	Version           string   `toml:"version"`
}

// URLString for instance
func (i *Instance) URLString() string {
	return fmt.Sprintf("%s://%s:%d", i.Protocol, i.Hostname, i.Port)
}

// URLString for instance
func (i *Instance) URLIPString() string {
	return fmt.Sprintf("%s://%s:%d", i.Protocol, i.IP, i.Port)
}

// GetPassword get password for instance
func (i *Instance) GetPassword() (string, error) {
	if Cnf.KeyRing {
		return keyring.Get(i.hostServiceName(), i.Username)
	}
	return i.Password, nil
}

// GetPasswordSimple Get password and not receive an error
func (i *Instance) GetPasswordSimple() string {
	var passwd string
	var err error
	if Cnf.KeyRing {
		passwd, err = keyring.Get(i.hostServiceName(), i.Username)
		if err != nil {
			output.Print(output.VERBOSE, "Failed retrieving password from keychain dropping back to")
			passwd = i.Password
		}
	}
	if passwd == "" {
		passwd = i.Password
	}
	return passwd
}

// SetPassword set password for instance
func (i *Instance) SetPassword(p string) error {
	return keyring.Set(i.hostServiceName(), i.Username, p)
}

func (i *Instance) hostServiceName() string {
	return fmt.Sprintf("%s-%s-%s-%s-%d", serviceName, Cnf.ProjectName, i.Name, i.Hostname, i.Port)
}

// GetVersion Get version for instance
func (i *Instance) GetVersion() string {
	if len(i.Version) > 0 {
		return i.Version
	}
	return Cnf.Version
}

// InstanceOf is instance of defined groups
func (i *Instance) InstanceOf(types []string) bool {
	if sliceutil.InSliceString(types, i.Type) {
		return true
	}
	return false
}
