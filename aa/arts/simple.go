package arts

import (
	"fmt"
	"github.com/xztaityozx/owari/aa"
	"strings"
)

// SimpleOwari は
//
//
//         糸冬
// ---------------------
//   制作・著作 ＮＨＫ
//
// です

type SimpleOwari struct {
	aa.AsciiArt
	author string
	text   string
}

// NewSimpleOwari は SimpleOwari を作って返す
func NewSimpleOwari(text, author string) SimpleOwari {
	return SimpleOwari{
		author: author, text: text,
	}
}

func (s *SimpleOwari) Load(_ string) error {
	if len(s.text) == 0 {
		s.text = "糸冬"
	}

	if len(s.author) == 0 {
		s.author = "ＮＨＫ"
	}
	signature := fmt.Sprintf("制作・著作 %s", s.author)

	maxLength := aa.MaxOfLooksLength(s.text, signature)
	horizontalBar := strings.Repeat("-", maxLength+4)

	// 最大長が奇数なら1つ足しておく
	// なぜならずれるから
	if maxLength%2 == 1 {
		maxLength++
	}

	s.AsciiArt = []string{
		"",
		"",
		fmt.Sprintf("%s%s", strings.Repeat(" ", (maxLength+2)/2-aa.GetLooksLength(s.text)/2), s.text),
		horizontalBar,
		fmt.Sprintf("%s%s", strings.Repeat(" ", (maxLength+2)/2-aa.GetLooksLength(signature)/2), signature),
	}

	return nil
}
