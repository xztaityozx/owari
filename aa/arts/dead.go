package arts

import (
	"fmt"
	"github.com/xztaityozx/owari/aa"
	"strings"
)

type Dead struct {
	aa.AsciiArt

	// 終わってしまったものの名前
	Name string
}

// NewDead は Dead を作って返す
func NewDead(name string) Dead {
	return Dead{Name: name}
}

var template = `
              ,, ＿
            ／       ｀､ 
           /   (_ﾉL_） ヽ 
          /  ´・  ・｀  l        %sは死んだんだ
         （l     し     l）       いくら呼んでも帰っては来ないんだ 
          l     ＿＿    l        もうあの時間は終わって、君も人生と向き合う時なんだ 
           >  ､ _     ィ 
         ／         ￣   ヽ 
        /  |             i ヽ 
        |＼|             |/| 
        |  ||/＼／＼／＼/| | 
`

func (d *Dead) Load(_ string) error {
	d.AsciiArt = strings.Split(fmt.Sprintf(template, d.Name), "\n")
	return nil
}
