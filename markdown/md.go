package markdown

import (
	"flag"
	"fmt"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var (
	infile string
)

func init() {
	flag.StringVar(&infile, "f", "", "input markdown file")
}

func Md() {
	flag.Parse()
	input := "## Hello,world"
	unsafe := blackfriday.MarkdownCommon([]byte(input))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Printf("html: %s", string(html))
}
