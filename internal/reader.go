package internal

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadIntFromFile(p string) (int, error) {
	f, err := os.Open(p)
	if err != nil {
		return 0, err
	}
	raw, err := bufio.NewReader(f).ReadString('\r')
	if err != nil && err != io.EOF {
		return 0, err
	}
	raw = strings.TrimRight(raw, "\r")
	raw = strings.TrimRight(raw, "\n")
	out, err := strconv.Atoi(raw)
	if err != nil {
		return 0, err
	}
	return out, nil

}

func ReadInt64FromFile(p string) (int64, error) {
	f, err := os.Open(p)
	if err != nil {
		return 0, err
	}
	raw, err := bufio.NewReader(f).ReadString('\r')
	if err != nil && err != io.EOF {
		return 0, err
	}
	raw = strings.TrimRight(raw, "\r")
	raw = strings.TrimRight(raw, "\n")
	out, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, err
	}
	return out, nil
}
