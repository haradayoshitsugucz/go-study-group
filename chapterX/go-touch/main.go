package main

import (
	"fmt"
	"os"
)

func main() {

	//fmt.Println(os.Args)
	//fmt.Println(os.Args[1])

	if len(os.Args) < 2 {
		fmt.Println(os.Stderr, "入力してください。")
	}

	path := os.Args[1]

	file, err := os.Open(path)
	defer file.Close()

	if err == nil {
		return
	}
	if err != nil {
		fmt.Fprint(os.Stderr, "エラーが発生しました。 err = %+v\n", err)
	}

	_, err = os.Create(path)
	if err != nil {
		fmt.Fprint(os.Stderr, "エラーが発生しました。 err = %+v\n", err)
	}

}
