package main

import (
	"github.com/go-git/go-git/v5"
	"fmt"
	"os"
	"flag"
	"log"
	"strings"
)


func CheckIfError(err error) {
	if err == nil {
		return
	}
	
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func Clone(url string) error{
	repo := strings.Split(url, "/")[1]
	_, err := git.PlainClone(repo, false, &git.CloneOptions{
		URL:      "https://github.com/" + url,
		Progress: os.Stdout,
	})
	return err
}

func CloneWithDir(url ,dir string) error{
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      "https://github.com/" + url,
		Progress: os.Stdout,
	})
	return err
}

func main() {
	flag.Parse()
	url := flag.Arg(0)
	dir := flag.Arg(1)
	if url == "" {
		log.Fatal("Please set repository  ex) kuroko1t/clone")
	}
	Info("git clone https://github.com/"+url)
	if dir == "" {
		err := Clone(url)
		CheckIfError(err)
	} else {
		err := CloneWithDir(url, dir)
		CheckIfError(err)
	}
}


