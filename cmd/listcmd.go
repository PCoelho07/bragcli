package cmd

import (
	"brag/internal/brag"
	"brag/internal/store"
	"fmt"

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

			list, err := brag.New(s).List(name)
			if err != nil {
				return err
			}

			for _, l := range list {
				fmt.Printf("Title: %s, Description: %s, Created At: %s\n", l.Title, l.Description, l.CreatedAt.Format("01/02/2006 15:04:05"))
			}
			return nil
		},
	}

	cmd.Flags().StringP("name", "n", "", "name of the bragdoc")
	return cmd
}
