package main

import (
	"github.com/jplanckeel/scope/internal"

	"github.com/spf13/cobra"
)

var config internal.ScopeConfig

var rootCmd = &cobra.Command{
	Version: "0.3.0",
	Use:     "scope",
	Short:   "a cli to sync helmchart to private registry",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Sync(config)
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}

}

func init() {
	rootCmd.PersistentFlags().StringVarP(&config.BinaryHelm, "binary", "b", "helm", "alias for binary helm3")
	rootCmd.PersistentFlags().BoolVarP(&config.Dryrun, "dryrun", "d", false, "enable dry-run mode")
	rootCmd.PersistentFlags().StringVarP(&config.ConfigFile, "config", "c", "", "path to configfile")
	rootCmd.PersistentFlags().StringVarP(&config.Registry, "registry", "r", "", "destination chart registry")
	rootCmd.PersistentFlags().StringVarP(&config.RegistryType, "registry-type", "t", "oci", "registry nexus or oci")
	rootCmd.PersistentFlags().StringVarP(&config.User, "user", "u", "", "user for nexus registry")
	rootCmd.PersistentFlags().StringVarP(&config.Password, "password", "p", "", "password for nexus registry")
}
