package main

import "github.com/spf13/cobra"

var debugFlag = false

var RootCommand = &cobra.Command{
	Use:   "taishoku",
	Short: "退職届を出力する",
	Long:  "退職届けを出力する",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cobra.OnInitialize()
	RootCommand.PersistentFlags().BoolVarP(&debugFlag, "debug", "X", false, "Debug logging flag.")
	RootCommand.Flags().IntP("year", "y", 2999, "year")
	RootCommand.Flags().IntP("month", "m", 12, "month")
	RootCommand.Flags().IntP("day", "d", 31, "day")
}
