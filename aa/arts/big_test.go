package arts

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestNewBigOwari(t *testing.T) {
	bo := NewBigOwari("text")
	assert.NotNil(t, bo)
	assert.Equal(t, "text", bo.text)
}

func TestBigOwari_Load(t *testing.T) {
	now := time.Now()
	if now.Month() == 12 && now.Day() == 25 {
		t.Skip("12/25はちゃんと動かないテストなのでスキップだよ")
	}

	bo := NewBigOwari("")

	assert.NotNil(t, bo)

	err := bo.Load("")
	assert.Nil(t, err)
	assert.Equal(t, bigAA, strings.Join(bo.AsciiArt, "\n"))
}

func TestBigOwari_Load2(t *testing.T) {
	now := time.Now()
	if now.Month() != 12 || now.Day() != 25 {
		t.Skip("12/25にしか動かないテストなのでスキップだよ")
	}

	bo := NewBigOwari("")

	assert.NotNil(t, bo)

	err := bo.Load("")
	assert.Nil(t, err)
	assert.Equal(t, bigHolly, strings.Join(bo.AsciiArt, "\n"))
}
