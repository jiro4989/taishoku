package main

import (
	"fmt"
	"strings"
)

const msg = `
このたび一身上の都合により、YYYY年MM月DD日をもって
退職いたしたくここにお願い申し上げます。
now

部署
氏名

社名
`

// https://tenshoku.mynavi.jp/knowhow/caripedia/58

type YearMonthDay struct {
	Year, Month, Day string
}

func main() {

}

func convertYYYYMMDDtoKanji(year, month, day int) YearMonthDay {
	convs := map[string]string{
		"0": "〇",
		"1": "一",
		"2": "二",
		"3": "三",
		"4": "四",
		"5": "五",
		"6": "六",
		"7": "七",
		"8": "八",
		"9": "九",
	}

	y := fmt.Sprintf("%d", year)
	for k, v := range convs {
		y = strings.Replace(y, k, v, -1)
	}
	m := fmt.Sprintf("%d", month)
	for k, v := range convs {
		m = strings.Replace(y, k, v, -1)
	}
	d := fmt.Sprintf("%d", day)
	for k, v := range convs {
		d = strings.Replace(y, k, v, -1)
	}
	ymd := YearMonthDay{
		Year:  y,
		Month: m,
		Day:   d,
	}
	return ymd
}
