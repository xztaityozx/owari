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
	"github.com/k0kubun/go-ansi"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// boomCmd represents the boom command
var boomCmd = &cobra.Command{
	Use:   "boom",
	Short: "A brief description of your command",
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
		offset, err := cmd.Flags().GetInt("offset")
		if err != nil {
			log.Fatal(err)
		}

		text := strings.Join(args, " ")
		if len(text) == 0 {
			text = ">>1"
		}

		count, err := cmd.Flags().GetInt("countDown")
		if err != nil {
			log.Fatal(err)
		}
		pid, err := cmd.Flags().GetInt("kill")
		if err != nil {
			log.Fatal(err)
		}
		yes, err := cmd.Flags().GetBool("yes")
		if err != nil {
			log.Fatal(err)
		}

		if pid != -1 {
			text = fmt.Sprintf("PID: %d", pid)
		}

		ba := NewBoomAA(text, offset, count)
		ba.Run(pid, yes)
	},
}

type BoomAA struct {
	SeqNum    int
	Width     int
	ThirdLine string
	Offset    int
	Max       int
}

func (ba BoomAA) Run(pid int, yes bool) {
	d, _ := time.ParseDuration("0.1s")

	line := 11

	for i := 0; i < ba.Max; i++ {
		PaddingPrintNoColor(fmt.Sprintf("%d/%d", i+1, ba.Max), ba.Offset)
		for _, v := range ba.NextCount() {
			PaddingPrintNoColor(v, ba.Offset)
		}
		time.Sleep(d)

		if overwrite {
			ansi.CursorUp(line)
		}

		ba.SeqNum++
	}

	if pid != -1 {
		p, err := os.FindProcess(pid)
		if err != nil {
			log.Println(err)
		}
		status := yes
		if !yes {
			fmt.Fprintf(os.Stderr, "PID:%d をKillっていいんですか？ (y/n)>>> ", pid)
			r := bufio.NewScanner(os.Stdin)
			r.Scan()
			status = r.Text() == "y"
			ansi.CursorUp(1)
		}
		if status {
			p.Kill()
		}
	}

	limit, err := strconv.Atoi(count)
	if err != nil || count == "inf" {
		limit = -1
	}

	d, err = time.ParseDuration(duration)
	if err != nil {
		d, _ = time.ParseDuration("0.5s")
	}

	ansi.CursorDown(1)

	for i := 0; i < limit || limit < 0; i++ {
		for _, v := range ba.NextFinish() {
			PaddingPrint(v, ba.Offset)
		}
		ba.SeqNum++

		if overwrite && i != limit-1 {
			ansi.CursorUp(line - 1)
		}
		time.Sleep(d)
	}
}

func NewBoomAA(text string, offset int, max int) BoomAA {

	textLength := GetLooksLength(text)
	if textLength > 15 {
		text = text[0:14]
	}
	upper := (15 - textLength) / 2
	lower := upper

	for upper+textLength+lower >= 15 {
		lower--
	}

	return BoomAA{
		SeqNum:    1,
		Width:     GetWidth(),
		ThirdLine: fmt.Sprintf("     %s|＿| |", strings.Repeat(" ", upper)+text+strings.Repeat(" ", lower)),
		Offset:    offset,
		Max:       max,
	}
}

func (ba BoomAA) NextCountHead() []string {

	cnt := ba.SeqNum % ((ba.Width - 28) / 10)
	max := (ba.Width - 28) / 10

	return []string{
		fmt.Sprintf("                 __|...|     %s%s", strings.Repeat("ピッ・・・", cnt), strings.Repeat("        ", max-cnt)),
		"...                |％| |",
		ba.ThirdLine,
	}
}

func (ba BoomAA) NextCount() []string {
	return append(ba.NextCountHead(), ba.Base()...)
}

func (ba BoomAA) NextFinish() []string {

	cnt := (ba.Width - 29) / 2

	return append([]string{
		fmt.Sprintf("                 __|...|   ピ%s", strings.Repeat("ー", cnt)),
		"...                |  | |",
		ba.ThirdLine,
	}, ba.FinishBase()...)
}

func (ba BoomAA) Base() []string {
	return strings.Split(`     ＿ .／￣＼＿__ノ  |
   /|  | | ^o^ ||ノ    |
   ||, ～～'⌒⌒ヽ～-.､.|
   ||＼ '   ,⌒ ｀    ﾞヽ
   ||＼＼||￣|￣|￣|￣|||
..     ＼||￣￣￣￣￣￣||
         ||￣￣￣￣￣￣||`, "\n")
}

func (ba BoomAA) FinishBase() []string {
	return strings.Split(`     ＿ .／￣＼＿__ノ  |
   /|  | | ^o^ ||ノ    |
   ||, ～～'⌒⌒ヽ～-.､.|
   ||＼ '   ,⌒ ｀    ﾞヽ
   ||＼＼||￣|￣|￣|￣|||
..     ＼||￣￣￣￣￣￣||                ┼ヽ   -|r‐､.  ﾚ  |
         ||￣￣￣￣￣￣||               ｄ⌒)  ./|  _ﾉ   __ﾉ`, "\n")
}

func init() {
	rootCmd.AddCommand(boomCmd)

	boomCmd.Flags().Int("kill", -1, "PIDを受け取って，対象のプロセスをキルします")
	boomCmd.Flags().Int("countDown", 3, "カウントダウンします")
	boomCmd.Flags().Int("offset", 0, "左端からの距離です")

	boomCmd.Flags().BoolP("yes", "y", false, "キルするときの確認を飛ばします")
}
