package cmd

import "testing"

func TestMakePaddedText(t *testing.T) {
	type TestData struct {
		texts  []string
		ret    []string
		length int
	}
	testDatas := []TestData{
		{texts: []string{"a", "bbbbb"}, ret: []string{"   a  ", " bbbbb"}, length: 6},
		{texts: []string{"あいう", "えお"}, ret: []string{"あいう", " えお "}, length: 6},
		{texts: []string{"", ""}, ret: []string{"", ""}, length: 0},
		{texts: []string{}, ret: []string{}, length: 0},
	}
	for _, v := range testDatas {
		ret, l := makePaddedText(v.texts)
		for i, r := range ret {
			text := v.texts[i]
			if r != text {
				t.Errorf("文字列が不一致:引数 = %v, 期待値 = %s, 実際 = %s", v.texts, r, text)
			}
		}
		if v.length != l {
			t.Errorf("最長文字列の長さが不一致:引数 = %v, 期待値 = %d, 実際 = %d", v.texts, v.length, l)
		}
	}
}
func TestPadSpace(t *testing.T) {
	type TestData struct {
		text string
		max  int
		ret  string
	}
	testDatas := []TestData{
		{text: "a ", max: 2, ret: "a "},
		{text: "a ", max: 4, ret: " a  "},
		{text: "あ", max: 2, ret: "あ"},
		{text: "あ", max: 4, ret: " あ "},
	}
	for _, v := range testDatas {
		ret := padSpace(v.text, v.max)
		if v.ret != ret {
			t.Errorf("パディングされた文字列が不一致:引数text = <%s>, 引数max = <%d>, 期待値 = %s, 実際 = %s", v.text, v.max, v.ret, ret)
		}
	}
}
