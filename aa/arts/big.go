package arts

import (
	"github.com/xztaityozx/owari/aa"
	"strings"
	"time"
)

// BigOwari は
var bigAA = `
          了                了了了
          終終了            了終了
        了終終              終終了
      了終終了            了終終了了了了終終了
      終終終    了了      了終終終終終終終終終了
  了了終終了  了終終了  了終終了了了了了終終終
了終終終終了了終終了    了終終終      了終終了
了終終終終了了終了    了終終了終了  了終終終
  了終終終終終終了  了終終了了終終了終終終了
    了終終終終了了  了終了    了終終終終了
      了終終了了終了  了        了終終終
      了終終    終終了          了終終了了
了了了終終終了終終終終了    了了終終終終終了
終終終終終終終終了終終了了終終終終了終終終終終終了了
終終了了了終終    了了終終終終了了    了終終終終終了
了了了    終終  了了了終終終了了終了了了了終終終終了
  終終終了終終了終終了終終了  了終終終終了了了了了
  終終終了終終了終終了了了    了終終終終終終了
  終終終了終終  了終終          了了終終終終了
了終終終  終終  了終終了    了      了終終了
了終終終  終終    終終了  了終終了了了
了終終了  終終    終終了  了終終終終終了了
了終終了  終終    了了    了終終終終終終終終了了
了終終了  終終              了了終終終終終終終了
          終終                    了終終終終終了
          終終                        了了了了
`

// です
type BigOwari struct {
	aa.AsciiArt
	text string
}

// NewBigOwari は BigOwari を作って返す
func NewBigOwari(text string) BigOwari {
	return BigOwari{text: text}
}

func (bo *BigOwari) Load(_ string) error {
	baseAA := func() string {
		now := time.Now()
		if now.Month() == 12 && now.Day() == 25 {
			return bigHolly
		} else {
			return bigAA
		}
	}()

	length := len(bo.text)
	if length == 0 {
		bo.AsciiArt = strings.Split(baseAA, "\n")
		return nil
	}

	// textが与えられているなら文字を置き換える
	bo.AsciiArt = []string{}
	textIdx := 0
	for _, line := range strings.Split(baseAA, "\n") {
		var replaced []rune
		for _, r := range line {
			if r == '終' || r == '了' || r == '木' || r == '柊' {
				replaced = append(replaced, rune(bo.text[textIdx]))
				textIdx = (textIdx + 1) % length
			} else {
				replaced = append(replaced, ' ')
			}
		}

		bo.AsciiArt = append(bo.AsciiArt, string(replaced))
	}

	return nil
}

var bigHolly = `
         柊柊木            木木木
         柊柊木            木柊木
         柊柊木            柊柊木
         柊柊木          木柊柊木木木木柊柊木
         柊柊木          木柊柊柊柊柊柊柊柊柊木
 柊柊柊柊柊柊柊柊柊木  木柊柊木木木木木柊柊柊
 柊柊柊柊柊柊柊柊柊木  木柊柊柊      木柊柊木
 木木木木柊柊木木木  木柊柊木柊木  木柊柊柊
       木柊柊柊  木柊柊木木柊柊木柊柊柊木
       柊柊柊柊木  木柊木    木柊柊柊柊木
     木柊柊柊柊柊木  木        木柊柊柊
     木柊柊柊木柊柊木          木柊柊木木
   木柊柊柊柊木柊柊柊木    木木柊柊柊柊柊木
   柊柊柊柊柊木木柊柊木木柊柊柊柊木柊柊柊柊柊柊木木
 柊柊木  柊柊木  木木柊柊柊柊木木    木柊柊柊柊柊木
木柊柊    柊柊木  木木柊柊柊木木柊木木木木柊柊柊柊木
柊柊木    柊柊木  木木柊柊木  木柊柊柊柊木木木木木
 柊      柊柊木    木木木    木柊柊柊柊柊柊木
         柊柊木              木木柊柊柊柊木
         柊柊木          木      木柊柊木
         柊柊木        木柊柊木木木
         柊柊木        木柊柊柊柊柊木木
         柊柊木        木柊柊柊柊柊柊柊柊木木
         柊柊木            木木柊柊柊柊柊柊柊木
         柊柊木                  木柊柊柊柊柊木
         柊柊木                      木木木木
`
