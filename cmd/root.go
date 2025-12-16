package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "doberman",
	Short: "Doberman CLI - the canem of applications",
	Long:  ` This is a lightweight CLI focused on automating and monitoring local and remote applications through simple, friendly, human-readable commands.`,
	Run: func(cmd *cobra.Command, args []string) {

		ShowBanner()
		fmt.Println(`
			Usage:
			dob status 
			dob list
			dob ports
			dob docker
		
		`)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
