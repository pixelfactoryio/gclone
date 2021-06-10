package repo

import (
	"net/url"
	"path"
	"path/filepath"
	"strings"

	giturls "github.com/whilp/git-urls"
)

type Repo struct {
	GitHost     string
	ProjectName string
	UserName    string
	URL         string
}

func New(repoUrl string) (*Repo, error) {
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
	r := &Repo{
		URL:         u.String(),
		GitHost:     u.Host,
		ProjectName: strings.TrimSuffix(basename, filepath.Ext(basename)),
		UserName:    path.Dir(u.Path),
	}

	return r, nil
}
