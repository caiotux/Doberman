package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}
}

func printHelp() {
	fmt.Println(Banner)
	fmt.Println(` 
	
Usage :
	doberman status
	doberman list 
	doberman new 
	
	`)
}
