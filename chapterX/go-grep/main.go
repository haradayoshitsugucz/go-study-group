package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	//fmt.Println(os.Args)
	//fmt.Println(os.Args[1])

	var reader io.Reader

	pattern := os.Args[1]

	if len(os.Args) == 2 {
		reader = os.Stdin
	} else {
		path := os.Args[2]
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			fmt.Println(os.Stderr, "err=%+v", err)
		}
		reader = file
	}

	scan := bufio.NewScanner(reader)

	for scan.Scan() {
		txt := scan.Text()
		if strings.Contains(txt, pattern) {
			fmt.Fprintf(os.Stdout, "%v\n", txt)
		}
	}


	// build のあとに go install うつとコマンド実行できる

	// goの慣習で err==nil だったら、fileがnilになることはない！！

	// model, ok, err で返すとか
	// nil, nil で返すべきではない！


}
