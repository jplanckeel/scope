package config

type Flags struct {
	Dryrun                bool
	InsecureSkipTLSverify bool
	PasswordFromStdinOpt  bool
	SourceFile            string
	Registry              string
	Scheme                string
	Namespace             string
	Username              string
	Password              string
	CaFile                string
	CertFile              string
	KeyFile               string
}
