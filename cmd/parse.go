/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse [string to parse]",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("usage: parse <path>")
			os.Exit(1)
		}
		files, err := ioutil.ReadDir(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if file.IsDir() {
				exist := fileExists(args[0] + "/" + file.Name() + "/.git")
				if exist {
					res := GitStatus(string(args[0] + "/" + file.Name()))
					color.Green("[%v] %v", res, file.Name())
				} else {
					color.White("--- %v", file.Name())
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GitStatus(dir string) int {
	return 0
}

//TODO https://github.com/muesli/gitomatic/blob/master/main.go
