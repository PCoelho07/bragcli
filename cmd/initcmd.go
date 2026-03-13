package cmd

import (
	"brag/internal/brag"
	"brag/internal/store"

	"github.com/spf13/cobra"
)

func CreateInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		Short:   "Initialize a new bragdoc",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			s, err := store.NewDiskStore(name)
			if err != nil {
				return err
			}
			return brag.New(s).Init()
		},
	}

	cmd.Flags().StringP("name", "n", "", "name of the bragdoc")
	cmd.MarkFlagRequired("name")

	return cmd
}
