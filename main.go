package main

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ldez/go-git-cmd-wrapper/v2/clone"
	"github.com/ldez/go-git-cmd-wrapper/v2/git"
	giturls "github.com/whilp/git-urls"
)

type repo struct {
	gitHost     string
	projectName string
	userName    string
	URL         string
}

func main() {
	srcPath := os.Getenv("GCLONE_SRC_PATH")
	if len(srcPath) == 0 {
		fmt.Println("GCLONE_SRC_PATH is not set")
		os.Exit(1)
	}

	r, err := parseUrl(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	projectPath := fmt.Sprintf("%s/%s/%s/%s", srcPath, r.gitHost, r.userName, r.projectName)
	err = os.MkdirAll(projectPath, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := git.Clone(clone.Repository(r.URL), clone.Directory(projectPath), clone.Verbose)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(out)
}

func parseUrl(repoUrl string) (*repo, error) {
	var u *url.URL
	var err error

	if strings.HasPrefix(repoUrl, "git") {
		u, err = giturls.Parse(repoUrl)
		if err != nil {
			return nil, err
		}
	}

	if strings.HasPrefix(repoUrl, "https") {
		u, err = url.Parse(repoUrl)
		if err != nil {
			return nil, err
		}
	}

	basename := path.Base(u.Path)
	r := &repo{
		URL:         u.String(),
		gitHost:     u.Host,
		projectName: strings.TrimSuffix(basename, filepath.Ext(basename)),
		userName:    path.Dir(u.Path),
	}

	return r, nil
}
