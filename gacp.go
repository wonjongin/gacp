package main

import (
	"github.com/wonjongin/gacp/cmd"
)

func main() {
	// 	if len(os.Args) != 2 {
	// 		fmt.Println(aurora.Red(`
	// 문법에 맞지 않아요. 다음처럼 입력하세요.
	// gacp "커밋 메시지"
	// 		`))
	// 		os.Exit(1)
	// 	}
	// 	commitMsg := os.Args[1]

	// 	git.GitAddCommitPush(commitMsg)

	cmd.Execute()
}
