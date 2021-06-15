package repo

import (
	"net/url"
	"path"
	"path/filepath"
	"strings"

	giturls "github.com/whilp/git-urls"
)

// Repo respresents a git repository
type Repo struct {
	GitHost     string
	ProjectName string
	UserName    string
	URL         string
}

// New creates a new git Repo using the given url
func New(repoURL string) (*Repo, error) {
	var u *url.URL
	var err error

	if strings.HasPrefix(repoURL, "git") {
		u, err = giturls.Parse(repoURL)
		if err != nil {
			return nil, err
		}
	}

	if strings.HasPrefix(repoURL, "https") {
		u, err = url.Parse(repoURL)
		if err != nil {
			return nil, err
		}
	}

	basename := path.Base(u.Path)
	r := &Repo{
		URL:         u.String(),
		GitHost:     u.Host,
		ProjectName: strings.TrimSuffix(basename, filepath.Ext(basename)),
		UserName:    strings.TrimPrefix(path.Dir(u.Path), "/"),
	}

	return r, nil
}
