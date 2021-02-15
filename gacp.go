package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/logrusorgru/aurora"
)

func main() {
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
		fmt.Println(aurora.Red("git 프로젝트가 아닙니다!"))
		os.Exit(1)
	}
	fmt.Println(".git 위치: ", locGitNum, locGitPath)
	cmdGitAdd := exec.Command("git", "add", strings.Replace(locGitPath, ".git", "", 1))
	cmdGitAddOut, cmdGitAddErr := cmdGitAdd.Output()
	if cmdGitAddErr != nil {
		fmt.Println(cmdGitAddErr)
		fmt.Println(aurora.Red("add 에러"))
		os.Exit(1)
	} else {
		fmt.Println(string(cmdGitAddOut))
		fmt.Println(aurora.Green("add 완료"))
	}

	prompt := &survey.Input{
		Message: "Commit Message (cancel: q):",
	}
	survey.AskOne(prompt, &commitMsg)
	// fmt.Print(aurora.Blue("커밋메시지(취소:q): "))
	// fmt.Scanln(&commitMsg)
	if commitMsg == "q" {
		os.Exit(0)
		fmt.Println(aurora.Blue("취소합니다."))
	}
	cmdGitCommit := exec.Command("git", "commit", "-m", commitMsg)
	cmdGitCommitOut, cmdGitCommitErr := cmdGitCommit.Output()
	if cmdGitCommitErr != nil {
		fmt.Println(cmdGitCommitErr)
		fmt.Println(aurora.Red("commit 에러"))
		os.Exit(1)
	} else {
		fmt.Println(string(cmdGitCommitOut))
		fmt.Println(aurora.Green("commit 완료"))
	}

	isPush := true
	promptPush := &survey.Confirm{
		Message: "Do you want to push? ",
		Default: true,
	}
	survey.AskOne(promptPush, &isPush)
	if isPush == true {
		cmdGitPush := exec.Command("git", "push")
		cmdGitPushOut, cmdGitPushErr := cmdGitPush.Output()
		if cmdGitPushErr != nil {
			fmt.Println(cmdGitPushErr)
			fmt.Println(aurora.Red("push 에러"))
			os.Exit(1)
		} else {
			fmt.Println(string(cmdGitPushOut))
			fmt.Println(aurora.Green("push 완료"))
		}
	}

}
