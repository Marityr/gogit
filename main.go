<<<<<<< HEAD
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
=======
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gogit",
		Usage: "fight the loneliness!",
		Action: func(*cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
>>>>>>> 09634b2632d58f0a923245e28d4795cca160ffa6
}
