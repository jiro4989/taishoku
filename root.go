package main

import (
	"fmt"
	"time"

	"github.com/jiro4989/taishoku/log"
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "taishoku",
	Short: "退職届を出力する",
	Long:  "退職届けを出力する",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("start 'taishoku'")
		f := cmd.Flags()

		year, err := f.GetInt("year")
		if err != nil {
			panic(err)
		}

		month, err := f.GetInt("month")
		if err != nil {
			panic(err)
		}

		day, err := f.GetInt("day")
		if err != nil {
			panic(err)
		}

		department, err := f.GetString("department")
		if err != nil {
			panic(err)
		}

		team, err := f.GetString("team")
		if err != nil {
			panic(err)
		}

		yourName, err := f.GetString("your-name")
		if err != nil {
			panic(err)
		}

		company, err := f.GetString("company")
		if err != nil {
			panic(err)
		}

		president, err := f.GetString("president")
		if err != nil {
			panic(err)
		}

		presidentName, err := f.GetString("president-name")
		if err != nil {
			panic(err)
		}

		taishokuDate := fmt.Sprintf("%d年%d月%d日", year, month, day)
		taishokuDate = convertStringNumberToKanji(taishokuDate)
		today := time.Now().Format("2006年01月02日")
		today = convertStringNumberToKanji(today)
		text := makeTaishokuText(taishokuDate, today, department, team, yourName, company, president, presidentName)
		fmt.Println(text)

		log.Debug("end 'taishoku'")
	},
}

func init() {
	cobra.OnInitialize()
	RootCommand.PersistentFlags().BoolVarP(&log.DebugFlag, "debug", "X", false, "Debug logging flag.")
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
