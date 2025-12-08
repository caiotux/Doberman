package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

//go:embed banner.ascii
var myBanner string

func ShowBanner() {
	const artWidth = 50

	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println(myBanner)
		return
	}

	paddingSize := (termWidth - artWidth) / 2
	if paddingSize < 0 {
		paddingSize = 0
	}

	padding := strings.Repeat(" ", paddingSize)

	lines := strings.Split(myBanner, "\n")
	for _, line := range lines {
		if line != "" {
			fmt.Printf("%s%s\n", padding, line)
		}
	}
}
