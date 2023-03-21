package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"flag"
	"github.com/cli/go-gh"
)

const (
	ROOT = "src"
	VCS  = "github.com"
)

func main() {
	l := log.New(os.Stderr, "", 0)
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(0)
	}
	var owner, repo, nwo string
	if len(args) == 1 && strings.Contains(args[0], "/") {
		owner = strings.Split(args[0], "/")[0]
		repo = strings.Split(args[0], "/")[1]
	} else if len(args) == 2 {
		owner = args[0]
		repo = args[1]
	} else {
		flag.Usage()
		os.Exit(0)
	}
	nwo = fmt.Sprintf("%s/%s", owner, repo)
	home := os.Getenv("HOME")
	path := filepath.Join(home, ROOT, VCS, nwo)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			l.Fatal(err)
		}
		args := []string{"repo", "clone", nwo, path}
		l.Println(fmt.Sprintf("Cloning into '%s'...", path))
		_, stdErr, err := gh.Exec(args...)
		if err != nil {
			log.Fatal(err)
			l.Println(stdErr)
		}
	}
	fmt.Println(path)
}
