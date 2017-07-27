package markdown

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var (
	infile    string
	targetDir string
)

func init() {
	flag.StringVar(&infile, "f", "hello.md", "input markdown filepath")
	flag.StringVar(&targetDir, "t", "", "target dir for output html file")
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
	arr := strings.Split(infile, ".md")
	out := arr[0] + ".html"
	if targetDir != "" {
		if !strings.HasSuffix(targetDir, "/") {
			targetDir = targetDir + "/"
		}
		out = targetDir + out
	}
	err = ioutil.WriteFile(out, html, 0644)
	if err != nil {
		panic(err)
	}
}
