package main

import "strings"

const taishokuText = `
このたび一身上の都合により、{taishokuDate}をもって
退職いたしたくここにお願い申し上げます。
{today}

　　　　　　　　　　　　　　{department}　{team}
　　　　　　　　　　　　　　{yourName}　印

{company}
{president}　{presidentName}
`

// https://tenshoku.mynavi.jp/knowhow/caripedia/58

func makeTaishokuText(taishokuDate, today, department, team, yourName, company, president, presidentName string) string {
	text := taishokuText
	convs := map[string]string{
		"{taishokuDate}":  taishokuDate,
		"{today}":         today,
		"{department}":    department,
		"{team}":          team,
		"{yourName}":      yourName,
		"{company}":       company,
		"{president}":     president,
		"{presidentName}": presidentName,
	}
	for k, v := range convs {
		text = strings.Replace(text, k, v, 1)
	}
	return text[1 : len(text)-1]
}

func convertStringNumberToKanji(s string) string {
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

	for k, v := range convs {
		s = strings.Replace(s, k, v, -1)
	}
	return s
}
