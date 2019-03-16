package main

import "github.com/spf13/cobra"

var debugFlag = false

var RootCommand = &cobra.Command{
	Use:   "taishoku",
	Short: "退職届を出力する",
	Long:  "退職届けを出力する",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cobra.OnInitialize()
	RootCommand.PersistentFlags().BoolVarP(&debugFlag, "debug", "X", false, "Debug logging flag.")
	RootCommand.Flags().IntP("year", "y", 2999, "year")
	RootCommand.Flags().IntP("month", "m", 12, "month")
	RootCommand.Flags().IntP("day", "d", 31, "day")
	RootCommand.Flags().StringP("department", "D", "ジョークコマンド開発部", "day")
	RootCommand.Flags().StringP("team", "T", "第一課", "day")
	RootCommand.Flags().StringP("your-name", "n", "真面目田マジメ", "day")
	RootCommand.Flags().StringP("company", "C", "株式会社ジョークコマンド", "day")
	RootCommand.Flags().StringP("president", "P", "代表取締役", "day")
	RootCommand.Flags().StringP("president-name", "N", "ジョーク山悪ふざけ太郎", "day")
	RootCommand.Flags().Bool("html", false, "output html")
}
