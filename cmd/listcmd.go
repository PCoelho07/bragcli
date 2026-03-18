package cmd

import (
	"brag/internal/brag"
	"brag/internal/presenter"
	"brag/internal/store"

	"github.com/spf13/cobra"
)

func CreateListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List your bragdoc items",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")

			s, err := store.NewDiskStore(name)
			if err != nil {
				return err
			}

			_, err = brag.New(s).List(presenter.NewTextPresenter())
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringP("name", "n", "", "name of the bragdoc")
	return cmd
}
