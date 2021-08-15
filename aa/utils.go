package aa

import (
	"github.com/mattn/go-runewidth"
)

// GetLooksLength は見た目上のテキストの長さを返す。
// 例：
//   a  == 1
//   あ == 2
//   漢 == 2
//   ｂ == 2
// 見た目上の長さを特別扱いしたい文字などが登場したらここに追加していく感じでヨロ！
func GetLooksLength(text string) int {
	cond := runewidth.NewCondition()
	// EastAsianWidthを強制します
	cond.EastAsianWidth = true

	return cond.StringWidth(text)
}

// MaxOfLooksLength は引数に渡された文字列のうち見た目上の長さが最大なものの長さを返す
func MaxOfLooksLength(s ...string) int {
	max := 0
	for _, v := range s {
		l := GetLooksLength(v)
		if l > max {
			max = l
		}
	}

	return max
}
