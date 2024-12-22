package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	fetchCmd.Flags().StringP("from", "f", "", "Start date to start fetching (YYYY-MM).")
	fetchCmd.Flags().StringP("to", "t", "", "End date to start fetching (YYYY-MM, not included).")
}

var rootCmd = &cobra.Command{
	Use:               "lastfm-fetcher",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var fetchCmd = &cobra.Command{
	Use: "fetch",
	RunE: func(cmd *cobra.Command, args []string) error {
		fromStr, _ := cmd.Flags().GetString("from")
		toStr, _ := cmd.Flags().GetString("to")

		var (
			err      error
			from, to time.Time
		)

		if fromStr != "" {
			from, err = time.ParseInLocation("2006-01", fromStr, time.UTC)
			if err != nil {
				return err
			}
		}

		if toStr != "" {
			to, err = time.ParseInLocation("2006-01", toStr, time.UTC)
			if err != nil {
				return err
			}
		}

		c, err := parseConfig()
		if err != nil {
			return err
		}

		return newLastFm(c).fetchBetween(cmd.Context(), from, to)
	},
}
