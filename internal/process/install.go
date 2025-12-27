package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Askinstall() bool {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Install binary doberman to /usr/local/bin ? [Y/n]")

	answer, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	answer = strings.TrimSpace(strings.ToLower(answer))

	return answer == "" || answer == "y" || answer == "yes"
}

func Install(src, dst string) error {

	input, err := os.ReadFile(src)

	if err != nil {
		return err

	}
	err = os.WriteFile(dst, input, 0755)
	if err != nil {
		return err
	}

	return nil
}
