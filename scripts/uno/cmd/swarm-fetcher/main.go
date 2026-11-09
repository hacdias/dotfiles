package main

import (
	"fmt"
	"os"
	"uno"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(fetchCmd)

	uno.AddMonthFlags(fetchCmd)
}

var rootCmd = &cobra.Command{
	Use:               "swarm-fetcher",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var loginCmd = &cobra.Command{
	Use: "login",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := uno.ParseConfig[config]()
		if err != nil {
			return err
		}

		t, err := newSwarm(c)
		if err != nil {
			return err
		}

		return t.interactiveLogin(8050)
	},
}

var fetchCmd = &cobra.Command{
	Use: "fetch",
	RunE: func(cmd *cobra.Command, args []string) error {
		from, to, err := uno.MonthFlags(cmd)
		if err != nil {
			return err
		}

		c, err := uno.ParseConfig[config]()
		if err != nil {
			return err
		}

		t, err := newSwarm(c)
		if err != nil {
			return err
		}

		return t.fetchBetween(cmd.Context(), from, to)
	},
}
