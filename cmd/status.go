package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func infoOS() (map[string]string, error) {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return nil, err
	}

	info := make(map[string]string)
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := strings.Trim(parts[1], `"`)

		info[key] = value
	}

	return info, nil
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show system status",
	RunE: func(cmd *cobra.Command, args []string) error {

		now := time.Now().Format("15:04:05")

		uptimeRaw, err := os.ReadFile("/proc/uptime")
		if err != nil {
			return err
		}

		uptimeFields := strings.Fields(string(uptimeRaw))
		uptimeSeconds := uptimeFields[0]

		sec, err := time.ParseDuration(uptimeSeconds + "s")
		if err != nil {
			return err
		}

		uptime := fmt.Sprintf(
			"--> Device on %dh%dm <--",
			int(sec.Hours()),
			int(sec.Minutes())%60,
		)
		osInfo, err := infoOS()
		if err != nil {
			return err
		}
		osName := osInfo["PRETTY_NAME"]
		if osName == "" {
			osName = osInfo["NAME"]
		}

		fmt.Printf(
			"%s\n --> %s <-- Are you still awake?\n  %s\n",
			uptime,
			now,
			osName,
		)

		return nil

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
