package cmd

import "github.com/spf13/cobra"

func Register(c *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "folder",
		Short: "Gestão de pastas",
	}

	cmd.AddCommand(create())
	cmd.AddCommand(list())

	c.AddCommand(cmd)
}
