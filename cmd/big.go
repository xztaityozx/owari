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
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var BigAA = `
          了                了了了
          終終了            了終了
        了終終              終終了
      了終終了            了終終了了了了終終了
      終終終    了了      了終終終終終終終終終了
  了了終終了  了終終了  了終終了了了了了終終終
了終終終終了了終終了    了終終終      了終終了
了終終終終了了終了    了終終了終了  了終終終
  了終終終終終終了  了終終了了終終了終終終了
    了終終終終了了  了終了    了終終終終了
      了終終了了終了  了        了終終終
      了終終    終終了          了終終了了
了了了終終終了終終終終了    了了終終終終終了
終終終終終終終終了終終了了終終終終了終終終終終終了了
終終了了了終終    了了終終終終了了    了終終終終終了
了了了    終終  了了了終終終了了終了了了了終終終終了
  終終終了終終了終終了終終了  了終終終終了了了了了
  終終終了終終了終終了了了    了終終終終終終了
  終終終了終終  了終終          了了終終終終了
了終終終  終終  了終終了    了      了終終了
了終終終  終終    終終了  了終終了了了
了終終了  終終    終終了  了終終終終終了了
了終終了  終終    了了    了終終終終終終終終了了
了終終了  終終              了了終終終終終終終了
          終終                    了終終終終終了
          終終                        了了了了
`

var BigHolly = `
          柊柊木            木木木
          柊柊木            木柊木
          柊柊木            柊柊木
          柊柊木          木柊柊木木木木柊柊木
          柊柊木          木柊柊柊柊柊柊柊柊柊木
  柊柊柊柊柊柊柊柊柊木  木柊柊木木木木木柊柊柊
  柊柊柊柊柊柊柊柊柊木  木柊柊柊      木柊柊木
  木木木木柊柊木木木  木柊柊木柊木  木柊柊柊
        木柊柊柊  木柊柊木木柊柊木柊柊柊木
        柊柊柊柊木  木柊木    木柊柊柊柊木
      木柊柊柊柊柊木  木        木柊柊柊
      木柊柊柊木柊柊木          木柊柊木木
    木柊柊柊柊木柊柊柊木    木木柊柊柊柊柊木
    柊柊柊柊柊木木柊柊木木柊柊柊柊木柊柊柊柊柊柊木木
  柊柊木  柊柊木  木木柊柊柊柊木木    木柊柊柊柊柊木
木柊柊    柊柊木  木木柊柊柊木木柊木木木木柊柊柊柊木
柊柊木    柊柊木  木木柊柊木  木柊柊柊柊木木木木木
  柊      柊柊木    木木木    木柊柊柊柊柊柊木
          柊柊木              木木柊柊柊柊木
          柊柊木          木      木柊柊木
          柊柊木        木柊柊木木木
          柊柊木        木柊柊柊柊柊木木
          柊柊木        木柊柊柊柊柊柊柊柊木木
          柊柊木            木木柊柊柊柊柊柊柊木
          柊柊木                  木柊柊柊柊柊木
          柊柊木                      木木木木
`

// bigCmd represents the big command
var bigCmd = &cobra.Command{
	Use:   "big",
	Short: "デカい終を出しますよ",
	Long:  fmt.Sprintf("%s\n\nを出力します\n引数を与えると，終了の文字がそれになります\n", BigAA),
	Run: func(cmd *cobra.Command, args []string) {
		offset, _ := cmd.Flags().GetInt("offset")

		PrintBig(strings.Join(args, ""), offset)
	},
}

func init() {
	rootCmd.AddCommand(bigCmd)
	bigCmd.Flags().Int("offset", 0, "左端からの距離です")
}

func PrintBig(text string, offset int) {

	status := len(text) == 0
	t := time.Now()
	baseAA := BigAA

	if t.Month() == 12 && t.Day() == 25 {
		baseAA = BigHolly
	}

	if status {
		PrintAA(strings.Split(baseAA, "\n"), offset)
		return
	}

	runeText := []rune(text)
	textIdx := 0
	maxLength := len(runeText)

	var aa []string

	for _, line := range strings.Split(baseAA, "\n") {
		if status {
			PaddingPrint(line, offset)
		} else {
			var replaced []rune
			for _, c := range line {
				if c == rune('終') || c == rune('了') || c == rune('柊') || c == rune('木') {
					replaced = append(replaced, runeText[textIdx])
					textIdx = (textIdx + 1) % maxLength
				} else {
					replaced = append(replaced, rune(' '))
				}
			}
			aa = append(aa, string(replaced))
		}
	}

	PrintAA(aa, offset)
}
