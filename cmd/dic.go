package cmd

import (
	"fmt"

	"github.com/caiotux/Doberman/internal/gemini"
	"github.com/spf13/cobra"
)

var model string

var dicCmd = &cobra.Command{
	Use:   "dic [prompt]",
	Short: "Doberman CLI useing Gemini API to generate text based on a prompt",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		prompt := args[0]

		client, err := gemini.NewClient()
		if err != nil {
			return err
		}

		resp, err := client.Generate(cmd.Context(), prompt, model)
		if err != nil {
			return err
		}

		fmt.Println(resp)
		return nil
	},
}

func init() {
	dicCmd.Flags().StringVarP(
		&model,
		"model",
		"m",
		"gemini-2.5-flash",
		"Model to use for generation",
	)

	rootCmd.AddCommand(dicCmd)
}
