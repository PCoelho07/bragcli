package cmd

import (
	"brag/internal/brag"
	"brag/internal/store"

	"github.com/spf13/cobra"
)

func CreateCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Add a new entry to a bragdoc",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			title, _ := cmd.Flags().GetString("title")
			description, _ := cmd.Flags().GetString("description")

			s, err := store.NewDiskStore(name)
			if err != nil {
				return err
			}

			return brag.New(s).Create(title, description)
		},
	}

	cmd.Flags().StringP("name", "n", "", "name of the bragdoc")
	cmd.Flags().StringP("title", "t", "", "title of the brag item")
	cmd.Flags().StringP("description", "d", "", "description of the brag item")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("title")

	return cmd
}
