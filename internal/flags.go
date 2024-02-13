package internal

type ScopeConfig struct {
	Dryrun       bool
	BinaryHelm   string
	ConfigFile   string
	Registry     string
	RegistryType string
	User         string
	Password     string
}
