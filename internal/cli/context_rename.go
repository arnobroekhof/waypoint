package cli

import (
	"strings"

	"github.com/posener/complete"

	"github.com/hashicorp/waypoint/internal/pkg/flag"
	"github.com/hashicorp/waypoint/sdk/terminal"
)

type ContextRenameCommand struct {
	*baseCommand
}

func (c *ContextRenameCommand) Run(args []string) int {
	// Initialize. If we fail, we just exit since Init handles the UI.
	flagSet := c.Flags()
	if err := c.Init(
		WithArgs(args),
		WithFlags(flagSet),
		WithNoConfig(),
	); err != nil {
		return 1
	}
	args = flagSet.Args()

	if len(args) != 2 {
		c.ui.Output(c.Flags().Help(), terminal.WithErrorStyle())
		return 1
	}

	if err := c.contextStorage.Rename(args[0], args[1]); err != nil {
		c.ui.Output(err.Error(), terminal.WithErrorStyle())
		return 1
	}

	c.ui.Output("Context %q has been renamed to %q.", args[0], args[1], terminal.WithSuccessStyle())
	return 0
}

func (c *ContextRenameCommand) Flags() *flag.Sets {
	return c.flagSet(0, nil)
}

func (c *ContextRenameCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *ContextRenameCommand) AutocompleteFlags() complete.Flags {
	return c.Flags().Completions()
}

func (c *ContextRenameCommand) Synopsis() string {
	return "Rename a context."
}

func (c *ContextRenameCommand) Help() string {
	helpText := `
Usage: waypoint context rename [options] FROM TO

  Rename a context FROM to TO.

  This will error if FROM does not exist. This will overwrite TO if it
  exists. If the current default is FROM, the default will be set to TO.

` + c.Flags().Help()

	return strings.TrimSpace(helpText)
}