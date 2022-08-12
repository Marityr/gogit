/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

var (
	pull     = flag.Bool("pull", true, "automatically pull changes")
	push     = flag.Bool("push", true, "automatically push changes")
	author   = flag.String("author", "gitomatic", "author name for git commits")
	email    = flag.String("email", "gitomatic@fribbledom.com", "email address for git commits")
	interval = flag.String("interval", "1m", "how often to check for changes")
	privkey  = flag.String("privkey", "~/.ssh/id_rsa", "location of private key used for auth")
	username = flag.String("username", "", "username used for auth")
	password = flag.String("password", "", "password used for auth")
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse [string to parse]",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("usage: parse <path>")
			os.Exit(1)
		}
		files, err := ioutil.ReadDir(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if file.IsDir() {
				exist := fileExists(args[0] + "/" + file.Name() + "/.git")
				if exist {
					res := GitStatus(string(args[0] + "/" + file.Name()))
					color.Green("[%v] %v", res, file.Name())
				} else {
					color.White("--- %v", file.Name())
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GitStatus(dir string) int {
	allit := 0

	path := dir

	for {
		log.Println("Checking repository:", path)
		r, err := git.PlainOpen(path)
		if err != nil {
			fatal("cannot open repository: %s\n", err)
		}
		w, err := r.Worktree()
		if err != nil {
			fatal("cannot access repository: %s\n", err)
		}

		if *push {
			status, err := w.Status()
			if err != nil {
				fatal("cannot retrieve git status: %s\n", err)
			}

			changes := 0
			// msg := ""
			//TODO парсинг статуса файлов в каталоге с репозиториями - отдать общее число изменений в каталоге
			for path, s := range status {
				switch s.Worktree {
				case git.Modified:
					changes++
				case git.Deleted:
					changes++
				default:
					changes++
				}
				allit = changes
			}
		}

	}
	return allit
}

func fatal(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}

func gitAdd(w *git.Worktree, path string) error {
	log.Printf("Adding file to work-tree: %s\n", path)
	_, err := w.Add(path)
	return err
}

func gitHasRemote(r *git.Repository) bool {
	remotes, _ := r.Remotes()
	return len(remotes) > 0
}

func gitRemove(w *git.Worktree, path string) error {
	log.Printf("Removing file from work-tree: %s\n", path)
	_, err := w.Remove(path)
	return err
}

func gitPush(r *git.Repository, auth transport.AuthMethod) error {
	if !gitHasRemote(r) {
		log.Println("Not pushing: no remotes configured.")
		return nil
	}

	log.Println("Pushing changes...")
	return r.Push(&git.PushOptions{
		Auth: auth,
	})
}

func gitPull(r *git.Repository, w *git.Worktree, auth transport.AuthMethod) error {
	if !gitHasRemote(r) {
		log.Println("Not pulling: no remotes configured.")
		return nil
	}

	log.Println("Pulling changes...")
	err := w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth:       auth,
	})
	if err == transport.ErrEmptyRemoteRepository {
		return nil
	}
	if err == git.NoErrAlreadyUpToDate {
		return nil
	}
	return err
}

func parseAuthArgs() (transport.AuthMethod, error) {
	if len(*username) > 0 {
		return &http.BasicAuth{
			Username: *username,
			Password: *password,
		}, nil
	}

	*privkey, _ = homedir.Expand(*privkey)
	auth, err := ssh.NewPublicKeysFromFile("git", *privkey, "")
	if err != nil {
		return nil, err
	}
	return auth, nil
}

//TODO https://github.com/muesli/gitomatic/blob/master/main.go
