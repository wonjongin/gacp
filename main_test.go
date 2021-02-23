package main

import (
	"fmt"
	"os"
)

func TestMain()  {
	// commitMsg := ""
	// fmt.Print(aurora.Blue("커밋메시지(취소:q): "))
	// fmt.Scanln(&commitMsg)
	// fmt.Println(commitMsg)

	// name := ""
	// prompt := &survey.Input{
    // 	Message: "ping",
	// }
	// survey.AskOne(prompt, &name)
	// fmt.Println(name)
	fmt.Println("0: " + os.Args[0])
	fmt.Println("1: " + os.Args[1])
	fmt.Println("2: " + os.Args[2])
}