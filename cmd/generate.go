package cmd

import (
	"fmt"
	"strings"

	"github.com/gkwa/hoppinghare/internal/boilerplate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	templateURL    string
	outputFolder   string
	vars           []string
	varFiles       []string
	nonInteractive bool
	disableHooks   bool
	disableShell   bool
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a project from a template",
	Long: `Generate a project from a template using Gruntwork's Boilerplate.
This command processes a template at the given URL or path and outputs
the generated files to the specified output folder.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Convert vars string slice to map
		varsMap, err := parseVars(vars)
		if err != nil {
			return fmt.Errorf("failed to parse variables: %w", err)
		}

		// Merge with any vars from config file
		if viper.IsSet("vars") {
			configVars := viper.GetStringMapString("vars")
			for k, v := range configVars {
				// Command line takes precedence
				if _, exists := varsMap[k]; !exists {
					varsMap[k] = v
				}
			}
		}

		// Prepare options
		opts := boilerplate.Options{
			TemplateURL:             templateURL,
			OutputFolder:            outputFolder,
			NonInteractive:          nonInteractive,
			Vars:                    varsMap,
			VarFiles:                varFiles,
			MissingKeyAction:        viper.GetString("missing_key_action"),
			MissingConfigAction:     viper.GetString("missing_config_action"),
			DisableHooks:            disableHooks,
			DisableShell:            disableShell,
			DisableDependencyPrompt: viper.GetBool("disable_dependency_prompt"),
		}

		// Generate the project
		return boilerplate.Generate(opts)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Command-specific flags
	generateCmd.Flags().StringVar(&templateURL, "template-url", "", "URL or path to the template (required)")
	generateCmd.Flags().StringVar(&outputFolder, "output-folder", "", "folder where generated files will be written (required)")
	generateCmd.Flags().StringSliceVar(&vars, "var", []string{}, "specify a variable in the format NAME=VALUE (can be used multiple times)")
	generateCmd.Flags().StringSliceVar(&varFiles, "var-file", []string{}, "load variables from a YAML file (can be used multiple times)")
	generateCmd.Flags().BoolVar(&nonInteractive, "non-interactive", false, "do not prompt for input variables")
	generateCmd.Flags().BoolVar(&disableHooks, "disable-hooks", false, "do not run template hooks")
	generateCmd.Flags().BoolVar(&disableShell, "disable-shell", false, "do not run shell commands in templates")

	// Required flags
	generateCmd.MarkFlagRequired("template-url")
	generateCmd.MarkFlagRequired("output-folder")

	// Set default values from config
	viper.SetDefault("missing_key_action", "error")
	viper.SetDefault("missing_config_action", "exit")
	viper.SetDefault("disable_dependency_prompt", false)
}

// parseVars converts a slice of KEY=VALUE strings to a map
func parseVars(vars []string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for _, v := range vars {
		parts := strings.SplitN(v, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid variable format: %s (expected NAME=VALUE)", v)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key == "" {
			return nil, fmt.Errorf("variable name cannot be empty")
		}

		result[key] = value
	}

	return result, nil
}
