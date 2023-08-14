package main

import (
	"github.com/jplanckeel/scope/internal"

	"github.com/spf13/cobra"
)

var binaryHelm string
var dryrun bool
var config string
var registry string
var registryType string
var user string
var password string

var rootCmd = &cobra.Command{
	Version: "0.0.2",
	Use:     "scope",
	Short:   "a cli to sync helmchart to private registry",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Sync(binaryHelm, config, registry, registryType, user, password, dryrun)
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}

}

func init() {
	rootCmd.PersistentFlags().StringVarP(&binaryHelm, "binary", "b", "helm", "alias for binary helm3")
	rootCmd.PersistentFlags().BoolVarP(&dryrun, "dryrun", "d", false, "enable dry-run mode")
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "path to configfile")
	rootCmd.PersistentFlags().StringVarP(&registry, "registry", "r", "", "destination chart registry")
	rootCmd.PersistentFlags().StringVarP(&registryType, "registry-type", "t", "oci", "registry nexus or ecr (default: oci)")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "user for nexus registry")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password for nexus registry")
}
