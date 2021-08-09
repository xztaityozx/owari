package cmd

import (
	"bufio"
	"github.com/xztaityozx/owari/aa/arts"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var graveCmd = &cobra.Command{
	Use:   "grave",
	Short: "先祖代々のお墓を出力します",
	Long: `
    ┌─┐
    │先│
    │祖│
    │代│
    │々│
    │之│
    │ば│
    │か│
  ┌┴─┴┐
 │| 三三 |│
￣￣￣￣￣￣￣
先祖代々のお墓です。引数を与えると文字を入れ替えることができます。ただし必ず一列になります

`,

	Run: func(cmd *cobra.Command, args []string) {
		stdin, _ := cmd.Flags().GetBool("stdin")
		//t, _ := cmd.Flags().GetString("type")
		text := func() []string {
			if stdin {
				var lines []string
				sc := bufio.NewScanner(os.Stdin)
				for sc.Scan() {
					lines = append(lines, sc.Text())
				}
				if err := sc.Err(); err != nil {
					panic(err)
				}

				return lines
			} else if len(args) != 0 {
				return args
			} else {
				return []string{"先祖代々之ばか"}
			}
		}()

		grave := arts.NewGrave(strings.Join(text, ""))
		if err := grave.Load(""); err != nil {
			log.Fatal(err)
		}

		if err := writer.Write(grave.AsciiArt); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(graveCmd)
	graveCmd.Flags().BoolP("stdin", "i", false, "標準入力を受取ります")
	//graveCmd.Flags().String("type", "default", "フォントや形状などを選択することができます。存在しないものの場合はdefaultが代わりに使われます")
}
