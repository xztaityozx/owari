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

		PrintKanban(strings.Join(args, " "), offset, gikoneko)
	},
}

func init() {
	rootCmd.AddCommand(kanbanCmd)
	kanbanCmd.Flags().BoolP("giko", "g", false, "ギコ猫を付けます")
	kanbanCmd.Flags().Int("offset", 0, "左端からの距離です")
}

func PrintKanban(text string, offset int, gikoneko bool) {

	if len(text) == 0 {
		text = "終"
	}

	topString := "￣"
	bottomString := "＿"

	length := GetLooksLength(text)

	if length%2 == 1 {
		length++
		text = " " + text
	}

	defaultLength := 6
	upperLength := 6
	lowerLength := 6

	sideLength := 4
	if length < defaultLength {
		w := defaultLength - length
		upperLength += w / 2
		lowerLength += w / 2
	} else {
		sideLength = (length-10)/2 + upperLength
	}

	AA := []string{
		"|" + strings.Repeat(topString, (upperLength+length+lowerLength)/2) + "|",
		"|" + strings.Repeat(" ", upperLength) + text + strings.Repeat(" ", lowerLength) + "|",
		"|" + strings.Repeat(" ", sideLength) + "制作・著作" + strings.Repeat(" ", sideLength) + "|",
		"|  " + strings.Repeat(topString, sideLength+5-2) + "  |",
		"|" + strings.Repeat(" ", sideLength) + " Ｎ Ｈ Ｋ " + strings.Repeat(" ", sideLength) + "|",
		"|" + strings.Repeat(bottomString, (upperLength+length+lowerLength)/2) + "|",
	}

	if gikoneko {
		AA = append(AA, []string{
			strings.Repeat(" ", sideLength-2) + " ∧∧  ||",
			strings.Repeat(" ", sideLength-2) + "( ﾟдﾟ)||",
			strings.Repeat(" ", sideLength-2) + "/　づΦ",
		}...)
	}

	PrintAA(AA, offset)

}
