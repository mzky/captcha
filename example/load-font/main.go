package main

import (
	"bytes"
	"fmt"
	"github.com/mzky/captcha"
	"golang.org/x/image/font/gofont/goregular"
	"html/template"
	"io"
	"net/http"
)

func writeToBuffer(w io.Writer) []byte {
	buf := new(bytes.Buffer)
	io.Copy(w, buf)
	return buf.Bytes()
}
func main() {
	err := captcha.LoadFont(goregular.TTF)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/captcha", captchaHandle)
	fmt.Println("Server start at port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func indexHandle(w http.ResponseWriter, _ *http.Request) {
	doc, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	doc.Execute(w, nil)
}

func captchaHandle(w http.ResponseWriter, _ *http.Request) {
	img, err := captcha.New(120, 35, func(options *captcha.Options) {
		options.FontScale = 0.8
	})
	if err != nil {
		fmt.Fprint(w, nil)
		fmt.Println(err.Error())
		return
	}
	img.WriteImage(w)
}
