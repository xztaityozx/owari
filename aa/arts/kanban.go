package arts

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xztaityozx/owari/aa"
)

// Kanban は
// //
////|￣￣￣￣￣￣￣￣￣|
////|        終        |
////|    制作・著作    |
////|  ￣￣￣￣￣￣￣  |
////|     Ｏ Ｗ Ｒ     |
////|＿＿＿＿＿＿＿＿＿|
////   ∧∧  ||
////  ( ﾟдﾟ)||
////  /    づΦ
// です
type Kanban struct {
	aa.AsciiArt
	// 終 の代わりになるテキスト。改行区切りの[]string
	text []string
	// ギコ猫の代わりにこなたを出力する
	konata bool
	// ギコ猫に看板を持たせる
	giko bool
	// 制作・著作の下に入れる文字を変える
	author string
	// ギコ猫の位置を左から右にする
	reverse bool
	// ギコ猫が二匹になる
	twinGiko bool
}

// NewKanban は Kanban を作って返す
func NewKanban(text []string) Kanban {
	if text == nil {
		text = []string{"終"}
	}
	return Kanban{
		text:   text,
		author: "Ｏ Ｗ Ｒ",
	}
}

func (k *Kanban) SetKonata(b bool) {
	k.konata = b
}

func (k *Kanban) SetGikoNeko(b bool) {
	k.giko = b
}

func (k *Kanban) SetTwinGikoNeko(b bool) {
	k.twinGiko = b
	if b {
		k.giko = b
	}
}

func (k *Kanban) SetReverse(b bool) {
	k.reverse = b
}

func (k *Kanban) SetAuthor(author string) {
	if len(author) == 0 {
		k.author = "Ｏ Ｗ Ｒ"
	} else {
		k.author = author
	}
}

type KanbanTemplate struct {
	TopString    string
	BottomString string
	SideString   string
	Giko         map[string][]string
	ReverseGiko  map[string][]string
	TwinGiko     map[string][]string
	Konata       map[string][]string
}

type GikoTemplateLine struct {
	Text     string
	CenterAt int
}

// Load は AsciiArt メンバーに出力するAAを格納する
func (k *Kanban) Load(font string) error {
	var kt KanbanTemplate
	if b, err := raw.ReadFile("raw/kanban.json"); err != nil {
		return err
	} else {
		if err := json.Unmarshal(b, &kt); err != nil {
			return err
		}
	}

	// 見た目上の長さが奇数なら偶数に合わせる
	for i := 0; i < len(k.text); i++ {
		if aa.GetLooksLength(k.text[i])%2 == 1 {
			k.text[i] = " " + k.text[i]
		}
	}

	// Kanbanのテキスト部分の最小幅。k.textがどんな値だろうとこれ以下の幅にならない
	const minOfKanbanWidth = 14
	k.AsciiArt = []string{}

	// 一番長い文字列の長さを得る
	maxOfLooksLength := aa.MaxOfLooksLength(k.text...)
	// 一番長くても minOfKanbanWidth 文字未満なら最大の長さを18としておく。幅維持のため
	if maxOfLooksLength < minOfKanbanWidth {
		maxOfLooksLength = minOfKanbanWidth
	}

	// 看板の上端部分を生成
	k.AsciiArt = append(
		k.AsciiArt,
		fmt.Sprintf("%s%s%s", kt.SideString, strings.Repeat(kt.TopString, (maxOfLooksLength+4)/2), kt.SideString),
	)

	// 看板の中身
	for _, v := range append(k.text, "制作・著作") {
		k.AsciiArt = append(
			k.AsciiArt,
			fmt.Sprintf("%s  %s  %s", kt.SideString, alignCenter(v, maxOfLooksLength), kt.SideString),
		)
	}

	// 内側の水平線, 名前, 下端
	k.AsciiArt = append(
		k.AsciiArt,
		fmt.Sprintf("%s  %s  %s", kt.SideString, strings.Repeat(kt.TopString, maxOfLooksLength/2), kt.SideString),
		fmt.Sprintf("%s  %s  %s", kt.SideString, alignCenter(k.author, maxOfLooksLength), kt.SideString),
		fmt.Sprintf("%s%s%s", kt.SideString, strings.Repeat(kt.BottomString, (maxOfLooksLength+4)/2), kt.SideString),
	)

	// ギコ猫やこなたなどのキャラクター
	var character []string
	var ok bool
	const def = "default"
	if k.giko && k.reverse {
		character, ok = kt.ReverseGiko[font]
		if !ok {
			character = kt.ReverseGiko[def]
		}
	} else if k.giko && k.twinGiko {
		character, ok = kt.TwinGiko[font]
		if !ok {
			character = kt.TwinGiko[def]
		}
	} else if k.giko {
		character, ok = kt.Giko[font]
		if !ok {
			character = kt.Giko[def]
		}
	} else if k.konata {
		character, ok = kt.Konata[font]
		if !ok {
			character = kt.Konata[def]
		}
	} else {
		character = nil
	}

	for _, v := range character {
		k.AsciiArt = append(
			k.AsciiArt,
			alignCenter(v, maxOfLooksLength+6),
		)
	}

	return nil
}

// alignCenter は 見た目状の長さが width な文字列の中心に str を配置した文字列を返す
func alignCenter(str string, width int) string {
	ll := aa.GetLooksLength(str)
	// もともと str のほうが長かったら、後ろから削り、widthにおさまった時点で返す
	if ll > width {
		for i := len(str); i >= 0; i-- {
			t := str[0:i]
			l := aa.GetLooksLength(t)
			if l == width {
				// 長さが一致したらそのまま返す
				return t
			} else if l == width-1 {
				// 1つだけ短かったら戦闘に半角スペースを足して返す
				return " " + t
			}
		}

		// width が負だったときに来るか…？
		return ""
	}

	// if ll <= width {
	diff := width - ll
	return fmt.Sprintf("%s%s%s",
		strings.Repeat(" ", diff-(diff/2)),
		str,
		strings.Repeat(" ", diff/2),
	)
	// }
}
