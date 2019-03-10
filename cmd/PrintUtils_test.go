package cmd

import "testing"

func TestGetLooksLength(t *testing.T) {
	type TestData struct {
		text   string
		length int
	}
	testDatas := []TestData{
		{text: "a", length: 1},
		{text: "あ", length: 2},
		{text: "漢", length: 2},
		{text: "ａ", length: 2},
		{text: "Ａ", length: 2},
		{text: " ", length: 1},
		{text: "　", length: 2},
		{text: "あいうえお", length: 10},
	}
	for _, v := range testDatas {
		l := GetLooksLength(v.text)
		if v.length != l {
			t.Errorf("文字列の長さが不一致:引数 = <%s>, 期待値 = %d, 実際 = %d", v.text, v.length, l)
		}
	}
}
