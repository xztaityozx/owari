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
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/xztaityozx/owari/aa/arts"
)

// kanbanCmd represents the kanban command
var kanbanCmd = &cobra.Command{
	Use:     "kanban",
	Aliases: []string{"kan"},
	Short:   "看板を持ちます",
	Long: `

|￣￣￣￣￣￣￣￣￣|
|        終        |
|    制作・著作    |
|  ￣￣￣￣￣￣￣  |
|     Ｏ Ｗ Ｒ     |
|＿＿＿＿＿＿＿＿＿|
  ∧∧  ||
 ( ﾟдﾟ)||
 /    づΦ


看板がでます
引数を与えると「終」の部分が変わります
`,
	Run: func(cmd *cobra.Command, args []string) {
		gikoneko, _ := cmd.Flags().GetBool("giko")
		useStdin, _ := cmd.Flags().GetBool("stdin")
		konata, _ := cmd.Flags().GetBool("konata")
		textimg, _ := cmd.Flags().GetBool("textimg")
		author, _ := cmd.Flags().GetString("author")
		reverse, _ := cmd.Flags().GetBool("reverse")
		twin, _ := cmd.Flags().GetBool("twin")
		font, _ := cmd.Flags().GetString("font")

		text := func() []string {
			if useStdin {
				var rt []string
				buf := bufio.NewScanner(os.Stdin)
				for buf.Scan() {
					rt = append(rt, buf.Text())
				}
				return rt
			} else if len(args) != 0 {
				return args
			} else {
				return nil
			}
		}()

		if textimg {
			// textimgフラグが建っていたら、fontをNotSansCJKにしちゃう
			font = "NotoSansCJK"
		}

		kanban := arts.NewKanban(text)
		kanban.SetKonata(konata)
		kanban.SetReverse(reverse)
		kanban.SetGikoNeko(gikoneko)
		kanban.SetTwinGikoNeko(twin)
		kanban.SetAuthor(author)

		if err := kanban.Load(font); err != nil {
			log.Fatalln(err)
		}

		if err := writer.Write(kanban.AsciiArt); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(kanbanCmd)
	kanbanCmd.Flags().BoolP("giko", "g", false, "ギコ猫を付けます")
	kanbanCmd.Flags().BoolP("stdin", "i", false, "標準入力を受取ります")
	kanbanCmd.Flags().Bool("konata", false, "こなた")
	kanbanCmd.Flags().BoolP("textimg", "t", false, "textimg用に出力を整えます。実際には--font=NotoSansCJKのShortHandです")
	kanbanCmd.Flags().StringP("author", "a", "Ｏ Ｗ Ｒ", "制作・著作者を指定します")
	kanbanCmd.Flags().StringP("font", "f", "default", "指定したフォントでの描画用に出力を整えます")
	kanbanCmd.Flags().Bool("reverse", false, "ギコ猫を反転します")
	kanbanCmd.Flags().Bool("twin", false, "ギコ猫を二匹に増やします")
}
