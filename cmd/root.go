package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dob",
	Short: "Doberman CLI - the canem of applications",
	Long: `Doberman CLI is a lightweight tool focused on automating
and monitoring local and remote applications through simple,
friendly, human-readable commands.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
