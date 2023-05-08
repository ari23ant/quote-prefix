package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

var (
	prefix string
)

func init() {
	flag.StringVar(&prefix, "p", "> ", "prefix string")
}

func main() {

	// コマンドライン引数をパースして引用接頭語を取得
	flag.Parse()

	text, err := clipboard.ReadAll()

	if err != nil {
		// The operation completed successfully.(err)
		//fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "clipboard is NOT string")
		return
	}

	// 行頭に接頭語を結合
	text = prefix + text
	text = strings.Replace(text, "\n", "\n"+prefix, -1)

	clipboard.WriteAll(text)
	//println(text)
}
