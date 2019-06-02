// Copyright © 2019 xztaityozx
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// kanbanCmd represents the kanban command
var kanbanCmd = &cobra.Command{
	Use:     "kanban",
	Aliases: []string{"kan"},
	Short:   "看板を持ちます",
	Long: `
|￣￣￣￣￣￣￣￣￣|
|　　　　終　　　　|
|　  制作・著作  　|
|  ￣￣￣￣￣￣￣  |
|     Ｎ Ｈ Ｋ     |
|＿＿＿＿＿＿＿＿＿|
   ∧∧  ||
  ( ﾟдﾟ)||
  /　づΦ

看板がでます
引数を与えると「終」の部分が変わります
`,
	Run: func(cmd *cobra.Command, args []string) {
		offset, _ := cmd.Flags().GetInt("offset")
		gikoneko, _ := cmd.Flags().GetBool("giko")
		useStdin, _ := cmd.Flags().GetBool("stdin")
		konata, _ := cmd.Flags().GetBool("konata")
		author, _ := cmd.Flags().GetString("author")

		gikoneko = gikoneko || konata

		if useStdin {
			lines := readStdin()
			printKanban(lines, offset, gikoneko, konata, author)
			return
		}

		PrintKanban(strings.Join(args, " "), offset, gikoneko, konata, author)
	},
}

func init() {
	rootCmd.AddCommand(kanbanCmd)
	kanbanCmd.Flags().BoolP("giko", "g", false, "ギコ猫を付けます")
	kanbanCmd.Flags().Int("offset", 0, "左端からの距離です")
	kanbanCmd.Flags().BoolP("stdin", "i", false, "標準入力を受取ります")
	kanbanCmd.Flags().Bool("konata", false, "こなた")
	kanbanCmd.Flags().String("author", " Ｎ Ｈ Ｋ ", "制作/著作者を指定します")
}

func PrintKanban(text string, offset int, gikoneko, konata bool, author string) {
	if len(text) == 0 {
		text = "終"
	}
	texts := []string{text}
	printKanban(texts, offset, gikoneko, konata, author)
}

func printKanban(texts []string, offset int, gikoneko, konata bool, author string) {
	const topString = "￣"
	const bottomString = "＿"

	// テキストを改行文字で分割し、最長文字列に合わせて空白詰め
	// 最長文字列の長さも返す
	texts, textLen, author := makePaddedText(texts, author)

	sideLength := 4 // テキストの横に確保しておきたい空白
	maxLength := 18 // 看板の横幅
	// デフォルト値よりも大きい問だけ幅を拡張する
	if maxLength < textLen+sideLength {
		maxLength = textLen + sideLength
	}

	var AA []string
	AA = append(AA, fmt.Sprintf("|%s|", strings.Repeat(topString, maxLength/2)))
	for _, t := range texts {
		// 看板の横幅に合わせて空白埋め
		t = padSpace(t, maxLength)
		s := fmt.Sprintf("|%s|", t)
		AA = append(AA, s)
	}
	AA = append(AA, fmt.Sprintf("|%s|", padSpace("制作・著作", maxLength)))
	AA = append(AA, fmt.Sprintf("|  %s  |", strings.Repeat(topString, (maxLength/2)-2)))
	AA = append(AA, fmt.Sprintf("|%s|", padSpace(author, maxLength)))
	AA = append(AA, fmt.Sprintf("|%s|", strings.Repeat(bottomString, maxLength/2)))

	if gikoneko {

		if konata {
			AA = append(AA, fmt.Sprintf("%s ", padSpace("∧∧  ||", maxLength-3)))
			AA = append(AA, fmt.Sprintf("%s ", padSpace("(≡ω≡.)||", maxLength-4)))
			AA = append(AA, fmt.Sprintf("%s ", padSpace("/     づΦ", maxLength-3)))
		} else {
			AA = append(AA, fmt.Sprintf("%s ", padSpace("∧∧ ||", maxLength-2)))
			AA = append(AA, fmt.Sprintf("%s ", padSpace("( ﾟдﾟ)||", maxLength-2)))
			AA = append(AA, fmt.Sprintf("%s ", padSpace("/   づΦ", maxLength-2)))
		}

	}

	PrintAA(AA, offset)
}

func makePaddedText(ts []string, author string) ([]string, int, string) {
	if len(ts) < 1 {
		return []string{}, 0, author
	}

	maxLength := GetLooksLength(author)

	var texts = ts[:]
	for i := 0; i < len(texts); i++ {
		t := texts[i]
		l := GetLooksLength(t)

		// もっとも長さの長い文字列の長さを取得
		if maxLength < l {
			maxLength = l
		}

		// 文字列の長さを偶数に統一するための空白埋め
		if l%2 == 1 {
			texts[i] = " " + t
		}
	}

	// 文字列の長さを空白で埋めて一番長い文字列に合わせる
	for i := 0; i < len(texts); i++ {
		t := texts[i]
		l := GetLooksLength(t)

		if l < maxLength {
			texts[i] = padSpace(t, maxLength)
		}
	}

	// 最長文字列自体は偶数に合わせているので
	// もし奇数の最長値がセットされてたら修正
	if maxLength%2 == 1 {
		maxLength++
	}

	// author も偶数文字長へ調整
	if GetLooksLength(author)%2 == 1 {
		author = " " + author
	}

	return texts, maxLength, author
}

// padSpace は前後を半角スペースで埋めた文字列を返す。
// 文字列の長さは見た目上の長さで偶数でなければならない。
func padSpace(t string, max int) string {
	l := GetLooksLength(t)
	diff := max - l
	padLen := diff / 2
	var pad string
	if 0 < padLen {
		pad = strings.Repeat(" ", padLen)
	}
	return pad + t + pad
}

func readStdin() (ret []string) {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		ret = append(ret, line)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return
}
