package repo_test

import (
	"testing"

	"github.com/pixelfactoryio/gclone/pkg/repo"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	r, err := repo.New("https://github.com/pixelfactoryio/gclone.git")
	is.NoError(err)
	is.NotEmpty(r)

	is.Equal(r.GitHost, "github.com")
	is.Equal(r.UserName, "pixelfactoryio")
	is.Equal(r.ProjectName, "gclone")
	is.Equal(r.URL, "https://github.com/pixelfactoryio/gclone.git")
}
