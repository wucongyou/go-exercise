package markdown

import (
	"flag"
	"fmt"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
)

var (
	infile string
)

func init() {
	flag.StringVar(&infile, "f", "hello.md", "input markdown file")
}

func Md() {
	flag.Parse()
	f, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	input := string(fd)
	unsafe := blackfriday.MarkdownCommon([]byte(input))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Printf("html: %s", string(html))
}
