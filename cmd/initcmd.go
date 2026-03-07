package cmd

import (
	"brag/internal/brag"
	"brag/internal/fs"

	"github.com/spf13/cobra"
)

func CreateInitCmd() *cobra.Command {
    initcmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		RunE: func(cmd *cobra.Command, args []string) error {
            name, err := cmd.Flags().GetString("name")
            if err != nil {
                return err
            }

            disk := fs.NewDisk()
            b := brag.New(disk, brag.Options{ Name: name })

			return b.Init()
		},
	}

    initcmd.Flags().StringP("name", "n", "", "name of the bragdoc")

    return initcmd
}
