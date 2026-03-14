package cmd

import "github.com/spf13/cobra"

func CreateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "brag",
		Short: "Your bragdoc at a distance command",
	}

	cmd.AddCommand(CreateInitCmd())
	cmd.AddCommand(CreateCreateCmd())
	cmd.AddCommand(CreateListCmd())

	return cmd
}
