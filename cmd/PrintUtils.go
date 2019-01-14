package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/k0kubun/go-ansi"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var FrontColors = []color.Attribute{
	color.FgBlack,
	color.FgBlue,
	color.FgCyan,
	color.FgGreen,
	color.FgHiBlack,
	color.FgHiBlue,
	color.FgHiCyan,
	color.FgHiGreen,
	color.FgHiMagenta,
	color.FgHiRed,
	color.FgHiWhite,
	color.FgHiYellow,
}

func PaddingPrint(text string, c int) {
	for i := 0; i < c; i++ {
		fmt.Printf(" ")
	}
	if colorful {
		// カラフルなやつ
		for _, v := range text {
			idx := rand.Intn(len(FrontColors))
			ansi.Print(FrontColors[idx])
			ansi.Print(v)
			ansi.Print(color.Reset)
		}
		fmt.Println()
	} else {
		fmt.Println(text)
	}
}

// 端末の幅を取る
func GetWidth() int {
	if reqWidth == "auto" {
		// sttyがあれば
		if _, err := exec.LookPath("stty"); err == nil {
			stty := exec.Command("stty", "size")
			stty.Stdin = os.Stdin
			out, err := stty.Output()
			if err != nil {
				log.Fatal(err)
			}

			hw := strings.Split(string(out), " ")
			rt, err := strconv.Atoi(hw[1])
			if err != nil {
				log.Fatal(err)
			}

			return rt
		}

		// windowsなら
		if runtime.GOOS == "windows" {
			out, err := exec.Command("powershell.exe", "/C", "Write-Output", "$Host.UI.RawUI.WindowSize.Width").Output()
			if err != nil {
				log.Fatal(err)
			}

			trimed := strings.Trim(string(out), "\r\n")
			rt, err := strconv.Atoi(trimed)
			if err != nil {
				log.Fatal(err)
			}

			return rt

		}

		// そうじゃなかったら80を返す
		return 80
	}

	rt, err := strconv.Atoi(reqWidth)
	if err != nil {
		log.Fatal(err)
	}

	return rt
}
