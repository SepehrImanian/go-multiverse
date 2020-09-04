package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/yondero/multiverse/ipfs"
	"github.com/yondero/multiverse/repo"
)

var logCmd = &cobra.Command{
	Use:          "log",
	Short:        "Log change history.",
	Long:         `Log change history.`,
	SilenceUsage: true,
	RunE:         executeLog,
}

func init() {
	rootCmd.AddCommand(logCmd)
}

func executeLog(cmd *cobra.Command, args []string) error {
	ipfs, err := ipfs.NewDefault(context.TODO())
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	r, err := repo.Open(cwd)
	if err != nil {
		return err
	}

	return r.Log(ipfs)
}
