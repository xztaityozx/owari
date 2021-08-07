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
	Version: "2.0",
	Long: `
author: xztaityozx
repository: https://github.com/xztaityozx/owari


       糸冬
-------------------
 制作・著作 ＮＨＫ

を出力するよ。ほかにもいろいろあるよ。


`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 全体で使うaa.Writerを作る
		writer = aa.NewWriter(os.Stdout)
		cf, err := cmd.PersistentFlags().GetBool("colorful")
		if err != nil {
			return err
		}
		cfa, err := cmd.PersistentFlags().GetBool("colorful-always")
		if err != nil {
			return err
		}
		overwrite, err := cmd.PersistentFlags().GetBool("overwrite")
		if err != nil {
			return err
		}
		d, err := cmd.PersistentFlags().GetDuration("duration")
		if err != nil {
			return err
		}
		c, err := func() (int, error) {
			count, err := cmd.PersistentFlags().GetString("count")
			if count == "inf" {
				return -1, err
			}
			return strconv.Atoi(count)
		}()
		if err != nil {
			return err
		}
		offset, err := cmd.PersistentFlags().GetInt("offset")
		if err != nil {
			return err
		}
		ie, err := cmd.PersistentFlags().GetBool("insert-empty")
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
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return os.Stdout.Close()
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
	rootCmd.PersistentFlags().BoolP("colorful-always", "C", false, "colorfulフラグが有効なとき、パイプやリダイレクト時にもCOLOR＿CODEが適用されるよう強制します")
	rootCmd.PersistentFlags().StringP("width", "w", "auto", "表示幅です。autoにすると端末の幅を取得します")
	_ = rootCmd.PersistentFlags().MarkDeprecated("width", "AAの最大幅を指定することはできなくなりました")

	rootCmd.PersistentFlags().Bool("overwrite", false, "複数回出力するときに同じ場所に上書きし続けます")
	rootCmd.PersistentFlags().StringP("count", "n", "1", "指定回数繰り返します。負数かinfを指定すると無限になります")

	defaultDuration, _ := time.ParseDuration("0.5s")
	rootCmd.PersistentFlags().Duration("duration", defaultDuration, "繰り返しのインターバルです")

	rootCmd.PersistentFlags().Int("offset", 0, "左からの距離です")

	rootCmd.PersistentFlags().BoolP("insert-empty", "E", true, "出力の1行目に必ず空白行を挿入します")
}
