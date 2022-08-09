package main

import (
	"os"

	"github.com/fatih/color"
)

func main() {
	exist := fileExists("/Users/marityr/Git/gogit/app/.git")
	if exist {
		color.Cyan("+++")
	} else {
		color.Red("---")
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
