package main

import (
	"bytes"
	"embed"
	_ "embed"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/defkoi/passgen"
	webwiev "github.com/webview/webview_go"
)

var (
	length int
	script string
	view   bool
)

func init() {
	flag.IntVar(&length, "l", 8, "password length")
	flag.StringVar(&script, "s", "", "external script")
	flag.BoolVar(&view, "w", false, "enable webview")
	flag.Parse()
}

func init() {
	if script != "" {
		source, err := os.ReadFile(script)
		if err != nil {
			log.Fatal(err)
		}
		passgen.SetScript(string(source))
	}
}

func main() {
	if view {
		webView()
		return
	}

	if v, err := passgen.GeneratePassword(length); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}

func webView() {
	w := webwiev.New(false)
	defer w.Destroy()

	w.SetSize(400, 400, webwiev.HintFixed)
	w.SetTitle("password generator")
	w.SetHtml(parseWebViewHTML())

	w.Bind("generatePassword", func(length int) (string, error) {
		return passgen.GeneratePassword(int(length))
	})

	w.Run()
}

//go:embed webview
var webViewFiles embed.FS

func parseWebViewHTML() string {
	t, err := template.ParseFS(webViewFiles, "**/*.html", "**/*.css", "**/*.js")
	if err != nil {
		panic(err)
	}

	var buf = bytes.NewBuffer([]byte{})
	data := map[string]any{"length": length}
	if err := t.Lookup("index.html").Execute(buf, data); err != nil {
		panic(err)
	}

	return buf.String()
}
