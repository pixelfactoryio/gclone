package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ldez/go-git-cmd-wrapper/v2/clone"
	"github.com/ldez/go-git-cmd-wrapper/v2/git"
	"github.com/pixelfactoryio/gclone/pkg/repo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go.pixelfactory.io/pkg/version"
)

func initConfig() {
	viper.Set("revision", version.REVISION)
	viper.SetEnvPrefix("GCLONE")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}

// NewRootCmd create new rootCmd
func NewRootCmd() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:           "gclone <repo>",
		Short:         "gclone",
		SilenceErrors: true,
		PreRunE:       preStart,
		RunE:          start,
	}

	rootCmd.PersistentFlags().String("src-path", "", "source directory path")
	err := viper.BindPFlag("src-path", rootCmd.PersistentFlags().Lookup("src-path"))
	if err != nil {
		return nil, err
	}

	rootCmd.PersistentFlags().String("log-level", "info", "log level (debug, info, warn, error, fatal, panic)")
	err = viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))
	if err != nil {
		return nil, err
	}

	rootCmd.PersistentFlags().Bool("debug", false, "run in debug mode")
	err = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	if err != nil {
		return nil, err
	}

	return rootCmd, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	rootCmd, err := NewRootCmd()
	if err != nil {
		return err
	}

	cobra.OnInitialize(initConfig)
	return rootCmd.Execute()
}

func preStart(c *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("the <repo> argument is required")
	}

	if len(args) > 1 {
		return errors.New("the <repo> argument is required")
	}
	return nil
}

func start(c *cobra.Command, args []string) error {

	debug := viper.GetBool("debug")
	if debug {
		fmt.Println(viper.GetString("debug"), viper.GetString("src-path"), viper.GetString("log-level"), args)
	}

	srcPath := viper.GetString("src-path")
	if len(srcPath) == 0 {
		return errors.New("no source directory specified \nPlease set GCLONE_SRC_PATH or use `--src-path` flag")
	}

	r, err := repo.New(args[0])
	if err != nil {
		return err
	}

	projectPath := fmt.Sprintf("%s/%s/%s/%s", srcPath, r.GitHost, r.UserName, r.ProjectName)
	err = os.MkdirAll(projectPath, 0755)
	if err != nil {
		return err
	}

	out, err := git.Clone(clone.Repository(r.URL), clone.Directory(projectPath), clone.Verbose)
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}
