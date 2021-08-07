package cmd

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//
//	"github.com/spf13/cobra"
//)
//
//var graveCmd = &cobra.Command{
//	Use:   "grave",
//	Short: "先祖代々のお墓を出力します",
//	Long: `
//     ┌─┐
//     │先│
//     │祖│
//     │代│
//     │々│
//     │之│
//     │ば│
//     │か│
//   ┌┴─┴┐
//  │| 三三 |│
//￣￣￣￣￣￣￣
//先祖代々のお墓です。引数を与えると文字を入れ替えることができます。ただし必ず一列になります
//
//`,
//
//	Run: func(cmd *cobra.Command, args []string) {
//		stdin, _ := cmd.Flags().GetBool("stdin")
//		text := func() []string {
//			if stdin {
//				lines := []string{}
//				sc := bufio.NewScanner(os.Stdin)
//				for sc.Scan() {
//					lines = append(lines, sc.Text())
//				}
//				if err := sc.Err(); err != nil {
//					panic(err)
//				}
//
//				return lines
//			} else if len(args) != 0 {
//				return args
//			} else {
//				return []string{"先祖代々之ばか"}
//			}
//		}()
//
//		aa := []string{
//			"     ┌─┐",
//		}
//		for _, line := range text {
//			for _, c := range line {
//				if GetLooksLength(fmt.Sprintf("%c", c)) == 1 {
//					aa = append(aa, fmt.Sprintf("     │%c │", c))
//				} else {
//					aa = append(aa, fmt.Sprintf("     │%c│", c))
//				}
//			}
//		}
//
//		aa = append(aa, "   ┌┴─┴┐")
//		aa = append(aa, "  │| 三三 |│")
//		aa = append(aa, "￣￣￣￣￣￣￣")
//
//		PrintAA(aa, 0)
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(graveCmd)
//	graveCmd.Flags().BoolP("stdin", "i", false, "標準入力を受取ります")
//}
