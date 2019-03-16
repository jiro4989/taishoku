package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jiro4989/taishoku/log"
	"github.com/spf13/cobra"
)

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

	l := utf8.RuneCountInString(s[0])
	leftPad := strings.Repeat("　", padSize)
	for i := 0; i < l; i++ {
		var line []string
		for _, row := range s {
			var colIdx int // rangeで指定する数値は配列の印字ではないため
			for _, c := range row {
				if colIdx == i {
					line = append(line, string(c))
				}
				colIdx++
			}
		}
		if len(line) == 0 {
			continue
		}
		line = reverse(line)
		s := strings.Join(line, "")
		fmt.Println(leftPad + s)
	}
	// 一番下に空白をもうけたい
	for i := 0; i < padSize; i++ {
		fmt.Println()
	}
}
