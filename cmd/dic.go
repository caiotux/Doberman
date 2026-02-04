package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/caiotux/Doberman/internal/gemini"
	"github.com/spf13/cobra"
)

var model string

var dicCmd = &cobra.Command{
	Use:   "dic [prompt]",
	Short: "Doberman CLI useing Gemini API to generate text based on a prompt",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var prompt string

		stat, err := os.Stdin.Stat()
		if err != nil {
			return err
		}

		hasStdin := (stat.Mode() & os.ModeCharDevice) == 0

		var stdinContent string
		if hasStdin {
			b, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			stdinContent = strings.TrimSpace(string(b))
		}

		if len(args) > 0 {
			prompt = strings.Join(args, " ")
		}

		switch {
		case stdinContent != "" && prompt != "":
			prompt = fmt.Sprintf(
				"Context:\n%s\n\nQuestion:\n%s",
				stdinContent,
				prompt,
			)

		case stdinContent != "":
			prompt = stdinContent
		}

		if prompt == "" {
			return fmt.Errorf("not prompt provided (args or stdin)")
		}

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
