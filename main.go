package main

import (
	"fmt"
	"github.com/cli/go-gh/v2/pkg/browser"
	"github.com/cli/go-gh/v2/pkg/repository"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("error: need one argument.")
	}
	// TODO: should validate arg to comply with https://github.com/mislav/hub/blob/master/commands/compare.go
	arg := os.Args[1]

	repo, err := repository.Current()
	if err != nil {
		log.Fatal(err)
	}
	url := fmt.Sprintf("https://%s/%s/%s/compare/%s", repo.Host, repo.Owner, repo.Name, arg)

	b := browser.New("", os.Stdout, os.Stderr)
	err = b.Browse(url)
	if err != nil {
		log.Fatal(err)
	}
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
