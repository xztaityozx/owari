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
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "owari",
	Short: "終了を知らせるAAを出力するコマンドだよ！仲良く使ってね！",
	Long: `
       糸冬
-------------------
 制作・著作 ＮＨＫ

を出力します
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		// デフォルトを呼ぶ
		offset, _ := cmd.Flags().GetInt("offset")
		PrintDefault("", offset)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var colorful bool
var reqWidth string

func init() {
	rootCmd.PersistentFlags().BoolVar(&colorful, "colorful", false, "カラフルにします")
	rootCmd.PersistentFlags().StringVarP(&reqWidth, "reqWidth", "w", "auto", "表示幅です．autoにすると端末幅を取得します")

	rootCmd.Flags().Int("offset", 0, "左からの距離です")
}
