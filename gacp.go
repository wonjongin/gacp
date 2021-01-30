package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/logrusorgru/aurora"
)

func setupUI(){
	mainwin := ui.NewWindow("GACP", 400, 200, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	grid := ui.NewGrid()
	grid.SetPadded(true)

	editor := ui.NewMultilineEntry()
	btnGo := ui.NewButton("커밋")
	btnGo.OnClicked(func(*ui.Button){
		commitMsg := editor.Text()
		fmt.Println(commitMsg)
		GitAdd()
		GitCommit(commitMsg)
		GitPush()
	})
	btnCancel := ui.NewButton("최소")
	btnCancel.OnClicked(func(*ui.Button){
		ui.Quit()
	})

	grid.Append(ui.NewLabel("커밋메시지를 입력하세요."), 0, 0, 4, 1, false, ui.AlignCenter, false, ui.AlignFill)
	grid.Append(editor, 0, 1, 4, 3, false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(btnGo, 2, 4, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(btnCancel, 3, 4, 1, 1, false, ui.AlignFill, false, ui.AlignFill)


	mainwin.SetChild(grid)

	mainwin.Show()

}

// CheckGitDir 함수는 깃의 위치를 찾는다.
func CheckGitDir() string{
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
func GitAdd(){
	locGitPath := CheckGitDir()
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
func GitCommit(commitMsg string){
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
func GitPush(){
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

func main() {
	fmt.Println(aurora.Bold(`
    __ _   __ _   ___  _ __  
   / _' | / _' | / __|| '_ \ 
  | (_| || (_| || (__ | |_) |
   \__, | \__,_| \___|| .__/ 
    __/ |             | |    
   |___/              |_|    
  `))
	  

	ui.Main(setupUI);

	

	// prompt := &survey.Input{
	// 	Message: "Commit Message (cancel: q):",
	// }
	// survey.AskOne(prompt, &commitMsg)
	// fmt.Print(aurora.Blue("커밋메시지(취소:q): "))
	// fmt.Scanln(&commitMsg)
	// if commitMsg == "q" {
	// 	os.Exit(0)
	// 	fmt.Println(aurora.Blue("취소합니다."))
	// }
}
