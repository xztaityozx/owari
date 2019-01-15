package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/k0kubun/go-ansi"
	"golang.org/x/text/width"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
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

func GetLooksLength(text string) int {
	length := 0
	for _, v := range []rune(text) {
		kind := width.LookupRune(v).Kind()

		if kind == width.EastAsianWide {
			length += 2
		} else {
			length += 1
		}
	}
	return length
}

func PrintAA(aa []string, padding int) {

	length := len(aa)

	d, err := time.ParseDuration(duration)
	if err != nil {
		d, _ = time.ParseDuration("0.5s")
	}

	limit, err := strconv.Atoi(count)
	if err != nil || count == "inf" {
		limit = -1
	}
	// ひたすら
	for i := 0; i < limit || limit < 0; i++ {
		for _, v := range aa {
			PaddingPrint(v, padding)
		}

		if overwrite && i != limit-1 {
			ansi.CursorUp(length)
		}

		time.Sleep(d)

	}
}

func PaddingPrint(text string, c int) {
	for i := 0; i < c; i++ {
		fmt.Printf(" ")
	}

	rand.Seed(time.Now().UnixNano())

	if colorful {
		// カラフルなやつ
		for _, v := range []rune(text) {
			idx := rand.Intn(len(FrontColors))
			p := color.New(FrontColors[idx])
			p.Printf("%c", v)
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
			rt, err := strconv.Atoi(strings.Trim(hw[1], "\n"))
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
