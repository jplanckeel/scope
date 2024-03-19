package config

type Flags struct {
	Dryrun                bool
	InsecureSkipTLSverify bool
	PasswordFromStdinOpt  bool
	AppendSource          bool
	SourceFile            string
	Registry              string
	Scheme                string
	Type                  string
	Namespace             string
	Username              string
	Password              string
	CaFile                string
	CertFile              string
	KeyFile               string
	AppendSuffix          string
}
