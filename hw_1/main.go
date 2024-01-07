package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	b := true
	if b {
		w = os.Stdout
		paint(25, 25, w)
	} else {
		w, err := os.OpenFile("test/paint_.txt", os.O_CREATE|os.O_RDWR, 0777)
		defer w.Close()
		if err != nil {
			panic(err)
		}
		paint(25, 25, w)
	}

}

func paint(sizeW, sizeH int, writer io.Writer) {
	for h := 0; h < sizeH; h++ {
		for w := 0; w < sizeW; w++ {
			if h == 0 || w == 0 || w == sizeW-1 || h == sizeH-1 {
				fmt.Fprint(writer, " # ")
			} else {
				fmt.Fprint(writer, " . ")
			}
		}
		fmt.Fprint(writer, "\n")
	}
}
