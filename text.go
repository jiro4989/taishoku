package main

import (
	"strings"
	"unicode/utf8"
)

const taishokuText = `
　　　退　職　願
　　　　　　　　　　　　　　　　　　　　　　　　　私議

このたび一身上の都合により、{taishokuDate}をもって
退職いたしたくここにお願い申し上げます。
{today}

　　　　　　　　　　　　　　{department}　{team}
　　　　　　　　　　　　　　{yourName}　印

{company}
{president}　{presidentName}
`

// makeTaishokuText は退職願テキストを生成する。
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

// convertStringNumberToKanji は半角英数字を全角漢数字に変換する。
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

// padSpace は文字列のスライスのうち、一番長い文字列にあわせてテキストの後ろに全角空白を埋める。
// 返却する文字列は新しい配列なので、引数に渡した配列に破壊的変更は与えない。
// また、空白はマルチバイト全角空白である。
// ここでの空白埋めは見た目上のテキストの長さを指す。
func padSpace(s []string) []string {
	text := make([]string, len(s))
	copy(text, s)

	var maxLength int
	// 一番長い文字列の長さを取得
	for _, t := range text {
		l := utf8.RuneCountInString(t)
		if maxLength < l {
			maxLength = l
		}
	}

	// 一番長い文字列の長さに合わせて空白を追加
	for i := 0; i < len(text); i++ {
		diff := maxLength - utf8.RuneCountInString(text[i])
		if 0 < diff {
			pad := strings.Repeat("　", diff)
			text[i] += pad
		}
	}

	return text
}

// reverse は配列の並びを反転させる。
// 文字順で逆順ソートとかしたりはしない。
func reverse(s []string) []string {
	s2 := make([]string, len(s))
	var j int
	for i := len(s) - 1; 0 <= i; i-- {
		s2[j] = s[i]
		j++
	}
	return s2
}

// toVertical は横書き文字列配列を縦書き配列に変換する。
// 前提として、引数の配列はrune文字列としてすべて長さが同じでなければならない。
func toVertical(s []string) (ret []string) {
	if len(s) < 1 {
		return []string{}
	}

	l := utf8.RuneCountInString(s[0])
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
		ret = append(ret, s)
	}
	return
}
