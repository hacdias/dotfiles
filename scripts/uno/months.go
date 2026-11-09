package uno

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

// AddMonthFlags adds the --from and --to flags read by MonthFlags.
func AddMonthFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("from", "f", "", "Start date to start fetching (YYYY-MM).")
	cmd.Flags().StringP("to", "t", "", "End date to start fetching (YYYY-MM, not included).")
}

// MonthFlags returns --from and --to as UTC times, zero when unset.
func MonthFlags(cmd *cobra.Command) (from, to time.Time, err error) {
	fromStr, _ := cmd.Flags().GetString("from")
	toStr, _ := cmd.Flags().GetString("to")

	if fromStr != "" {
		from, err = time.ParseInLocation("2006-01", fromStr, time.UTC)
		if err != nil {
			return
		}
	}

	if toStr != "" {
		to, err = time.ParseInLocation("2006-01", toStr, time.UTC)
	}

	return
}

// SaveByMonth writes items to dir as one YYYY-MM.json file per UTC month,
// replacing the files of the months it has items for. UTC matches MonthFlags.
func SaveByMonth[T any](dir string, items []T, at func(T) time.Time) error {
	months := map[string][]T{}
	for _, item := range items {
		month := at(item).UTC().Format("2006-01")
		months[month] = append(months[month], item)
	}

	for month, items := range months {
		bytes, err := json.MarshalIndent(items, "", "\t")
		if err != nil {
			return err
		}

		err = os.WriteFile(filepath.Join(dir, month+".json"), bytes, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
