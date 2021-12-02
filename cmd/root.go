package cmd

import (
	"errors"

	"github.com/fwwieffering/advent-of-code-2021/cmd/day"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "advent",
	Short: "Run advent of code days",
	RunE: func(cmd *cobra.Command, args []string) error {
		if dayInput == 0 || dayInput > 31 {
			return errors.New("day must be between 1-31")
		}
		return day.DayFunc[dayInput]()
	},
}

var dayInput int

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&dayInput, "day", "d", 0, "day to run")
}
