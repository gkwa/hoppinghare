package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// These variables are set during build time using -ldflags
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Print the version, build commit, and build date information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("HoppingHare v%s (commit: %s, built: %s)\n", version, commit, date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

