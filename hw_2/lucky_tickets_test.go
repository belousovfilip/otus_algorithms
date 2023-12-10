package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLuckyTickets(t *testing.T) {
	data, err := getData(10)
	require.NoError(t, err)
	for _, v := range data {
		require.Equal(t, v.out, LuckyTickets(v.in))
		_, err := fmt.Fprintf(os.Stdout, "success: in: %d; out: %d \n", v.in, v.out)
		require.NoError(t, err)
	}
}
func BenchmarkLuckyTickets(b *testing.B) {
	b.StopTimer()
	data, err := getData(10)
	require.NoError(b, err)
	b.StartTimer()
	for _, v := range data {
		b.StopTimer()
		name := fmt.Sprintf("in:__%d__", v.in)
		b.StartTimer()
		b.Run(name, func(b *testing.B) {
			n := v.in
			for i := 0; i < b.N; i++ {
				LuckyTickets(n)
			}
		})
	}
}

func readIntFromFile(p string) (int64, error) {
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
	return int64(out), nil

}

func getData(n int) ([]struct {
	in  uint16
	out int64
}, error) {
	var data []struct {
		in  uint16
		out int64
	}
	for i := 0; i <= n; i++ {
		in, err := readIntFromFile(fmt.Sprintf("./test/test.%d.in", i))
		if err != nil {
			return nil, err
		}
		out, err := readIntFromFile(fmt.Sprintf("./test/test.%d.out", i))
		if err != nil {
			return nil, err
		}
		data = append(data, struct {
			in  uint16
			out int64
		}{in: uint16(in), out: out})
	}
	return data, nil
}
