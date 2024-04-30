package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func GormGen() {
	prompt := promptui.Prompt{
		Label: "Enter your name",
	}

	result, _ := prompt.Run()
	fmt.Printf("Hello, %s!\n", result)
}
