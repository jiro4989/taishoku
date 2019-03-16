package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/jiro4989/taishoku/log"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	RootCommand.PersistentFlags().BoolVarP(&log.DebugFlag, "debug", "X", false, "debug logging flag.")
	RootCommand.Flags().SortFlags = false
	RootCommand.Flags().IntP("year", "y", 2999, "year")
	RootCommand.Flags().IntP("month", "m", 12, "month")
	RootCommand.Flags().IntP("day", "d", 31, "day")
	RootCommand.Flags().StringP("department", "D", "ジョークコマンド開発部", "your department")
	RootCommand.Flags().StringP("team", "T", "第一課", "your team")
	RootCommand.Flags().StringP("your-name", "n", "真面目田マジメ", "your name")
	RootCommand.Flags().StringP("company", "C", "株式会社ジョークコマンド", "company name")
	RootCommand.Flags().StringP("president", "P", "代表取締役", "president")
	RootCommand.Flags().StringP("president-name", "N", "ジョーク山悪ふざけ太郎", "president name")
	// RootCommand.Flags().Bool("html", false, "output html")
}

var RootCommand = &cobra.Command{
	Use:   "taishoku",
	Short: "退職届を出力する",
	Long: `
taishoku prints Japanese one's resignation ASCII Art.

Repository: https://github.com/jiro4989/taishoku
    Author: jiro4989

Example:

　　　代株　　　　二退こ　　　
　　　表式　　　　〇職の　　　
　　　取会　　　　一いた　　　
　　　締社　　　　九たび　　退
　　　役ジ　　　　年し一　　　
　　　　ョ　　　　〇た身　　職
　　　ジー　　　　三く上　　　
　　　ョク　　　　月この　　願
　　　ーコ　　　　一こ都　　　
　　　クマ　　　　七に合　　　
　　　山ン　　　　日おに　　　
　　　悪ド　　　　　願よ　　　
　　　ふ　　　　　　いり　　　
　　　ざ　　　　　　申、　　　
　　　け　　真ジ　　し二　　　
　　　太　　面ョ　　上九　　　
　　　郎　　目ー　　げ九　　　
　　　　　　田ク　　ま九　　　
　　　　　　マコ　　す年　　　
　　　　　　ジマ　　。一　　　
　　　　　　メン　　　二　　　
　　　　　　　ド　　　月　　　
　　　　　　印開　　　三　　　
　　　　　　　発　　　一　　　
　　　　　　　部　　　日　　　
　　　　　　　　　　　を　私　
　　　　　　　第　　　も　議　
　　　　　　　一　　　っ　　　
　　　　　　　課　　　て　　　

	`,
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
		arr := strings.Split(text, "\n")
		arr = padSpace(arr)
		printVertical(arr, 3)

		log.Debug("end 'taishoku'")
	},
}

func printVertical(s []string, padSize int) {
	// 一番上に空白をもうけたい
	for i := 0; i < padSize; i++ {
		fmt.Println()
	}

	leftPad := strings.Repeat("　", padSize)
	vertical := toVertical(s)
	for _, v := range vertical {
		fmt.Println(leftPad + v)
	}

	// 一番下に空白をもうけたい
	for i := 0; i < padSize; i++ {
		fmt.Println()
	}
}
