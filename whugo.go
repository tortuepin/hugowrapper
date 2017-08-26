package main

import (
	"fmt"
	//	"os"

	"github.com/urfave/cli"
)

var commands = []cli.Command{
	{
		Name:    "new",
		Aliases: []string{"n"},
		Usage:   "create new post",
		Action:  cmdNew,
	},
	{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "post list",
		Action:  cmdList,
	},
	{
		Name:    "edit",
		Aliases: []string{"e"},
		Usage:   "edit post",
		Action:  cmdEdit,
	},
}

func cmdNew() {
	fmt.Print("new")
}
func cmdList() {
	fmt.Print("list")
}
func cmdEdit() {
	fmt.Print("Edit")
}

func main() {
	fmt.Print("hogehoge")
}
