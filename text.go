package main

import (
	"strings"
	"unicode/utf8"
)

const taishokuNegai = `
　　　退　職　願
　　　　　　　　　　　　　　　　　　　　　　　　　私儀

このたび一身上の都合により、{taishokuDate}をもって
退職いたしたくここにお願い申し上げます。
{today}

　　　　　　　　　　　　　　{department}　{team}
　　　　　　　　　　　　　　{yourName}　印

{company}
{president}　{presidentName}殿
`

const taishokuTodoke = `
　　　退　職　届
　　　　　　　　　　　　　　　　　　　　　　　　　私儀

このたび一身上の都合により、{taishokuDate}をもって
退職いたします。
{today}

　　　　　　　　　　　　　　{department}　{team}
　　　　　　　　　　　　　　{yourName}　印

{company}
{president}　{presidentName}殿
`

// makeTaishokuText は退職テキストを生成する。
func makeTaishokuText(templateText, taishokuDate, today, department, team, yourName, company, president, presidentName string) string {
	text := templateText
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

func toZenkaku(s string) string {
	convs := map[string]string{
		"a": "Ａ",
		"b": "Ｂ",
		"c": "Ｃ",
		"d": "Ｄ",
		"e": "Ｅ",
		"f": "Ｆ",
		"g": "Ｇ",
		"h": "Ｈ",
		"i": "Ｉ",
		"j": "Ｊ",
		"k": "Ｋ",
		"l": "Ｌ",
		"m": "Ｍ",
		"n": "Ｎ",
		"o": "Ｏ",
		"p": "Ｐ",
		"q": "Ｑ",
		"r": "Ｒ",
		"s": "Ｓ",
		"t": "Ｔ",
		"u": "Ｕ",
		"v": "Ｖ",
		"w": "Ｗ",
		"x": "Ｘ",
		"y": "Ｙ",
		"z": "Ｚ",
		"A": "ａ",
		"B": "ｂ",
		"C": "ｃ",
		"D": "ｄ",
		"E": "ｅ",
		"F": "ｆ",
		"G": "ｇ",
		"H": "ｈ",
		"I": "ｉ",
		"J": "ｊ",
		"K": "ｋ",
		"L": "ｌ",
		"M": "ｍ",
		"N": "ｎ",
		"O": "ｏ",
		"P": "ｐ",
		"Q": "ｑ",
		"R": "ｒ",
		"S": "ｓ",
		"T": "ｔ",
		"U": "ｕ",
		"V": "ｖ",
		"W": "ｗ",
		"X": "ｘ",
		"Y": "ｙ",
		"Z": "ｚ",
		"0": "０",
		"1": "１",
		"2": "２",
		"3": "３",
		"4": "４",
		"5": "５",
		"6": "６",
		"7": "７",
		"8": "８",
		"9": "９",
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
	for i, j := len(s)-1, 0; 0 <= i; i, j = i-1, j+1 {
		s2[j] = s[i]
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
			for colIdx, c := range []rune(row) {
				if colIdx == i {
					line = append(line, string(c))
				}
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
