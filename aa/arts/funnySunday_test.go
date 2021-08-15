package arts

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewFunnySunday(t *testing.T) {
	fs := NewFunnySunday()
	assert.NotNil(t, fs)
}

func TestFunnySunday_Load(t *testing.T) {
	fs := NewFunnySunday()
	_ = fs.Load("")
	//assert.Nil(t, err)

	expect := `たのしい休日

                               /||￣￣|| ∧∧
                              l ||＿＿||(,/⌒ヽ
                             [l  | ,,― とﾉ    ヽ
                        |￣￣￣￣￣     (＿＿＿ )

                                              ┼ヽ   -|r‐､.  ﾚ  |
                                             ｄ⌒)  ./| _ﾉ    __ﾉ` + "`\n"

	assert.Equal(t, expect, strings.Join(fs.AsciiArt, "\n"))
}
