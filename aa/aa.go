package aa

// IPrintableAsciiArt は出力可能なアスキーアートを表す
type IPrintableAsciiArt interface {
	Load(font string) error
}

type AsciiArt []string
