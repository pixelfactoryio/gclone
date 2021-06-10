package main

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/pixelfactoryio/gclone/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(aurora.Red("Error:"), err)
		os.Exit(1)
	}
}
