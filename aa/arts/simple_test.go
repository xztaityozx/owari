package arts

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewSimpleOwari(t *testing.T) {
	so := NewSimpleOwari("text", "author")

	assert.Equal(t, "text", so.text)
	assert.Equal(t, "author", so.author)
}

func TestSimpleOwari_Load(t *testing.T) {
	so := NewSimpleOwari("", "")

	err := so.Load("")
	assert.Nil(t, err)

	expect := []string{
		"", "",
		"        糸冬",
		"---------------------",
		"  制作・著作 ＮＨＫ",
	}

	assert.Equal(t, strings.Join(expect, "\n"), strings.Join(so.AsciiArt, "\n"))
}
