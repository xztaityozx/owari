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
	"github.com/spf13/cobra"
	"github.com/xztaityozx/owari/aa/arts"
	"log"
)

// boomCmd represents the boom command
var boomCmd = &cobra.Command{
	Use:        "boom",
	Short:      "ブーム君がおわるよ",
	Deprecated: "owariはブーム君の提供を終了しました。ご利用ありがとうございました。",
	Long: `

                 __|...|     ピッ・・・ピッ・・・ピッ・・・
...                 |％| |
          >>1      |＿| |
    ＿ .／￣＼＿__ノ  |
  /|  | | ^o^ ||ノ    |
 ||, ～～'⌒⌒ヽ～-.､.|
 ||＼ '   ,⌒ ｀    ﾞヽ
 ||＼＼||￣|￣|￣|￣|||
..    ＼||￣￣￣￣￣￣||
       ||￣￣￣￣￣￣||

                __|...|     ピーーーーーーーーーー
..                 |  | |
         >>1      |＿| |
    ＿ ／￣＼＿__ノ   |
  /|  || ^o^ ||ノ     |
 ||, ～～'⌒⌒ヽ～-.､.|
 ||＼ '   ,⌒ ｀    ﾞヽ
 ||＼＼||￣|￣|￣|￣|||
..    ＼||￣￣￣￣￣￣||                ┼ヽ   -|r‐､.  ﾚ  |
       ||￣￣￣￣￣￣||                ｄ⌒) ./|  _ﾉ    __ﾉ

を表示します．引数を与えると 「>>1」の部分を置き換えますが，半角幅15文字までです．
`,
	Run: func(cmd *cobra.Command, args []string) {
		grave := arts.NewGrave("boomサブコマンドの墓")
		if err := grave.Load(""); err != nil {
			log.Fatalln(err)
		}

		if err := writer.Write(grave.AsciiArt); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(boomCmd)
}
