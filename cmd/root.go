package cmd

import (
	"os"

	"github.com/gkwa/hoppinghare/internal/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbosity int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hoppinghare",
	Short: "A tool for generating projects using Gruntwork's Boilerplate",
	Long: `HoppingHare is a tool that provides a convenient command-line interface
for generating projects using Gruntwork's Boilerplate as a library.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set log level based on verbosity flag
		for i := 0; i < verbosity; i++ {
			log.IncreaseLevel()
		}
		log.Debug("Verbosity level set to %d", verbosity)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hoppinghare.yaml)")
	rootCmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "increase verbosity (can be used multiple times)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			log.Error("Error determining user home directory: %s", err)
			os.Exit(1)
		}

		// Search config in home directory with name ".hoppinghare" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".hoppinghare")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file: %s", viper.ConfigFileUsed())
	}
}

