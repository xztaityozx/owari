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
	"github.com/xztaityozx/owari/aa/arts"
	"log"
	"strings"
)

// bigCmd represents the big command
var bigCmd = &cobra.Command{
	Use:   "big",
	Short: "デカい終を出しますよ",
	Long:  fmt.Sprintf("デカい終を出力します\n引数を与えると，終了の文字がそれになります\n"),
	Run: func(cmd *cobra.Command, args []string) {
		bo := arts.NewBigOwari(strings.Join(args, ""))
		if err := bo.Load(""); err != nil {
			log.Fatal(err)
		}

		if err := writer.Write(bo.AsciiArt); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(bigCmd)
}
