package main

import (
	"github.com/jplanckeel/scope/internal"

	"github.com/spf13/cobra"
)

var binaryHelm string
var dryrun bool
var config string
var registry string

var rootCmd = &cobra.Command{
	Version: "0.0.1",
	Use:     "scope",
	Short:   "a cli to sync helmchart to private registry",
	Run: func(cmd *cobra.Command, args []string) {
		internal.Sync(binaryHelm, config, registry, dryrun)
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
}
