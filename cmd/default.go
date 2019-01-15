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
	"math"
	"strings"

	"github.com/spf13/cobra"
)

// defaultCmd represents the default command
var defaultCmd = &cobra.Command{
	Use:     "default",
	Short:   "基本の終わりを出力するよ",
	Aliases: []string{"def"},
	Long: `
       糸冬
-------------------
 制作・著作 ＮＨＫ

を出力します
引数を与えると「糸冬」の部分に置き換わります`,
	Run: func(cmd *cobra.Command, args []string) {
		offset, _ := cmd.Flags().GetInt("offset")
		PrintDefault(strings.Join(args, " "), offset)
	},
}

func init() {
	rootCmd.AddCommand(defaultCmd)

	defaultCmd.Flags().Int("offset", 0, "右からの距離です")
}
func PrintDefault(text string, offset int) {
	if len(text) == 0 {
		text = "糸冬"
	}

	// 幅を合わせていく
	// 上のテキストの見た目の幅を数える
	length := GetLooksLength(text)

	// ハイフンの数
	upper := 8
	lower := 8
	defaultBarSize := 20

	if length+upper+lower < defaultBarSize {
		div := defaultBarSize - length - upper - lower
		upper += div / 2
		lower += div - (div / 2)
	}

	AA := []string{
		"",
		"",
		fmt.Sprintf("%s%s", strings.Repeat(" ", upper), text),
		strings.Repeat("-", length+upper+lower),
		strings.Repeat(" ", int(math.Max(0, 2+float64(length+upper+lower-defaultBarSize)/2))) + "制作・著作 ＮＨＫ",
	}

	c := GetWidth() - length - upper - lower - offset

	PrintAA(AA, c)
}
