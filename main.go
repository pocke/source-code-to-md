package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	for _, fname := range os.Args[1:] {
		err := Translate(fname, os.Stdout)
		if err != nil {
			panic(err)
		}
	}
}

func Translate(fname string, w io.Writer) error {
	lang := FnameToLang(fname)
	fmt.Fprintf(w, "- %s\n\n```%s\n", fname, lang)

	defer fmt.Fprint(w, "\n```\n\n")

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	w.Write(b)

	return nil
}

func FnameToLang(fname string) string {
	if fname == "Makefile" {
		return "make"
	}
	ext := filepath.Ext(fname)
	if ext == "" {
		return ""
	}
	switch ext[1:] {
	case "js":
		return "javascript"
	default:
		return ext[1:]
	}
}
