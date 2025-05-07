package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gkwa/hoppinghare/internal/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var templateDir string

// listTemplatesCmd represents the list-templates command
var listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "List available templates",
	Long: `List templates available in the configured templates directory.
Templates are directories containing a boilerplate.yml file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if templateDir == "" {
			// Try to get template directory from config
			templateDir = viper.GetString("template_directory")
			if templateDir == "" {
				return fmt.Errorf("template directory not specified and not configured")
			}
		}

		// Resolve to absolute path
		absTemplateDir, err := filepath.Abs(templateDir)
		if err != nil {
			return fmt.Errorf("failed to resolve template directory path: %w", err)
		}

		log.Info("Searching for templates in %s", absTemplateDir)
		
		// Find templates (directories containing boilerplate.yml)
		templates, err := findTemplates(absTemplateDir)
		if err != nil {
			return fmt.Errorf("failed to find templates: %w", err)
		}

		// Print templates
		if len(templates) == 0 {
			fmt.Println("No templates found in", absTemplateDir)
			return nil
		}

		fmt.Println("Available templates:")
		for _, t := range templates {
			relPath, err := filepath.Rel(absTemplateDir, t)
			if err != nil {
				relPath = t
			}
			fmt.Println("-", relPath)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listTemplatesCmd)

	// Command-specific flags
	listTemplatesCmd.Flags().StringVar(&templateDir, "template-dir", "", "directory containing templates (default from config)")
}

// findTemplates finds all directories containing a boilerplate.yml file
func findTemplates(rootDir string) ([]string, error) {
	templates := []string{}

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Warn("Error accessing path %s: %v", path, err)
			return nil // Continue walking
		}

		if !info.IsDir() && info.Name() == "boilerplate.yml" {
			templates = append(templates, filepath.Dir(path))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return templates, nil
}

