package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xztaityozx/owari/aa/arts"
	"log"
)

var deadCmd = &cobra.Command{
	Use: "dead",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		dead := arts.NewDead(name)
		_ = dead.Load("")

		if err := writer.Write(dead.AsciiArt); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	deadCmd.Flags().String("name", "owariコマンド", "終わってしまったものの名前です")
	rootCmd.AddCommand(deadCmd)
}
