/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"os"
	"os/exec"

	"github.com/Marityr/gogit/cmd"
)

func main() {
	cmdC := exec.Command("clear") //Linux example, its tested
	cmdC.Stdout = os.Stdout
	cmdC.Run()

	cmd.Execute()
}
