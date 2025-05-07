package boilerplate

import (
	"github.com/gruntwork-io/boilerplate/options"
	"github.com/gruntwork-io/boilerplate/templates"
	"github.com/gruntwork-io/boilerplate/variables"
	"github.com/gkwa/hoppinghare/internal/log"
)

// Options represents the configuration for a boilerplate run
type Options struct {
	// Template URL or path to use
	TemplateURL string

	// Where to output the processed files
	OutputFolder string

	// If true, don't prompt for input
	NonInteractive bool

	// Variables to pass to the template
	Vars map[string]interface{}

	// Variable files to load
	VarFiles []string

	// Action to take when template references a missing key
	// Valid values: "invalid", "zero", "error"
	MissingKeyAction string

	// Action to take when template folder doesn't contain boilerplate.yml
	// Valid values: "exit", "ignore"
	MissingConfigAction string

	// If true, don't run hooks
	DisableHooks bool

	// If true, shell helpers will return "replace-me" instead of executing
	DisableShell bool

	// If true, don't prompt for confirmation to include dependencies
	DisableDependencyPrompt bool
}

// Generate processes a boilerplate template with the given options
func Generate(opts Options) error {
	log.Debug("Processing template with options: %+v", opts)

	// Parse variables from var files
	vars, err := parseVarFiles(opts.VarFiles)
	if err != nil {
		return err
	}

	// Merge with command-line vars (command-line takes precedence)
	for k, v := range opts.Vars {
		vars[k] = v
	}

	// Convert our options to the internal options struct
	internalOpts, err := convertOptions(opts, vars)
	if err != nil {
		return err
	}

	// The root boilerplate.yml is not itself a dependency, so we pass an empty Dependency
	emptyDep := variables.Dependency{}

	// Process the template
	log.Info("Processing template: %s -> %s", opts.TemplateURL, opts.OutputFolder)
	err = templates.ProcessTemplate(internalOpts, internalOpts, emptyDep)
	if err != nil {
		return err
	}

	log.Info("Template processed successfully")
	return nil
}

// parseVarFiles parses variables from the specified var files
func parseVarFiles(files []string) (map[string]interface{}, error) {
	if len(files) == 0 {
		return make(map[string]interface{}), nil
	}

	vars, err := variables.ParseVars(nil, files)
	if err != nil {
		return nil, err
	}

	return vars, nil
}

// convertOptions converts our options struct to boilerplate's internal options
func convertOptions(opts Options, vars map[string]interface{}) (*options.BoilerplateOptions, error) {
	// Set default values for optional fields
	missingKeyAction := options.DefaultMissingKeyAction
	if opts.MissingKeyAction != "" {
		var err error
		missingKeyAction, err = options.ParseMissingKeyAction(opts.MissingKeyAction)
		if err != nil {
			return nil, err
		}
	}

	missingConfigAction := options.DefaultMissingConfigAction
	if opts.MissingConfigAction != "" {
		var err error
		missingConfigAction, err = options.ParseMissingConfigAction(opts.MissingConfigAction)
		if err != nil {
			return nil, err
		}
	}

	templateURL, templateFolder, err := options.DetermineTemplateConfig(opts.TemplateURL)
	if err != nil {
		return nil, err
	}

	// Create the internal options struct
	internalOpts := &options.BoilerplateOptions{
		TemplateUrl:             templateURL,
		TemplateFolder:          templateFolder,
		OutputFolder:            opts.OutputFolder,
		NonInteractive:          opts.NonInteractive,
		Vars:                    vars,
		OnMissingKey:            missingKeyAction,
		OnMissingConfig:         missingConfigAction,
		DisableHooks:            opts.DisableHooks,
		DisableShell:            opts.DisableShell,
		DisableDependencyPrompt: opts.DisableDependencyPrompt,
	}

	// Validate the options
	err = internalOpts.Validate()
	if err != nil {
		return nil, err
	}

	return internalOpts, nil
}

