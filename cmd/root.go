package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/wonjongin/gacp/git"
)

var ShouldPush bool

var rootCmd = &cobra.Command{
	Use:   "gacp [flags] [commit_msg]",
	Short: "Git add commit push automatically",
	Long:  `Gacp is a cammand line tool for developer which use git and github. Gacp can git add, git commit, git push at once.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires to type commit message: gacp \"commit message\"")
		}
		return nil

	},
	Run: func(c *cobra.Command, args []string) {
		if ShouldPush {
			git.GitAddCommitPush(args[0])
		} else {
			git.GitAddCommit(args[0])
			if git.AskGitPush() {
				git.GitPush()
			}
		}
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&ShouldPush, "push", "p", false, "Run git push")
}
