package markdown

import (
	"flag"
	"fmt"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"strings"
)

var (
	infile  string
	outfile string
)

func init() {
	flag.StringVar(&infile, "f", "hello.md", "input markdown filepath")
	flag.StringVar(&outfile, "o", "", "out html filepath")
}

func Md() {
	flag.Parse()
	if outfile == "" {
		arr := strings.Split(infile, ".md")
		outfile = arr[0] + ".html"
	}
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
	err = ioutil.WriteFile(outfile, html, 0644)
	if err != nil {
		panic(err)
	}
}
