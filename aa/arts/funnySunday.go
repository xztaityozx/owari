package arts

import (
	"encoding/json"
	"github.com/xztaityozx/owari/aa"
)
import _ "embed"

type FunnySunday struct {
	aa.AsciiArt
}

func NewFunnySunday() FunnySunday { return FunnySunday{} }

func (fs *FunnySunday) Load(_ string) error {
	var err error
	var templateJson []byte
	templateJson, err = raw.ReadFile("raw/funnySunday.json")
	if err != nil {
		return err
	}

	var template map[string][]string
	if err := json.Unmarshal(templateJson, &template); err != nil {
		return err
	}

	// funnySunday.jsonにテンプレを追加したなら
	// この関数の引数を使ってテンプレを選択する感じに
	//fs.AsciiArt, ok = template[font]
	//if !ok {
	//	fs.AsciiArt = template["default"]
	//}

	fs.AsciiArt = template["default"]

	return nil
}
