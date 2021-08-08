package arts

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/xztaityozx/owari/aa"
	"log"
)

// go:embed raw/grave.json
var graveTemplateContent []byte

// Grave は
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
// ￣￣￣￣￣￣￣
// です
type Grave struct {
	aa.AsciiArt
	text string
}

func NewGrave(text string) Grave {
	if len(text) == 0 {
		text = "先祖代々之ばか"
	}
	return Grave{text: text}
}

func (g Grave) Load(font string) error {
	var graveTemplate map[string][]string

	log.Println(graveTemplateContent)

	if err := json.Unmarshal(graveTemplateContent, &graveTemplate); err != nil {
		return err
	}

	template, ok := graveTemplate[font]
	if !ok {
		template = graveTemplate["default"]
	}

	g.AsciiArt = []string{
		template[0],
	}

	for _, c := range g.text {
		t := string(c)
		if aa.GetLooksLength(t) == 1 {
			t = fmt.Sprintf(" %c", c)
		}

		g.AsciiArt = append(g.AsciiArt, fmt.Sprintf(template[1], t))
	}

	return nil
}
