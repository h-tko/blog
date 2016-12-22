package libraries

import (
    "github.com/russross/blackfriday"
)

func MarkdownToHtml(markdown string) []byte {
    return blackfriday.MarkdownCommon([]byte(markdown))
}
