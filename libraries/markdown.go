package libraries

import (
	"github.com/russross/blackfriday"
)

func MarkdownToHtml(markdown []byte) []byte {
	return blackfriday.MarkdownCommon(markdown)
}
