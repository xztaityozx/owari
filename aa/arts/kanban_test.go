package arts

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKanban(t *testing.T) {
	k := NewKanban([]string{"text"})
	assert.NotNil(t, k)
	assert.False(t, k.giko)
	assert.False(t, k.reverse)
	assert.False(t, k.konata)
	assert.False(t, k.twinGiko)
	assert.Equal(t, []string{"text"}, k.text)
}

func TestKanban_Load(t *testing.T) {
	t.Run("装飾がないやつ", func(t *testing.T) {
		k := NewKanban(nil)
		k.SetAuthor("")
		assert.NotNil(t, k)
		err := k.Load("default")
		assert.Nil(t, err)

		expect := []string{
			"|￣￣￣￣￣￣￣￣￣|",
			"|        終        |",
			"|    制作・著作    |",
			"|  ￣￣￣￣￣￣￣  |",
			"|     Ｏ Ｗ Ｒ     |",
			"|＿＿＿＿＿＿＿＿＿|",
		}

		assert.Equal(t, strings.Join(expect, "\n"), strings.Join(k.AsciiArt, "\n"))
	})

	t.Run("text1行", func(t *testing.T) {
		k := NewKanban([]string{"おーーーーーーーーーーーーわり！"})
		assert.NotNil(t, k)
		err := k.Load("default")
		assert.Nil(t, err)

		expect := []string{
			"|￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣|",
			"|  おーーーーーーーーーーーーわり！  |",
			"|             制作・著作             |",
			"|  ￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣  |",
			"|              Ｏ Ｗ Ｒ              |",
			"|＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿|",
		}

		assert.Equal(t, strings.Join(expect, "\n"), strings.Join(k.AsciiArt, "\n"))
	})

	t.Run("text2行", func(t *testing.T) {
		k := NewKanban([]string{
			"おーーーーーーーーーーーーわり！",
			"おーーーーーーーーーーーーわり！",
		})
		assert.NotNil(t, k)
		err := k.Load("default")
		assert.Nil(t, err)

		expect := []string{
			"|￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣|",
			"|  おーーーーーーーーーーーーわり！  |",
			"|  おーーーーーーーーーーーーわり！  |",
			"|             制作・著作             |",
			"|  ￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣  |",
			"|              Ｏ Ｗ Ｒ              |",
			"|＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿|",
		}

		assert.Equal(t, strings.Join(expect, "\n"), strings.Join(k.AsciiArt, "\n"))
	})

	t.Run("giko", func(t *testing.T) {
		k := NewKanban([]string{"終"})
		k.SetGikoNeko(true)
		assert.NotNil(t, k)
		err := k.Load("default")
		assert.Nil(t, err)

		expect := []string{
			"|￣￣￣￣￣￣￣￣￣|",
			"|        終        |",
			"|    制作・著作    |",
			"|  ￣￣￣￣￣￣￣  |",
			"|     Ｏ Ｗ Ｒ     |",
			"|＿＿＿＿＿＿＿＿＿|",
			"  ∧∧   ||         ",
			"  ( ﾟдﾟ)||           ",
			"  /    づΦ         ",
		}

		assert.Equal(t, strings.Join(expect, "\n"), strings.Join(k.AsciiArt, "\n"))
	})

	t.Run("twin giko", func(t *testing.T) {
		k := NewKanban([]string{"終"})
		k.SetTwinGikoNeko(true)
		assert.NotNil(t, k)
		err := k.Load("default")
		assert.Nil(t, err)

		expect := []string{
			"|￣￣￣￣￣￣￣￣￣|",
			"|        終        |",
			"|    制作・著作    |",
			"|  ￣￣￣￣￣￣￣  |",
			"|     Ｏ Ｗ Ｒ     |",
			"|＿＿＿＿＿＿＿＿＿|",
			"  ∧∧   ||  ∧∧   ",
			"  ( ﾟдﾟ)||(ﾟдﾟ )      ",
			"  /    づΦ⊂    \\  ",
		}

		assert.Equal(t, strings.Join(expect, "\n"), strings.Join(k.AsciiArt, "\n"))
	})

	t.Run("giko reverse", func(t *testing.T) {
		k := NewKanban([]string{"終"})
		k.SetGikoNeko(true)
		k.SetReverse(true)
		assert.NotNil(t, k)
		err := k.Load("default")
		assert.Nil(t, err)

		expect := []string{
			"|￣￣￣￣￣￣￣￣￣|",
			"|        終        |",
			"|    制作・著作    |",
			"|  ￣￣￣￣￣￣￣  |",
			"|     Ｏ Ｗ Ｒ     |",
			"|＿＿＿＿＿＿＿＿＿|",
			"         ||  ∧∧   ",
			"         ||(ﾟдﾟ )    ",
			"         Φ⊂    \\  ",
		}

		assert.Equal(t, strings.Join(expect, "\n"), strings.Join(k.AsciiArt, "\n"))
	})

	t.Run("konata", func(t *testing.T) {
		k := NewKanban([]string{"終"})
		k.SetKonata(true)
		assert.NotNil(t, k)
		err := k.Load("default")
		assert.Nil(t, err)

		expect := []string{
			"|￣￣￣￣￣￣￣￣￣|",
			"|        終        |",
			"|    制作・著作    |",
			"|  ￣￣￣￣￣￣￣  |",
			"|     Ｏ Ｗ Ｒ     |",
			"|＿＿＿＿＿＿＿＿＿|",
			" ∧ ∧   ||         ",
			"(≡ω≡.)||         ",
			"/      づΦ         ",
		}

		assert.Equal(t, strings.Join(expect, "\n"), strings.Join(k.AsciiArt, "\n"))
	})

	t.Run("custom author", func(t *testing.T) {
		k := NewKanban([]string{"終"})
		k.SetAuthor("TEST")
		assert.NotNil(t, k)
		err := k.Load("default")
		assert.Nil(t, err)

		expect := []string{
			"|￣￣￣￣￣￣￣￣￣|",
			"|        終        |",
			"|    制作・著作    |",
			"|  ￣￣￣￣￣￣￣  |",
			"|       TEST       |",
			"|＿＿＿＿＿＿＿＿＿|",
		}

		assert.Equal(t, strings.Join(expect, "\n"), strings.Join(k.AsciiArt, "\n"))
	})
}
