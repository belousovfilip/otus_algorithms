package main

import (
	"fmt"
	"io"
)

func paint3(sizeW, sizeH int, writer io.Writer) {
	for h := 0; h < sizeH; h++ {
		for w := 0; w < sizeW; w++ {
			if w == sizeH-h-1 {
				fmt.Fprint(writer, " # ")
			} else {
				fmt.Fprint(writer, " . ")
			}
		}
		fmt.Fprint(writer, "\n")
	}
}
func paint6(sizeW, sizeH int, writer io.Writer) {
	for h := 0; h < sizeH; h++ {
		for w := 0; w < sizeW; w++ {
			if h < 10 || w < 10 {
				fmt.Fprint(writer, " # ")
			} else {
				fmt.Fprint(writer, " . ")
			}
		}
		fmt.Fprint(writer, "\n")
	}
}

func paint8(sizeW, sizeH int, writer io.Writer) {
	for h := 0; h < sizeH; h++ {
		for w := 0; w < sizeW; w++ {
			if h == 0 || w == 0 {
				fmt.Fprint(writer, " # ")
			} else {
				fmt.Fprint(writer, " . ")
			}
		}
		fmt.Fprint(writer, "\n")
	}
}

func paint11(sizeW, sizeH int, writer io.Writer) {
	for h := 0; h < sizeH; h++ {
		for w := 0; w < sizeW; w++ {
			if h == 1 || w == 1 || w == sizeW-2 || h == sizeH-2 {
				fmt.Fprint(writer, " # ")
			} else {
				fmt.Fprint(writer, " . ")
			}
		}
		fmt.Fprint(writer, "\n")
	}
}

func paint13(sizeW, sizeH int, writer io.Writer) {
	for h := 0; h < sizeH; h++ {
		for w := 0; w < sizeW; w++ {
			if w+h > 19 && w+h < 29 {
				fmt.Fprint(writer, " # ")
			} else {
				fmt.Fprint(writer, " . ")
			}
		}
		fmt.Fprint(writer, "\n")
	}
}
