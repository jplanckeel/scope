package cmd

import (
	"os"

	"github.com/jplanckeel/scope/pkg/config"
	"github.com/jplanckeel/scope/pkg/helm"
	"github.com/spf13/cobra"
)

var flags config.Flags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scope",
	Short: "Sync Chart On Private Registry",
	Run: func(cmd *cobra.Command, args []string) {
		helm.Sync(flags)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&flags.SourceFile, "source-file", "s", "", "path to source file")
	rootCmd.PersistentFlags().StringVarP(&flags.Registry, "registry", "r", "", "destination chart registry")
	rootCmd.PersistentFlags().StringVarP(&flags.Namespace, "namespace", "n", "", "namespace destination chart registry")
	rootCmd.PersistentFlags().StringVarP(&flags.Type, "type", "t", "oci", "type for registry (nexus or oci)")
	rootCmd.PersistentFlags().StringVarP(&flags.Username, "user", "u", "", "chart destination repository user")
	rootCmd.PersistentFlags().StringVarP(&flags.Password, "password", "p", "", "chart destination repository password")
	rootCmd.PersistentFlags().BoolVarP(&flags.PasswordFromStdinOpt, "password-stdin", "", false, "read password or identity token from stdin")
	rootCmd.PersistentFlags().StringVar(&flags.CertFile, "cert-file", "", "identify HTTPS client using this SSL certificate file")
	rootCmd.PersistentFlags().StringVar(&flags.KeyFile, "key-file", "", "identify HTTPS client using this SSL key file")
	rootCmd.PersistentFlags().StringVar(&flags.CaFile, "ca-file", "", "verify certificates of HTTPS-enabled servers using this CA bundle")
	rootCmd.PersistentFlags().BoolVar(&flags.InsecureSkipTLSverify, "insecure-skip-tls-verify", false, "skip tls certificate checks")
}
