package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Hello!")
	locGitNum := 0
	locGitPath := ".git"
	pwd, _ := os.Getwd()
	numOfPath := strings.Count(pwd, "/")
	commitMsg := ""

	if _, err := os.Stat(locGitPath); err == nil {
		locGitNum = 0
		locGitPath = "."
	} else {
		for i := 1; i < numOfPath; i++ {
			locGitPath = "../" + locGitPath
			if _, err := os.Stat(locGitPath); err == nil {
				break
				locGitNum = i
			}
		}
	}
	fmt.Println(".git 위치: ", locGitNum, locGitPath)
	cmdGitAdd := exec.Command("git", "add", strings.Replace(locGitPath, ".git", "", 1))
	cmdGitAddOut, cmdGitAddErr := cmdGitAdd.Output()
	if cmdGitAddErr != nil {
		fmt.Println(cmdGitAddErr)
	} else {
		fmt.Println(string(cmdGitAddOut))
	}

	fmt.Printf("커밋메시지: ")
	fmt.Scanln(&commitMsg)
	cmdGitCommit := exec.Command("git", "commit", "-m", commitMsg)
	cmdGitCommitOut, cmdGitCommitErr := cmdGitCommit.Output()
	if cmdGitCommitErr != nil {
		fmt.Println(cmdGitCommitErr)
	} else {
		fmt.Println(string(cmdGitCommitOut))
	}
	cmdGitPush := exec.Command("git", "push")
	cmdGitPushOut, cmdGitPushErr := cmdGitPush.Output()
	if cmdGitPushErr != nil {
		fmt.Println(cmdGitPushErr)
	} else {
		fmt.Println(string(cmdGitPushOut))
	}
}
