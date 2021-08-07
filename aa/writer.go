package aa

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/k0kubun/go-ansi"
	"io"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

type Writer struct {
	// 出力先
	w io.Writer
	// 色付きで出力する
	colorful bool
	// パイプ・リダイレクト先にも強制する
	colorfulAlways bool
	// 左からの距離
	offset int
	// 複数回出力するとき、上書きするかどうか
	overwrite bool
	// 出力する回数
	times int
	// 出力のインターバル
	duration time.Duration
	// 改行文字
	newline string
	// アスキーアートの1行目に必ず空行を出力するかどうか
	beginEmpty bool
}

// NewWriter はAAをio.Writerに出力するやつ
func NewWriter(w io.Writer) Writer {
	return Writer{
		w:              w,
		colorful:       false,
		colorfulAlways: false,
		offset:         0,
		overwrite:      false,
		times:          1,
		duration:       0,
		newline: func() string {
			if runtime.GOOS == "windows" {
				return "\r\n"
			} else if runtime.GOOS == "darwin" {
				return "\r"
			} else {
				return "\n"
			}
		}(),
		beginEmpty: false,
	}
}

// SetInsertEmpty は出力時に1行目に必ず空行を出力するかどうかのON/OFFを設定する
func (w *Writer) SetInsertEmpty(b bool) {
	w.beginEmpty = b
}

// SetColorful はランダムに色付けするようにするかどうかのON/OFFを設定する
func (w *Writer) SetColorful(b bool) {
	w.colorful = b
}

// SetColorfulAlways は SetColorful に加えて
// パイプやリダイレクト時にも色付けを強制するかどうかのON/OFFを設定する
func (w *Writer) SetColorfulAlways(b bool) {
	w.colorful = b
	w.colorfulAlways = b
}

// SetOffset は左端からのAA出力位置の値をセットする
func (w *Writer) SetOffset(offset int) {
	w.offset = offset
}

// SetOverwrite は SetTimes で2回以上出力する場合に
// 同じ場所へ上書きするかどうかのON/OFFを設定する
func (w *Writer) SetOverwrite(b bool) {
	w.overwrite = b
}

// SetTimes はAAの出力回数を設定する
func (w *Writer) SetTimes(times int) {
	w.times = times
}

// SetDuration は SetTimes で2回以上出力するように指定したときのインターバルを設定する
func (w *Writer) SetDuration(d time.Duration) {
	w.duration = d
}

// Write はアスキーアートを出力する
func (w *Writer) Write(aa []string) error {
	// 1行目に空行を出力する
	// ほぼShellGei-Bot向けのコード
	if w.beginEmpty {
		if _, err := fmt.Fprintln(w.w, ""); err != nil {
			return err
		}
	}

	// ランダム色出力のためのSeed初期化
	if w.colorful {
		rand.Seed(time.Now().UnixNano())
	}

	padding := strings.Repeat(" ", w.offset)
	aaHeight := len(aa)

	for i := 0; i < w.times || w.times < 0; i++ {
		for _, line := range aa {
			if err := w.writeLine(line, padding); err != nil {
				return err
			}
		}

		// 上書きするときはカーソルを上に戻す。ただし最後の出力ではやらない
		if w.overwrite && i != w.times-1 {
			ansi.CursorUp(aaHeight)
		}

		// 複数回出力する場合は待機する
		if w.times != 1 {
			time.Sleep(w.duration)
		}
	}

	return nil
}

// writeLine は1行出力する
func (w *Writer) writeLine(line, padding string) error {
	// 左からのパディングを出力
	if _, err := fmt.Fprint(w.w, padding); err != nil {
		return err
	}

	if w.colorful {
		// 色付き出力
		for _, c := range line {
			if _, err := w.getRandomColor().Fprintf(w.w, "%c", c); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintln(w.w, ""); err != nil {
			return err
		}
	} else {
		// 色なし出力
		if _, err := fmt.Fprintln(w.w, line); err != nil {
			return err
		}
	}

	return nil
}

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

// getRandomColor はランダムに色を選んで color.Color を返す
func (w *Writer) getRandomColor() *color.Color {
	if w.colorfulAlways {
		color.NoColor = false
	}
	return color.New(FrontColors[rand.Intn(len(FrontColors))])
}
