package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func main()  {
	// commitMsg := ""
	// fmt.Print(aurora.Blue("커밋메시지(취소:q): "))
	// fmt.Scanln(&commitMsg)
	// fmt.Println(commitMsg)

	name := ""
	prompt := &survey.Input{
    	Message: "ping",
	}
	survey.AskOne(prompt, &name)
	fmt.Println(name)
}