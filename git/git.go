package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/logrusorgru/aurora"
)

// FindGitDir 함수는 깃의 위치를 찾는다.
func FindGitDir() string {
	locGitNum := 0
	locGitPath := ".git"
	pwd, _ := os.Getwd()
	numOfPath := strings.Count(pwd, "/")

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

	return locGitPath
}

// GitAdd 함수는 git add 명령어를 실행한다.
func GitAdd() {
	locGitPath := FindGitDir()
	cmdGitAdd := exec.Command("git", "add", strings.Replace(locGitPath, ".git", "", 1))
	cmdGitAddOut, cmdGitAddErr := cmdGitAdd.Output()
	var stderr bytes.Buffer
	cmdGitAdd.Stderr = &stderr
	if cmdGitAddErr != nil {
		fmt.Println(cmdGitAddErr)
		fmt.Println(stderr.String())
		fmt.Println(aurora.Red("add 에러"))
		os.Exit(1)
	} else {
		fmt.Println(string(cmdGitAddOut))
		fmt.Println(aurora.Green("add 완료"))
	}
}

// GitCommit 함수는 git commit 명령어를 실행한다.
func GitCommit(commitMsg string) {
	cmdGitCommit := exec.Command("git", "commit", "-m", commitMsg)
	cmdGitCommitOut, cmdGitCommitErr := cmdGitCommit.Output()
	var stderr bytes.Buffer
	cmdGitCommit.Stderr = &stderr
	if cmdGitCommitErr != nil {
		fmt.Println(cmdGitCommitErr)
		fmt.Println(stderr.String())
		fmt.Println(aurora.Red("commit 에러"))
		os.Exit(1)
	} else {
		fmt.Println(string(cmdGitCommitOut))
		fmt.Println(aurora.Green("commit 완료"))
	}
}

// GitPush 함수는 git push 명령어를 실행합니다.
func GitPush() {
	isPush := true
	promptPush := &survey.Confirm{
		Message: "Do you want to push? ",
		Default: true,
	}
	survey.AskOne(promptPush, &isPush)
	if isPush == true {
		cmdGitPush := exec.Command("git", "push")
		cmdGitPushOut, cmdGitPushErr := cmdGitPush.Output()
		var stderr bytes.Buffer
		cmdGitPush.Stderr = &stderr
		if cmdGitPushErr != nil {
			fmt.Println(cmdGitPushErr)
			fmt.Println(stderr.String())
			fmt.Println(aurora.Red("push 에러"))
			os.Exit(1)
		} else {
			fmt.Println(string(cmdGitPushOut))
			fmt.Println(aurora.Green("push 완료"))
		}
	}
}

// GitAddCommitPush 함수는 git add, commit, push 를 실행합니다.
func GitAddCommitPush(commitMsg string) {
	GitAdd()
	GitCommit(commitMsg)
	GitPush()
}
