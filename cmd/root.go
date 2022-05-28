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
	"github.com/spf13/cobra"
	"github.com/xztaityozx/owari/aa"
	"os"
	"strconv"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "owari",
	Short:   "終了を知らせるAAを出力するコマンドだよ！仲良く使ってね！",
	Version: "2.0.2",
	Long: `

       糸冬
-------------------
 制作・著作 ＯＷＲ

を出力するよ。ほかにもいろいろあるよ。


author: xztaityozx
repository: https://github.com/xztaityozx/owari

`,
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return os.Stdout.Close()
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 全体で使うaa.Writerを作る
		writer = aa.NewWriter(os.Stdout)
		cf, err := cmd.Flags().GetBool("colorful")
		if err != nil {
			return err
		}

		cfa, err := cmd.Flags().GetBool("colorful-always")
		if err != nil {
			return err
		}

		overwrite, err := cmd.Flags().GetBool("overwrite")
		if err != nil {
			return err
		}

		d, err := cmd.Flags().GetDuration("duration")
		if err != nil {
			return err
		}

		c, err := func() (int, error) {
			count, err := cmd.Flags().GetString("count")
			if err != nil {
				return 0, err
			}
			if count == "inf" {
				return -1, nil
			}
			return strconv.Atoi(count)
		}()
		if err != nil {
			return err
		}

		offset, err := cmd.Flags().GetInt("offset")
		if err != nil {
			return err
		}

		ie, err := cmd.Flags().GetBool("insert-empty")
		if err != nil {
			return err
		}

		writer.SetTimes(c)
		writer.SetColorfulAlways(cfa)
		writer.SetColorful(cf || cfa)
		writer.SetDuration(d)
		writer.SetOverwrite(overwrite)
		writer.SetOffset(offset)
		writer.SetInsertEmpty(ie)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		defaultCmd.Run(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var writer aa.Writer

func init() {
	rootCmd.PersistentFlags().BoolP("colorful", "c", false, "カラフルにします")
	rootCmd.PersistentFlags().BoolP("colorful-always", "C", false, "colorfulフラグが有効なとき、パイプやリダイレクト時にもCOLOR_CODEが適用されるよう強制します")
	// Deprecated {{{
	rootCmd.PersistentFlags().StringP("width", "w", "auto", "表示幅です。autoにすると端末の幅を取得します")
	_ = rootCmd.PersistentFlags().MarkDeprecated("width", "AAの最大幅を指定することはできなくなりました")
	// }}}
	rootCmd.PersistentFlags().BoolP("overwrite", "o", false, "複数回出力するときに同じ場所に上書きし続けます")
	rootCmd.PersistentFlags().StringP("count", "n", "1", "指定回数出力します。負数かinfを指定すると無限になります")
	defaultDuration, _ := time.ParseDuration("0.5s")
	rootCmd.PersistentFlags().DurationP("duration", "d", defaultDuration, "複数回出力のインターバルです")
	rootCmd.PersistentFlags().Int("offset", 0, "左からアスキーアートまでの距離です")
	rootCmd.PersistentFlags().BoolP("insert-empty", "E", true, "出力の1行目に必ず空白行を挿入します")
}
