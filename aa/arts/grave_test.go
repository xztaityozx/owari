package arts

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewGrave(t *testing.T) {
	g := NewGrave("text")
	assert.Equal(t, "text", g.text)
}

func TestGrave_Load(t *testing.T) {
	t.Run("textを指定しない場合", func(t *testing.T) {
		g := NewGrave("")
		err := g.Load("default")

		assert.Nil(t, err)
		expect := `    ┌─┐
    │先│
    │祖│
    │代│
    │々│
    │之│
    │ば│
    │か│
   ┌┴─┴┐
 │| 三三 |│
￣￣￣￣￣￣￣`
		assert.Equal(t, expect, strings.Join(g.AsciiArt, "\n"))
	})

	t.Run("textを指定する場合", func(t *testing.T) {
		g := NewGrave("あいうえお")
		err := g.Load("default")

		assert.Nil(t, err)
		expect := `    ┌─┐
    │あ│
    │い│
    │う│
    │え│
    │お│
   ┌┴─┴┐
 │| 三三 |│
￣￣￣￣￣￣￣`
		assert.Equal(t, expect, strings.Join(g.AsciiArt, "\n"))

	})
}
